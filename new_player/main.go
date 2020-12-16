package new_player

import (
	"math"
	"strconv"

	"github.com/Pauloo27/tuner/new_player/mpv"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

const (
	maxVolume = 150.0
)

var MpvInstance *mpv.Mpv
var State *PlayerState

func Initialize() {
	var err error

	initialVolume := 50.0

	// create a mpv instance
	MpvInstance = mpv.Create()

	// set options
	// disable video
	err = MpvInstance.SetOptionString("video", "no")
	utils.HandleError(err, "Cannot set mpv video option")

	// disable cache
	err = MpvInstance.SetOptionString("cache", "no")
	utils.HandleError(err, "Cannot set mpv cache option")

	// set default volume value
	err = MpvInstance.SetOption("volume", mpv.FORMAT_DOUBLE, initialVolume)
	utils.HandleError(err, "Cannot set mpv volume option")

	// set default volume value
	err = MpvInstance.SetOption("volume-max", mpv.FORMAT_DOUBLE, maxVolume)
	utils.HandleError(err, "Cannot set mpv volume-max option")

	// set quality to worst
	err = MpvInstance.SetOptionString("ytdl-format", "worst")
	utils.HandleError(err, "Cannot set mpv ytdl-format option")

	// add observers
	err = MpvInstance.ObserveProperty(0, "volume", mpv.FORMAT_DOUBLE)
	utils.HandleError(err, "Cannot observer volume property")

	err = MpvInstance.ObserveProperty(0, "loop-file", mpv.FORMAT_FLAG)
	utils.HandleError(err, "Cannot observer loo-file property")

	err = MpvInstance.ObserveProperty(0, "loop-playlist", mpv.FORMAT_FLAG)
	utils.HandleError(err, "Cannot observer loop-playlist property")

	// start event listener
	startEventHandler()

	// create the state
	State = &PlayerState{
		false,
		nil, nil, 0,
		initialVolume,
		0.0,
		false, false, false,
		SongLyric{},
		LOOP_NONE,
	}

	registerInternalHooks()

	// start the player
	err = MpvInstance.Initialize()
	utils.HandleError(err, "Cannot initialize mpv")

	callHooks(HOOK_PLAYER_INITIALIZED, err)
}

func registerInternalHooks() {
	RegisterHook(func(params ...interface{}) {
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
			callHooks(HOOK_PLAYLIST_SONG_CHANGED)
		}
	}, HOOK_FILE_ENDED)

	RegisterHook(func(params ...interface{}) {
		State.Volume = params[0].(float64)
	}, HOOK_VOLUME_CHANGED)

	RegisterHook(func(params ...interface{}) {
		if State.Loop == LOOP_NONE {
			State.Loop = LOOP_TRACK
		} else {
			State.Loop = LOOP_NONE
		}
	}, HOOK_LOOP_TRACK_CHANGED)

	RegisterHook(func(params ...interface{}) {
		if State.Loop == LOOP_NONE {
			State.Loop = LOOP_PLAYLIST
		} else {
			State.Loop = LOOP_NONE
		}
	}, HOOK_LOOP_PLAYLIST_CHANGED)
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
	callHooks(HOOK_GENERIC_UPDATE)
}

func PlayFromYouTube(result *search.YouTubeResult, playlist *storage.Playlist) error {
	// remove all entries from playlist
	ClearPlaylist()
	// remove pause
	Play()

	State.Result = result
	State.Playlist = playlist

	if result == nil {
		var err error
		for i, song := range playlist.Songs {
			if i == 0 {
				err = LoadFile(song.URL())
			} else {
				err = AppendFile(song.URL())
			}
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return LoadFile(result.URL())
	}
}

func LoadFile(filePath string) error {
	err := MpvInstance.Command([]string{"loadfile", filePath})
	callHooks(HOOK_FILE_LOAD_STARTED, err, filePath)
	return err
}

func AppendFile(filePath string) error {
	err := MpvInstance.Command([]string{"loadfile", filePath, "append"})
	callHooks(HOOK_FILE_LOAD_STARTED, err, filePath)
	return err
}

func PlayPause() error {
	if State.Paused {
		return Play()
	} else {
		return Pause()
	}
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
	err := MpvInstance.SetProperty("time-pos", mpv.FORMAT_DOUBLE, pos)
	callHooks(HOOK_POSITION_CHANGED, err, pos)
	return err
}
