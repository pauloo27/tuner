package mode

import (
	"os"
	"strings"

	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
)

func simplePlayModeHandler() {
	var playing chan bool

	// wait until finished playing...
	player.RegisterHook(func(params ...interface{}) {
		if playing != nil {
			playing <- false
		}
	}, player.HookIdle)

	query := strings.Join(os.Args[2:], " ")
	results := search.Search(query, 1, search.SourceYouTube, search.SourceSoundCloud)
	player.PlaySearchResult(results[0], nil)

	playing = make(chan bool)
	<-playing
}

var SimplePlayMode = Mode{Handler: simplePlayModeHandler, Displayed: false}
