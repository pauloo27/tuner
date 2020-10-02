package keybind

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
	ByKey    = map[keyboard.Key]Keybind{}
	ByChar   = map[rune]Keybind{}
	keybinds []Keybind
)

func RegisterDefaultKeybinds() {
	killMpv := Keybind{
		Description: "Stop the player",
		KeyName:     "Esc",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			_ = cmd.Process.Kill()
		},
	}

	ByKey[keyboard.KeyEsc] = killMpv
	killMpv.KeyName = "CtrlC"
	ByKey[keyboard.KeyCtrlC] = killMpv

	ByKey[keyboard.KeySpace] = Keybind{
		Description: "Play/Pause song",
		KeyName:     "Space",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			mpv.PlayPause()
		},
	}

	ByChar['9'] = Keybind{
		Description: "Decrease the volume",
		KeyName:     "9",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			volume, err := mpv.Player.GetVolume()
			utils.HandleError(err, "Cannot get MPV volume")
			err = mpv.Player.SetVolume(volume - 0.05)
			utils.HandleError(err, "Cannot set MPV volume")
		},
	}

	ByChar['0'] = Keybind{
		Description: "Increase the volume",
		KeyName:     "0",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			volume, err := mpv.Player.GetVolume()
			utils.HandleError(err, "Cannot get MPV volume")
			err = mpv.Player.SetVolume(volume + 0.05)
			utils.HandleError(err, "Cannot set MPV volume")
		},
	}

	ByChar['?'] = Keybind{
		Description: "Toggle keybind list",
		KeyName:     "?",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			mpv.ShowHelp = !mpv.ShowHelp
		},
	}

	ByChar['l'] = Keybind{
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

	ByChar['p'] = Keybind{
		Description: "Toggle lyric",
		KeyName:     "P",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			mpv.ShowLyric = !mpv.ShowLyric
		},
	}

	ByChar['w'] = Keybind{
		Description: "Scroll lyric up",
		KeyName:     "W",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			if mpv.LyricIndex > 0 {
				mpv.LyricIndex = mpv.LyricIndex - 1
			}
		},
	}

	ByChar['s'] = Keybind{
		Description: "Scroll lyric down",
		KeyName:     "S",
		Handler: func(cmd *exec.Cmd, mpv *controller.MPV) {
			if mpv.LyricIndex < len(mpv.LyricLines) {
				mpv.LyricIndex = mpv.LyricIndex + 1
			}
		},
	}
}

func HandlePress(c rune, key keyboard.Key, cmd *exec.Cmd, mpv *controller.MPV) {
	if bind, ok := ByKey[key]; ok {
		bind.Handler(cmd, mpv)
	} else if bind, ok := ByChar[c]; ok {
		bind.Handler(cmd, mpv)
	}
}

func ListBinds() []Keybind {
	if keybinds != nil {
		return keybinds
	}
	keybinds = []Keybind{}
	for _, bind := range ByKey {
		keybinds = append(keybinds, bind)
	}
	for _, bind := range ByChar {
		keybinds = append(keybinds, bind)
	}
	return keybinds
}
