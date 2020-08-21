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

	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
)

var close = make(chan os.Signal)

func searchFor() {
	rawSearchTerm, err := utils.AskFor("Search term (add ! prefix to play the first, :q or Ctrl+D to exit)")

	if err != nil {
		os.Exit(0)
	}

	if rawSearchTerm == ":q" {
		fmt.Println("Bye!")
		os.Exit(0)
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
		fmt.Printf("%sNo results found.\n", utils.ColorRed)
		return
	}

	realIndex := 0
	if !strings.HasPrefix(rawSearchTerm, "!") {
		for index, result := range results {
			fmt.Printf(" %s-> %d: %s\n", utils.ColorGreen, index+1, result.Title)
		}
		index, err := utils.AskFor("Your pick ID")

		if err != nil {
			os.Exit(0)
		}

		parsedIndex, err := strconv.ParseInt(index, 10, 32)

		utils.MoveCursorTo(1, 1)
		utils.ClearScreen()

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		realIndex = int(parsedIndex) - 1

		if len(results) <= realIndex || realIndex <= -1 {
			fmt.Printf("%sIndex too big\n", utils.ColorRed)
			return
		}
	}

	result := results[realIndex]
	url := fmt.Sprintf("https://youtube.com/watch?v=%s", result.ID)
	go utils.PrintWithLoadIcon(fmt.Sprintf("%sPlaying %s%s", utils.ColorGreen, result.Title, utils.ColorReset), c, 1000*time.Millisecond, true)

	cmd := exec.Command("mpv", url, "--no-video")

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
			fmt.Println("\n\nTo close search for `:q`")
		}
	}()
}

func main() {
	setupCloseHandler()
	for {
		utils.MoveCursorTo(1, 1)
		utils.ClearScreen()
		searchFor()
	}
}
