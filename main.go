package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/commands"
	"github.com/Pauloo27/tuner/display"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
	"github.com/eiannone/keyboard"
)

func play(result *search.YouTubeResult, playlist *storage.Playlist) {
	state.Playing = true
	cmd := exec.Command("mpv", player.ParametersFor(result, playlist)...)

	go func() {
		if state.MPVInstance != nil && !state.MPVInstance.Exitted {
			state.MPVInstance.Exit()
		}
		state.MPVInstance = player.ConnectToMPV(cmd, result, playlist,
			display.ShowPlaying, display.SaveToPlaylist,
		)
		go keybind.Listen()
	}()

	err := cmd.Run()

	if err != nil && err.Error() != "exit status 4" && err.Error() != "signal: killed" {
		utils.HandleError(err, "Cannot run MPV")
	}

	keyboard.Close()
	state.Playing = false
	state.MPVInstance.Exit()
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
		playlistIndex, err := strconv.Atoi(rawInput)
		if err != nil {
			state.Warning = "Invalid playlist"
			return
		}
		play(nil, state.Data.Playlists[playlistIndex-1])
		return
	}

	// search
	c := make(chan bool)

	// 'loading' message
	go utils.PrintWithLoadIcon(utils.Fmt("Searching for %s", rawInput), c, 100*time.Millisecond, true)
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
	utils.HandleError(err, "Cannot read user input")
	if index <= 0 || index > len(results) {
		state.Warning = "Invalid index"
		return
	}
	index--
	play(results[index], nil)
}

func main() {
	state.Start()
	commands.SetupDefaultCommands()
	keybind.RegisterDefaultKeybinds(state.Data)
	// handle sigterm (Ctrl+C)
	utils.OnSigTerm(func(sig *os.Signal) {
		if !state.Playing {
			utils.ClearScreen()
			fmt.Println("Bye!")
			os.Exit(0)
		}
	})

	// loop
	for {
		promptEntry()
	}
}
