package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/commands"
	"github.com/Pauloo27/tuner/options"
	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
)

var close = make(chan os.Signal)
var playing = false
var warning = ""

func searchFor() {
	if warning != "" {
		fmt.Printf("%s%s\n", utils.ColorYellow, warning)
		warning = ""
	}
	rawSearchTerm, err := utils.AskFor("Search term (add ! prefix to play the first, Ctrl+C to exit)")

	if err != nil {
		os.Exit(0)
	}

	if strings.HasPrefix(rawSearchTerm, "/") {
		found, out := command.InvokeCommand(strings.TrimPrefix(rawSearchTerm, "/"))
		if !found {
			warning = "Invalid command"
		} else {
			warning = out
		}
		return
	}

	searchTerm := strings.TrimPrefix(rawSearchTerm, "!")
	c := make(chan bool)
	go utils.PrintWithLoadIcon(fmt.Sprintf("Searching for %s", searchTerm), c, 100*time.Millisecond, true)
	results := search.SearchYouTube(searchTerm, 10)

	c <- true
	<-c

	utils.MoveCursorTo(1, 1)
	utils.ClearScreen()

	if len(results) == 0 {
		warning = "No results found"
		return
	}

	realIndex := 0
	if !strings.HasPrefix(rawSearchTerm, "!") {
		for index, result := range results {
			durationDisplay := result.Duration
			if result.Live {
				durationDisplay = utils.ColorRed + "LIVE"
			}

			bold := ""
			if index%2 == 0 {
				bold = utils.ColorBold
			}

			fmt.Printf(
				" %s-> %d: %s%s%s from %s%s%s %s%s\n",
				utils.ColorReset+bold,
				index+1,
				utils.ColorGreen+bold,
				result.Title,
				utils.ColorReset+bold,
				utils.ColorGreen+bold,
				result.Uploader,
				utils.ColorReset+bold,
				durationDisplay,
				utils.ColorReset,
			)
		}
		index, err := utils.AskFor("Your pick ID")

		if err != nil {
			os.Exit(0)
		}

		parsedIndex, err := strconv.ParseInt(index, 10, 32)

		utils.MoveCursorTo(1, 1)
		utils.ClearScreen()

		if err != nil {
			warning = "Invalid ID"
			return
		}
		realIndex = int(parsedIndex) - 1

		if len(results) <= realIndex || realIndex <= -1 {
			warning = "Invalid ID"
			return
		}
	}

	result := results[realIndex]
	url := fmt.Sprintf("https://youtube.com/watch?v=%s", result.ID)
	go utils.PrintWithLoadIcon(fmt.Sprintf("%sPlaying %s%s", utils.ColorGreen, result.Title, utils.ColorReset), c, 1000*time.Millisecond, true)

	playing = true
	parameters := []string{url}
	if !options.Options.ShowVideo {
		parameters = append(parameters, "--no-video")
	}
	if !options.Options.Cache {
		parameters = append(parameters, "--cache=no")
	}
	cmd := exec.Command("mpv", parameters...)

	err = cmd.Run()
	if err != nil {
		if err.Error() == "exit status 4" {
			c <- true
			<-c
			return
		}
		utils.HandleError(err, "Cannot run MPV")
	}
	c <- true
	<-c
}

func setupCloseHandler() {
	signal.Notify(close, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			<-close
			if !playing {
				utils.MoveCursorTo(1, 1)
				utils.ClearScreen()
				fmt.Println("Bye!")
				os.Exit(0)
			}
		}
	}()
}

func main() {
	commands.SetupDefaultCommands()
	setupCloseHandler()
	for {
		utils.MoveCursorTo(1, 1)
		utils.ClearScreen()
		fmt.Printf("%sUse /help to see the commands%s\n\n", utils.ColorBlue, utils.ColorReset)
		searchFor()
		playing = false
	}
}
