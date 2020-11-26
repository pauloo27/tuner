package player

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/lyric"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
	"github.com/godbus/dbus/v5"
)

type UpdateHandler func(result *search.YouTubeResult, mpv *MPV)
type SaveFunction func(result *search.YouTubeResult, mpv *MPV)

type MPV struct {
	Pid                                  int
	Cmd                                  *exec.Cmd
	Player                               *mpris.Player
	Playlist                             *storage.Playlist
	PlaylistIndex                        int
	ShowHelp, ShowLyric, ShowURL, Saving bool
	LyricIndex                           int
	LyricLines                           []string
	result                               *search.YouTubeResult
	onUpdate                             UpdateHandler
	save                                 SaveFunction
	Exitted                              bool
}

func ConnectToMPV(
	cmd *exec.Cmd,
	result *search.YouTubeResult,
	playlist *storage.Playlist,
	onUpdate UpdateHandler,
	save SaveFunction,
) *MPV {
	for {
		if cmd.Process != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	pid := cmd.Process.Pid

	nameWithPID := fmt.Sprintf("org.mpris.MediaPlayer2.mpv.instance%d", pid)

	conn, err := dbus.SessionBus()
	utils.HandleError(err, "Cannot connect to dbus")

	names, err := mpris.List(conn)
	utils.HandleError(err, "Cannot list players")

	playerName := ""

	for _, name := range names {
		if name == nameWithPID {
			playerName = name
			break
		}
	}

	if playerName == "" {
		playerName = "org.mpris.MediaPlayer2.mpv"
	}

	player := mpris.New(conn, playerName)
	utils.HandleError(err, "Cannot connect to mpv")

	mpv := MPV{
		pid, cmd, player,
		playlist, 0,
		false, false, false, false,
		0, []string{},
		result,
		onUpdate, save, false,
	}

	mpv.Update()

	go func() {
		ch := make(chan *dbus.Signal)
		err := mpv.Player.OnSignal(ch)
		utils.HandleError(err, "Cannot add signal handler")

		for sig := range ch {
			if mpv.Exitted {
				break
			}
			body := sig.Body[1].(map[string]dbus.Variant)
			if _, ok := body["Metadata"]; ok {
				metadata := body["Metadata"].Value().(map[string]dbus.Variant)
				rawTrackId := string(metadata["mpris:trackid"].Value().(dbus.ObjectPath))
				rawTrackId = strings.TrimPrefix(rawTrackId, "/")

				trackId, err := strconv.ParseInt(rawTrackId, 10, 32)
				utils.HandleError(err, "Cannot parse track id "+rawTrackId)

				mpv.PlaylistIndex = int(trackId)
			}

			mpv.Update()
		}
	}()

	return &mpv
}

func (i *MPV) Update() {
	i.onUpdate(i.GetPlaying(), i)
}

func (i *MPV) IsPlaylist() bool {
	return i.result == nil
}

func (i *MPV) GetPlaying() *search.YouTubeResult {
	if i.IsPlaylist() {
		return i.Playlist.Songs[i.PlaylistIndex]
	} else {
		return i.result
	}
}

func (i *MPV) Save() {
	i.save(i.GetPlaying(), i)
}

func (i *MPV) PlayPause() {
	_ = i.Player.PlayPause()
	i.Update()
}

func (i *MPV) Exit() {
	i.Exitted = true
}

func (i *MPV) Previous() error {
	err := i.Player.Previous()
	i.LyricLines = []string{}
	i.LyricIndex = 0
	return err
}

func (i *MPV) Next() error {
	err := i.Player.Next()
	i.LyricLines = []string{}
	i.LyricIndex = 0
	return err
}

func (i *MPV) FetchLyric() {
	path, err := lyric.SearchFor(fmt.Sprintf("%s %s", i.GetPlaying().Title, i.GetPlaying().Uploader))
	if err != nil {
		i.LyricLines = []string{"Cannot get lyric"}
		i.Update()
		return
	}

	l, err := lyric.Fetch(path)
	if err != nil {
		i.LyricLines = []string{"Cannot get lyric"}
		i.Update()
		return
	}

	i.LyricLines = strings.Split(l, "\n")
	i.Update()
}
