package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/Pauloo27/keyboard"
	"github.com/Pauloo27/tuner/album"
	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/commands"
	"github.com/Pauloo27/tuner/display"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/storage"
	"github.com/Pauloo27/tuner/utils"
)

func parametersFor(result *search.YouTubeResult,
	playlist *storage.Playlist) (parameters []string) {
	if result == nil {
		for _, song := range playlist.Songs {
			parameters = append(parameters, song.URL())
		}
	} else {
		parameters = append(parameters, result.URL())
	}

	if !state.Data.ShowVideo {
		parameters = append(parameters, "--no-video", "--ytdl-format=worst")
	}
	if !state.Data.Cache {
		parameters = append(parameters, "--cache=no")
	}
	return
}

func play(result *search.YouTubeResult, playlist *storage.Playlist) {
	state.Playing = true
	cmd := exec.Command("mpv", parametersFor(result, playlist)...)

	go func() {
		if state.MPVInstance != nil && !state.MPVInstance.Exitted {
			state.MPVInstance.Exit()
		}
		state.MPVInstance = player.ConnectToMPV(cmd, result, playlist,
			display.ShowPlaying, display.SaveToPlaylist, album.FetchAlbum,
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
