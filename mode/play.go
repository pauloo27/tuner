package mode

import (
	"os"
	"strings"

	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
)

func playModeHandler() {
	var playing chan bool

	// wait until finished playing...
	player.RegisterHook(func(params ...interface{}) {
		if playing != nil {
			playing <- false
		}
	}, player.HookIdle)

	query := strings.Join(os.Args[2:], " ")
	results := search.Search(query, 1, search.SourceYouTube)
	player.PlaySearchResult(results[0], nil)
	go keybind.Listen()

	playing = make(chan bool)
	<-playing
	keyboard.Close()
	utils.ShowCursor()
	utils.Exit()
}

var PlayMode = Mode{Handler: playModeHandler, Displayed: true}
