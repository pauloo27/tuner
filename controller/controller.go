package controller

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
	"github.com/godbus/dbus"
)

type UpdateHandler func(result *search.YouTubeResult, mpv *MPV)

type MPV struct {
	Pid                 int
	Player              *mpris.Player
	ShowHelp, ShowLyric bool
	LyricIndex          int
	LyricLines          []string
	Result              *search.YouTubeResult
	OnUpdate            UpdateHandler
}

func ConnectToMPV(cmd *exec.Cmd, result *search.YouTubeResult, onUpdate UpdateHandler) *MPV {
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

	mpv := MPV{pid, player, false, false, 0, []string{"Fetching lyric..."}, result, onUpdate}

	mpv.Update()

	return &mpv
}

func (i *MPV) Update() {
	i.OnUpdate(i.Result, i)
}

func (i *MPV) PlayPause() {
	_ = i.Player.PlayPause()
}
