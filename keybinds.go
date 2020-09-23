package main

import (
	"os/exec"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/controller"
	"github.com/Pauloo27/tuner/utils"
	"github.com/eiannone/keyboard"
)

type Keybind struct {
	Handler              func(cmd *exec.Cmd, mpv *controller.MPV)
	KeyName, Description string
}

var (
	byKey  = map[keyboard.Key]Keybind{}
	byChar = map[rune]Keybind{}
)

func registerDefaultKeybinds() {
	killMpv := Keybind{
		Description: "Stop the player",
		KeyName:     "Esc",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			_ = cmd.Process.Kill()
		},
	}

	byKey[keyboard.KeyEsc] = killMpv
	killMpv.KeyName = "CtrlC"
	byKey[keyboard.KeyCtrlC] = killMpv

	byKey[keyboard.KeySpace] = Keybind{
		Description: "Play/Pause song",
		KeyName:     "Space",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			mpv.PlayPause()
		},
	}

	byChar['9'] = Keybind{
		Description: "Decrease the volume",
		KeyName:     "9",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			volume, err := mpv.Player.GetVolume()
			utils.HandleError(err, "Cannot get MPV volume")
			mpv.Player.SetVolume(volume - 0.05)
		},
	}

	byChar['0'] = Keybind{
		Description: "Increase the volume",
		KeyName:     "0",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			volume, err := mpv.Player.GetVolume()
			utils.HandleError(err, "Cannot get MPV volume")
			mpv.Player.SetVolume(volume + 0.05)
		},
	}

	byChar['h'] = Keybind{
		Description: "Toggle keybind list",
		KeyName:     "H",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			mpv.ShowHelp = !mpv.ShowHelp
		},
	}

	byChar['l'] = Keybind{
		Description: "Toggle loop",
		KeyName:     "L",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			loop, err := mpv.Player.GetLoopStatus()
			utils.HandleError(err, "Cannot get mpv loop status")
			newLoopStatus := mpris.LoopNone
			if loop == mpris.LoopNone {
				newLoopStatus = mpris.LoopTrack
			}
			err = mpv.Player.SetLoopStatus(newLoopStatus)
			utils.HandleError(err, "Cannot set loop status")
		},
	}
}
