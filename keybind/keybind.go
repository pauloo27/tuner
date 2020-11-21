package keybind

import (
	"fmt"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
	"github.com/eiannone/keyboard"
)

type Keybind struct {
	Handler              func(mpv *player.MPV)
	KeyName, Description string
}

var (
	ByKey    = map[keyboard.Key]Keybind{}
	ByChar   = map[rune]Keybind{}
	keybinds []Keybind
)

func RegisterDefaultKeybinds(data *storage.TunerData) {
	ByKey[keyboard.KeyCtrlC] = Keybind{
		Description: "Stop the player",
		KeyName:     "Ctrl C",
		Handler: func(mpv *player.MPV) {
			_ = mpv.Cmd.Process.Kill()
		},
	}

	ByKey[keyboard.KeySpace] = Keybind{
		Description: "Play/Pause song",
		KeyName:     "Space",
		Handler: func(mpv *player.MPV) {
			mpv.PlayPause()
		},
	}

	ByChar['9'] = Keybind{
		Description: "Decrease the volume",
		KeyName:     "9",
		Handler: func(mpv *player.MPV) {
			volume, err := mpv.Player.GetVolume()
			// avoid crashing when the player is starting
			if err != nil {
				fmt.Println("Cannot get MPV volume")
				return
			}
			err = mpv.Player.SetVolume(volume - 0.05)
			utils.HandleError(err, "Cannot set MPV volume")
		},
	}

	ByChar['0'] = Keybind{
		Description: "Increase the volume",
		KeyName:     "0",
		Handler: func(mpv *player.MPV) {
			volume, err := mpv.Player.GetVolume()
			// avoid crashing when the player is starting
			if err != nil {
				fmt.Println("Cannot get MPV volume")
				return
			}
			err = mpv.Player.SetVolume(volume + 0.05)
			utils.HandleError(err, "Cannot set MPV volume")
		},
	}

	ByChar['?'] = Keybind{
		Description: "Toggle keybind list",
		KeyName:     "?",
		Handler: func(mpv *player.MPV) {
			mpv.ShowHelp = !mpv.ShowHelp
			mpv.Update()
		},
	}

	ByChar['l'] = Keybind{
		Description: "Toggle loop",
		KeyName:     "L",
		Handler: func(mpv *player.MPV) {
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
		Handler: func(mpv *player.MPV) {
			if len(mpv.LyricLines) == 0 {
				go mpv.FetchLyric()
			}
			mpv.ShowLyric = !mpv.ShowLyric
			mpv.Update()
		},
	}

	ByChar['w'] = Keybind{
		Description: "Scroll lyric up",
		KeyName:     "W",
		Handler: func(mpv *player.MPV) {
			if mpv.LyricIndex > 0 {
				mpv.LyricIndex = mpv.LyricIndex - 1
				mpv.Update()
			}
		},
	}

	ByChar['s'] = Keybind{
		Description: "Scroll lyric down",
		KeyName:     "S",
		Handler: func(mpv *player.MPV) {
			if mpv.LyricIndex < len(mpv.LyricLines) {
				mpv.LyricIndex = mpv.LyricIndex + 1
				mpv.Update()
			}
		},
	}

	ByChar['u'] = Keybind{
		Description: "Show video URL",
		KeyName:     "U",
		Handler: func(mpv *player.MPV) {
			mpv.ShowURL = !mpv.ShowURL
			mpv.Update()
		},
	}

	ByChar['b'] = Keybind{
		Description: "Save song to playlist",
		KeyName:     "B",
		Handler: func(mpv *player.MPV) {
			mpv.Saving = !mpv.Saving
			mpv.Update()
			if mpv.Saving {
				mpv.Save()
			}
		},
	}

	ByChar['>'] = Keybind{
		Description: "Next song in playlist",
		KeyName:     ">",
		Handler: func(mpv *player.MPV) {
			err := mpv.Player.Next()
			utils.HandleError(err, "Cannot skip song")
		},
	}

	ByChar['<'] = Keybind{
		Description: "Previous song in playlist",
		KeyName:     "<",
		Handler: func(mpv *player.MPV) {
			err := mpv.Player.Previous()
			utils.HandleError(err, "Cannot skip song")
		},
	}
}

func HandlePress(c rune, key keyboard.Key, mpv *player.MPV) {
	if bind, ok := ByKey[key]; ok {
		bind.Handler(mpv)
	} else if bind, ok := ByChar[c]; ok {
		bind.Handler(mpv)
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
