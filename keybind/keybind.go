package keybind

import (
	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/new_player"
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
			new_player.Seek(-5)
		},
	})

	BindKey(keyboard.KeyArrowRight, Keybind{
		Description: "Seek 5 seconds",
		KeyName:     "Arrow Right",
		Handler: func() {
			new_player.Seek(+5)
		},
	})

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
			new_player.SetVolume(new_player.State.Volume - 5.0)
		},
	})

	BindKey(keyboard.KeyArrowUp, Keybind{
		Description: "Increase the volume",
		KeyName:     "Arrow Up",
		Handler: func() {
			new_player.SetVolume(new_player.State.Volume + 5.0)
		},
	})

	BindChar('?', Keybind{
		Description: "Toggle keybind list",
		KeyName:     "?",
		Handler: func() {
			new_player.State.ShowHelp = !new_player.State.ShowHelp
			new_player.ForceUpdate()
		},
	})

	BindChar('l', Keybind{
		Description: "Toggle loop",
		KeyName:     "L",
		Handler: func() {
			loop := new_player.State.Loop

			if loop == new_player.LOOP_NONE {
				new_player.LoopTrack()
				return
			} else if loop == new_player.LOOP_TRACK && new_player.State.IsPlaylist() {
				new_player.LoopPlaylist()
				return
			}

			new_player.LoopNone()
		},
	})

	BindChar('p', Keybind{
		Description: "Toggle lyric",
		KeyName:     "P",
		Handler: func() {
			if len(new_player.State.Lyric.Lines) == 0 {
				go new_player.FetchLyric()
			}
			new_player.State.ShowLyric = !new_player.State.ShowLyric
			new_player.ForceUpdate()
		},
	})

	BindChar('w', Keybind{
		Description: "Scroll lyric up",
		KeyName:     "W",
		Handler: func() {
			if new_player.State.Lyric.Index > 0 {
				new_player.State.Lyric.Index--
				new_player.ForceUpdate()
			}
		},
	})

	BindChar('s', Keybind{
		Description: "Scroll lyric down",
		KeyName:     "S",
		Handler: func() {
			if new_player.State.Lyric.Index < len(new_player.State.Lyric.Lines) {
				new_player.State.Lyric.Index++
				new_player.ForceUpdate()
			}
		},
	})

	BindChar('u', Keybind{
		Description: "Show video URL",
		KeyName:     "U",
		Handler: func() {
			new_player.State.ShowURL = !new_player.State.ShowURL
			new_player.ForceUpdate()
		},
	})

	BindChar('b', Keybind{
		Description: "Save song to playlist",
		KeyName:     "B",
		Handler: func() {
			new_player.SaveToPlaylist()
		},
	})

	BindChar('>', Keybind{
		Description: "Next song in playlist",
		KeyName:     ">",
		Handler: func() {
			new_player.PlaylistNext()
		},
	})

	BindChar('<', Keybind{
		Description: "Previous song in playlist",
		KeyName:     "<",
		Handler: func() {
			new_player.PlaylistPrevious()
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
			break
		} else {
			HandlePress(c, key)
		}
	}
}
