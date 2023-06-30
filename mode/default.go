package mode

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/commands"
	"github.com/Pauloo27/tuner/display"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
	"github.com/Pauloo27/tuner/version"
)

var (
	playing chan bool
	warning string
	searchSources = []search.SearchSource{search.SourceYouTube}
)

func play(result *search.SearchResult, playlist *storage.Playlist) {
	player.PlaySearchResult(result, playlist)
	go keybind.Listen()
	// wait to the player to exit
	playing = make(chan bool)
	<-playing
	keyboard.Close()
	utils.ShowCursor()
}

func promptEntry() {
	utils.ClearScreen()
	utils.ShowCursor()

	fmt.Printf("%sTuner %s\n", utils.ColorGreen, version.Current)
	fmt.Printf("%sPlaylists:\n", utils.ColorBlue)
	display.ListPlaylists()
	fmt.Printf("%sUse #<id> to start a playlist%s\n", utils.ColorBlue, utils.ColorReset)

	if warning != "" {
		fmt.Printf("%s%s%s\n", utils.ColorYellow, warning, utils.ColorReset)
		warning = ""
	}

	fmt.Println()
	rawInput, err := utils.AskFor("Search")
	if err != nil {
		utils.Exit()
	}

	if rawInput == "" {
		warning = "Missing search term"
		return
	}

	prefix := rawInput[0]
	unprefixed := rawInput[1:]

	searchLimit := 10

	switch prefix {
	case '/':
		found, msg := command.InvokeCommand(unprefixed)
		if found {
			warning = msg
		} else {
			warning = "Command not found"
		}
		return
	case '!':
		searchLimit = 1
		rawInput = unprefixed
	case '#':
		rawInput = unprefixed
		index, err := strconv.Atoi(rawInput)
		if err != nil || index <= 0 || index > len(player.State.Data.Playlists) {
			warning = "Invalid playlist"
			return
		}
		play(nil, player.State.Data.Playlists[index-1])
		return
	}

	// 'loading' message
	c := make(chan bool)
	go utils.PrintWithLoadIcon(utils.Fmt("Searching for %s", rawInput), c, 100*time.Millisecond, true)
	// do search
	results := search.Search(rawInput, searchLimit, searchSources...)

	// ask the loading message to stop
	c <- true
	// wait until it stopped
	<-c

	if len(results) == 0 {
		warning = "No results found for " + rawInput
		return
	}

	if searchLimit == 1 {
		play(results[0], nil)
		return
	}

	display.ListResults(results)
	index, err := utils.AskForInt("Insert index of the video")
	if err != nil {
		warning = "Invalid input"
		return
	}

	if index <= 0 || index > len(results) {
		warning = "Invalid index"
		return
	}
	index--
	play(results[index], nil)
}

func defaultModeHandler() {
	// 'lock' for the prompt
	player.RegisterHook(func(params ...interface{}) {
		if playing != nil {
			playing <- false
		}
	}, player.HookIdle)

	commands.SetupDefaultCommands()
	// handle sigterm (Ctrl+C)
	utils.OnSigTerm(func(sig *os.Signal) {
		utils.Exit()
	})

	for {
		promptEntry()
	}
}

var DefaultMode = Mode{Handler: defaultModeHandler, Displayed: true}
