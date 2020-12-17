package keybind

import (
	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/player"
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
	BindKey(keyboard.KeyArrowLeft, Keybind{
		Description: "Seek 5 seconds back",
		KeyName:     "Arrow Left",
		Handler: func() {
			player.Seek(-5)
		},
	})

	BindKey(keyboard.KeyArrowRight, Keybind{
		Description: "Seek 5 seconds",
		KeyName:     "Arrow Right",
		Handler: func() {
			player.Seek(+5)
		},
	})

	BindKey(keyboard.KeyCtrlC, Keybind{
		Description: "Stop the player",
		KeyName:     "Ctrl C",
		Handler: func() {
			player.Stop()
		},
	})

	BindKey(keyboard.KeySpace, Keybind{
		Description: "Play/Pause song",
		KeyName:     "Space",
		Handler: func() {
			player.PlayPause()
		},
	})

	BindKey(keyboard.KeyArrowDown, Keybind{
		Description: "Decrease the volume",
		KeyName:     "Arrow Down",
		Handler: func() {
			player.SetVolume(player.State.Volume - 5.0)
		},
	})

	BindKey(keyboard.KeyArrowUp, Keybind{
		Description: "Increase the volume",
		KeyName:     "Arrow Up",
		Handler: func() {
			player.SetVolume(player.State.Volume + 5.0)
		},
	})

	BindChar('?', Keybind{
		Description: "Toggle keybind list",
		KeyName:     "?",
		Handler: func() {
			player.State.ShowHelp = !player.State.ShowHelp
			player.ForceUpdate()
		},
	})

	BindChar('l', Keybind{
		Description: "Toggle loop",
		KeyName:     "L",
		Handler: func() {
			loop := player.State.Loop

			if loop == player.LOOP_NONE {
				player.LoopTrack()
				return
			} else if loop == player.LOOP_TRACK && player.State.IsPlaylist() {
				player.LoopPlaylist()
				return
			}

			player.LoopNone()
		},
	})

	BindChar('p', Keybind{
		Description: "Toggle lyric",
		KeyName:     "P",
		Handler: func() {
			if len(player.State.Lyric.Lines) == 0 {
				go player.FetchLyric()
			}
			player.State.ShowLyric = !player.State.ShowLyric
			player.ForceUpdate()
		},
	})

	BindChar('w', Keybind{
		Description: "Scroll lyric up",
		KeyName:     "W",
		Handler: func() {
			if player.State.Lyric.Index > 0 {
				player.State.Lyric.Index--
				player.ForceUpdate()
			}
		},
	})

	BindChar('s', Keybind{
		Description: "Scroll lyric down",
		KeyName:     "S",
		Handler: func() {
			if player.State.Lyric.Index < len(player.State.Lyric.Lines) {
				player.State.Lyric.Index++
				player.ForceUpdate()
			}
		},
	})

	BindChar('u', Keybind{
		Description: "Show video URL",
		KeyName:     "U",
		Handler: func() {
			player.State.ShowURL = !player.State.ShowURL
			player.ForceUpdate()
		},
	})

	BindChar('b', Keybind{
		Description: "Save song to playlist",
		KeyName:     "B",
		Handler: func() {
			player.SaveToPlaylist()
		},
	})

	BindChar('>', Keybind{
		Description: "Next song in playlist",
		KeyName:     ">",
		Handler: func() {
			player.PlaylistNext()
		},
	})

	BindChar('<', Keybind{
		Description: "Previous song in playlist",
		KeyName:     "<",
		Handler: func() {
			player.PlaylistPrevious()
		},
	})
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
			if player.State.Idle {
				break
			}
		} else {
			HandlePress(c, key)
		}
	}
}
