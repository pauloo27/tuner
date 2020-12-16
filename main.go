package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/commands"
	"github.com/Pauloo27/tuner/display"
	"github.com/Pauloo27/tuner/img"
	"github.com/Pauloo27/tuner/new_display"
	"github.com/Pauloo27/tuner/new_keybind"
	"github.com/Pauloo27/tuner/new_player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

var (
	playing = make(chan bool)
)

func play(result *search.YouTubeResult, playlist *storage.Playlist) {
	if result == nil {
		fmt.Println("Not supported yet (1)")
		os.Exit(-1)
	} else {
		new_player.PlayFromYouTube(result)
	}
	go new_keybind.Listen()
	// wait to the player to exit
	<-playing
	keyboard.Close()
}

func promptEntry() {
	utils.ClearScreen()

	fmt.Printf("%sPlaylists:\n", utils.ColorBlue)
	display.ListPlaylists()
	fmt.Printf("%sUse #<id> to start a playlist%s\n", utils.ColorBlue, utils.ColorReset)

	if state.Warning != "" {
		fmt.Printf("%s%s%s\n", utils.ColorYellow, state.Warning, utils.ColorReset)
		state.Warning = ""
	}

	fmt.Println()
	rawInput, err := utils.AskFor("Search")
	utils.HandleError(err, "Cannot read user input")

	if rawInput == "" {
		state.Warning = "Missing search term"
		return
	}

	prefix := rawInput[0]
	unprefixed := rawInput[1:]

	searchLimit := 10

	switch prefix {
	case '/':
		found, msg := command.InvokeCommand(unprefixed)
		if found {
			state.Warning = msg
		} else {
			state.Warning = "Command not found"
		}
		return
	case '!':
		searchLimit = 1
		rawInput = unprefixed
	case '#':
		rawInput = unprefixed
		index, err := strconv.Atoi(rawInput)
		if err != nil || index <= 0 || index > len(state.Data.Playlists) {
			state.Warning = "Invalid playlist"
			return
		}
		play(nil, state.Data.Playlists[index-1])
		return
	}

	// 'loading' message
	c := make(chan bool)
	go utils.PrintWithLoadIcon(utils.Fmt("Searching for %s", rawInput), c, 100*time.Millisecond, true)
	// do search
	results := search.SearchYouTube(rawInput, searchLimit)

	// ask the loading message to stop
	c <- true
	// wait until it stopped
	<-c

	if len(results) == 0 {
		state.Warning = "No results found for " + rawInput
		return
	}

	if searchLimit == 1 {
		play(results[0], nil)
		return
	}

	display.ListResults(results)
	index, err := utils.AskForInt("Insert index of the video")
	if err != nil {
		state.Warning = "Invalid input"
		return
	}

	if index <= 0 || index > len(results) {
		state.Warning = "Invalid index"
		return
	}
	index--
	play(results[index], nil)
}

func main() {
	state.Start()
	new_player.Initialize()

	new_player.RegisterHook(func(params ...interface{}) {
		playing <- false
	}, new_player.HOOK_FILE_ENDED)

	new_keybind.RegisterDefaultKeybinds()

	new_display.StartPlayerDisplay()

	commands.SetupDefaultCommands()
	// handle sigterm (Ctrl+C)
	utils.OnSigTerm(func(sig *os.Signal) {
		if !state.Playing {
			utils.ClearScreen()
			fmt.Println("Bye!")
			os.Exit(0)
		}
	})
	if state.Data.FetchAlbum {
		img.StartDaemon()
	}
	// loop
	for {
		promptEntry()
	}
}
