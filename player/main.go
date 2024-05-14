package player

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pauloo27/tuner/player/mpv"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
	"github.com/kkdai/youtube/v2"
)

const (
	maxVolume = 150.0
)

var (
	MpvInstance   *mpv.Mpv
	State         *PlayerState
	defaultClient = youtube.Client{}
)

func Initialize() {
	var err error

	// create a mpv instance
	MpvInstance = mpv.Create()

	// load data
	data := storage.Load()

	// math.Min is used to avoid high volumes
	initialVolume := math.Min(data.DefaultVolume, 100)
	// if none is defined, 50% is used
	if initialVolume == 0.0 {
		initialVolume = 50
	}

	// set options
	// disable video
	err = MpvInstance.SetOptionString("video", "no")
	utils.HandleError(err, "Cannot set mpv video option")

	// disable cache
	if data.Cache {
		err = MpvInstance.SetOptionString("cache", "no")
		utils.HandleError(err, "Cannot set mpv cache option")
	}

	// set default volume value
	err = MpvInstance.SetOption("volume", mpv.FORMAT_DOUBLE, initialVolume)
	utils.HandleError(err, "Cannot set mpv volume option")

	// set default volume value
	err = MpvInstance.SetOption("volume-max", mpv.FORMAT_DOUBLE, maxVolume)
	utils.HandleError(err, "Cannot set mpv volume-max option")

	// add observers
	err = MpvInstance.ObserveProperty(0, "volume", mpv.FORMAT_DOUBLE)
	utils.HandleError(err, "Cannot observer volume property")

	err = MpvInstance.ObserveProperty(0, "pause", mpv.FORMAT_FLAG)
	utils.HandleError(err, "Cannot observer pause property")

	err = MpvInstance.ObserveProperty(0, "loop-file", mpv.FORMAT_FLAG)
	utils.HandleError(err, "Cannot observer loop-file property")

	err = MpvInstance.ObserveProperty(0, "loop-playlist", mpv.FORMAT_FLAG)
	utils.HandleError(err, "Cannot observer loop-playlist property")

	// start event listener
	startEventHandler()

	// create the state
	State = &PlayerState{
		data,
		false, false,
		nil, nil, 0,
		initialVolume,
		0.0,
		false, false, false,
		SongLyric{},
		StatusLoopNone,
		false,
	}

	registerInternalHooks()

	// start the player
	err = MpvInstance.Initialize()
	utils.HandleError(err, "Cannot initialize mpv")

	callHooks(HookPlayerInitialized, err)
}

func registerInternalHooks() {
	RegisterHook(func(param ...interface{}) {
		if State.GetPlaying().Live {
			return
		}
		duration, err := MpvInstance.GetProperty("duration", mpv.FORMAT_DOUBLE)
		utils.HandleError(err, "Cannot get duration")
		State.Duration = duration.(float64)
	}, HookFileLoaded)

	RegisterHook(func(params ...interface{}) {
		// remove lyric from state
		State.Lyric.Lines = []string{}
		State.Lyric.Index = 0

		pos, err := MpvInstance.GetProperty("playlist-pos", mpv.FORMAT_INT64)
		if err != nil {
			return
		}
		newPos := int(pos.(int64))
		if newPos != State.PlaylistIndex {
			if newPos == -1 {
				newPos = 0
			}
			State.PlaylistIndex = newPos
			callHooks(HookPlaylistSongChanged)
		}
	}, HookFileEnded)

	RegisterHook(func(params ...interface{}) {
		State.Volume = params[0].(float64)
	}, HookVolumeChanged)

	RegisterHook(func(params ...interface{}) {
		if State.Loop == StatusLoopNone {
			State.Loop = StatusLoopTrack
		} else {
			State.Loop = StatusLoopNone
		}
	}, HookLoopTrackChanged)

	RegisterHook(func(params ...interface{}) {
		if State.Loop == StatusLoopNone {
			State.Loop = StatusLoopPlaylist
		} else {
			State.Loop = StatusLoopNone
		}
	}, HookLoopPlaylistChanged)

	RegisterHook(func(param ...interface{}) {
		State.Idle = true
		if State.IsPlaylist() {
			State.Playlist.Unshuffle()
		}
	}, HookIdle)

	RegisterHook(func(param ...interface{}) {
		State.Idle = false
	}, HookFileLoadStarted)
}

func ClearPlaylist() error {
	return MpvInstance.Command([]string{"playlist-clear"})
}

func RemoveCurrentFromPlaylist() error {
	return MpvInstance.Command([]string{"playlist-remove", "current"})
}

