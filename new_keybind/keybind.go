package new_keybind

import (
	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/new_player"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/utils"
)

type Keybind struct {
	Handler              func()
	KeyName, Description string
}

var (
	ByKey    = map[keyboard.Key]Keybind{}
	ByChar   = map[rune]Keybind{}
	keybinds []*Keybind
)

func BindKey(key keyboard.Key, bind Keybind) {
	ByKey[key] = bind
	keybinds = append(keybinds, &bind)
}

func BindChar(c rune, bind Keybind) {
	ByChar[c] = bind
	keybinds = append(keybinds, &bind)
}

func RegisterDefaultKeybinds() {
	/*
		BindKey(keyboard.KeyArrowLeft, Keybind{
			Description: "Seek 5 seconds back",
			KeyName:     "Arrow Left",
			Handler: func(mpv *player.MPV) {
				_ = mpv.Player.Seek(-5)
				mpv.Update()
			},
		})

		BindKey(keyboard.KeyArrowRight, Keybind{
			Description: "Seek 5 seconds",
			KeyName:     "Arrow Right",
			Handler: func(mpv *player.MPV) {
				_ = mpv.Player.Seek(+5)
				mpv.Update()
			},
		})
	*/

	BindKey(keyboard.KeyCtrlC, Keybind{
		Description: "Stop the player",
		KeyName:     "Ctrl C",
		Handler: func() {
			new_player.Stop()
		},
	})

	BindKey(keyboard.KeySpace, Keybind{
		Description: "Play/Pause song",
		KeyName:     "Space",
		Handler: func() {
			new_player.PlayPause()
		},
	})

	BindKey(keyboard.KeyArrowDown, Keybind{
		Description: "Decrease the volume",
		KeyName:     "Arrow Down",
		Handler: func() {
			new_player.SetVolume(new_player.State.Volume - 0.05)
		},
	})

	BindKey(keyboard.KeyArrowUp, Keybind{
		Description: "Increase the volume",
		KeyName:     "Arrow Up",
		Handler: func() {
			new_player.SetVolume(new_player.State.Volume + 0.05)
		},
	})

	/*
		BindChar('?', Keybind{
			Description: "Toggle keybind list",
			KeyName:     "?",
			Handler: func(mpv *player.MPV) {
				mpv.ShowHelp = !mpv.ShowHelp
				mpv.Update()
			},
		})
	*/

	/*
			BindChar('l', Keybind{
				Description: "Toggle loop",
				KeyName:     "L",
				Handler: func(mpv *player.MPV) {
					loop, err := mpv.Player.GetLoopStatus()
					// avoid crashing when the player is starting
					if err != nil {
						fmt.Println("Cannot get MPV loop status")
						return
					}
					newLoop := mpris.LoopNone

					if loop == mpris.LoopNone {
						newLoop = mpris.LoopTrack
					} else if loop == mpris.LoopTrack && mpv.IsPlaylist() {
						newLoop = mpris.LoopPlaylist
					}

					err = mpv.Player.SetLoopStatus(newLoop)
					utils.HandleError(err, "Cannot set loop status")
				},
			})

			BindChar('p', Keybind{
				Description: "Toggle lyric",
				KeyName:     "P",
				Handler: func(mpv *player.MPV) {
					if len(mpv.LyricLines) == 0 {
						go mpv.FetchLyric()
					}
					mpv.ShowLyric = !mpv.ShowLyric
					mpv.Update()
				},
			})

		BindChar('w', Keybind{
			Description: "Scroll lyric up",
			KeyName:     "W",
			Handler: func(mpv *player.MPV) {
				if mpv.LyricIndex > 0 {
					mpv.LyricIndex = mpv.LyricIndex - 1
					mpv.Update()
				}
			},
		})

		BindChar('s', Keybind{
			Description: "Scroll lyric down",
			KeyName:     "S",
			Handler: func(mpv *player.MPV) {
				if mpv.LyricIndex < len(mpv.LyricLines) {
					mpv.LyricIndex = mpv.LyricIndex + 1
					mpv.Update()
				}
			},
		})

		BindChar('u', Keybind{
			Description: "Show video URL",
			KeyName:     "U",
			Handler: func(mpv *player.MPV) {
				mpv.ShowURL = !mpv.ShowURL
				mpv.Update()
			},
		})

		BindChar('b', Keybind{
			Description: "Save song to playlist",
			KeyName:     "B",
			Handler: func(mpv *player.MPV) {
				mpv.Saving = !mpv.Saving
				mpv.Update()
				if mpv.Saving {
					mpv.Save()
				}
			},
		})

		BindChar('>', Keybind{
			Description: "Next song in playlist",
			KeyName:     ">",
			Handler: func(mpv *player.MPV) {
				mpv.Next()
			},
		})

		BindChar('<', Keybind{
			Description: "Previous song in playlist",
			KeyName:     "<",
			Handler: func(mpv *player.MPV) {
				mpv.Previous()
			},
		})
	*/
}

func HandlePress(c rune, key keyboard.Key) {
	if bind, ok := ByKey[key]; ok {
		bind.Handler()
	} else if bind, ok := ByChar[c]; ok {
		bind.Handler()
	}
}

func ListBinds() []*Keybind {
	return keybinds
}

func Listen() {
	err := keyboard.Open()
	utils.HandleError(err, "Cannot open keyboard")
	for {
		c, key, err := keyboard.GetKey()
		if err != nil {
			if !state.Playing {
				break
			}
		} else {
			HandlePress(c, key)
		}
	}
}