func Stop() error {
	return MpvInstance.Command([]string{"stop"})
}

func PlaylistNext() error {
	return MpvInstance.Command([]string{"playlist-next"})
}

func PlaylistPrevious() error {
	return MpvInstance.Command([]string{"playlist-prev"})
}

func ForceUpdate() {
	callHooks(HookGenericUpdate)
}

func SaveToPlaylist() {
	State.SavingToPlaylist = true
	callHooks(HookSavingTrackToPlaylist)
}

func extractLiveURL(live *youtube.Video) (string, error) {
	res, err := http.Get(live.HLSManifestURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	body := string(buffer)

	for _, line := range strings.Split(body, "\n") {
		if strings.HasPrefix(line, "https://") {
			return line, nil
		}
	}

	return "", errors.New("invalid HLS manifest")
}

func extractVideoURL(result *search.SearchResult) (string, error) {
	vid, err := defaultClient.GetVideo(result.URL)
	if result.Live {
		return extractLiveURL(vid)
	}
	if err != nil {
		return "", nil
	}

	formats := vid.Formats.Itag(251)

	if len(formats) > 0 {
		format := formats[0]
		return defaultClient.GetStreamURL(vid, &format)
	}

	return defaultClient.GetStreamURL(vid, &vid.Formats[0])
}

func PlaySearchResult(result *search.SearchResult, playlist *storage.Playlist) error {
	// remove all entries from playlist
	ClearPlaylist()

	// remove pause
	defer Play()

	State.Result = result
	State.Playlist = playlist

	if result == nil {
		for i := range playlist.Songs {
			song := playlist.SongAt(i)
			videoURL, err := extractVideoURL(song)
			if err != nil {
				return err
			}
			if i == 0 {
				err = LoadFile(videoURL, song.Title)
			} else {
				err = AppendFile(videoURL, song.Title)
			}
			if err != nil {
				return err
			}
		}
		return nil
	}
	videoURL, err := extractVideoURL(result)
	if err != nil {
		return err
	}
	return LoadFile(videoURL, result.Title)
}

func escapeTitle(title string) string {
	return fmt.Sprintf(`"%s"`, strings.ReplaceAll(title, `"`, "")) // FIXME
}

func LoadFile(filePath, title string) error {
	err := MpvInstance.Command([]string{"loadfile", filePath, "replace"})
	callHooks(HookFileLoadStarted, err, filePath)
	return err
}

func AppendFile(filePath, title string) error {
	err := MpvInstance.Command([]string{"loadfile", filePath, "append"})
	callHooks(HookFileLoadStarted, err, filePath)
	return err
}

func PlayPause() error {
	if State.Paused {
		return Play()
	}
	return Pause()
}

func Seek(seconds int) error {
	return MpvInstance.Command([]string{"seek", strconv.Itoa(seconds)})
}

func Pause() error {
	return MpvInstance.SetProperty("pause", mpv.FORMAT_FLAG, true)
}

func Play() error {
	return MpvInstance.SetProperty("pause", mpv.FORMAT_FLAG, false)
}

func LoopNone() error {
	err := MpvInstance.SetProperty("loop-file", mpv.FORMAT_FLAG, false)
	if err != nil {
		return err
	}
	err = MpvInstance.SetProperty("loop-playlist", mpv.FORMAT_FLAG, false)
	return err
}

func LoopTrack() error {
	err := MpvInstance.SetProperty("loop-file", mpv.FORMAT_FLAG, true)
	if err != nil {
		return err
	}
	err = MpvInstance.SetProperty("loop-playlist", mpv.FORMAT_FLAG, false)
	return err
}

func LoopPlaylist() error {
	err := MpvInstance.SetProperty("loop-file", mpv.FORMAT_FLAG, false)
	if err != nil {
		return err
	}
	err = MpvInstance.SetProperty("loop-playlist", mpv.FORMAT_FLAG, true)
	return err
}

func SetVolume(volume float64) error {
	volume = math.Min(maxVolume, volume)
	err := MpvInstance.SetProperty("volume", mpv.FORMAT_DOUBLE, volume)
	return err
}

func GetPosition() (float64, error) {
	position, err := MpvInstance.GetProperty("time-pos", mpv.FORMAT_DOUBLE)
	if err != nil {
		return 0.0, err
	}
	return position.(float64), err
}

func SetPosition(pos float64) error {
	return MpvInstance.SetProperty("time-pos", mpv.FORMAT_DOUBLE, pos)
}
