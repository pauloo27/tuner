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
	rawSearchTerm := utils.AskFor("Search term (add ! prefix to play the first)")
	if rawSearchTerm == ":q" {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	searchTerm := strings.TrimPrefix(rawSearchTerm, "!")
	c := make(chan bool)
	go utils.PrintWithLoadIcon(fmt.Sprintf("Searching for %s", searchTerm), c, 100*time.Millisecond)
	results := search.SearchYouTube(searchTerm, 10)

	c <- true
	<-c

	if len(results) == 0 {
		fmt.Printf("%sNo results found.\n", utils.ColorRed)
		return
	}

	realIndex := 0
	if !strings.HasPrefix(rawSearchTerm, "!") {
		for index, result := range results {
			fmt.Printf(" %s-> %d: %s\n", utils.ColorGreen, index+1, result.Title)
		}
		index := utils.AskFor("Your pick ID")
		parsedIndex, err := strconv.ParseInt(index, 10, 32)
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
	go utils.PrintWithLoadIcon(fmt.Sprintf("%sPlaying %s // %s%s", utils.ColorGreen, result.Title, url, utils.ColorReset), c, 1000*time.Millisecond)

	cmd := exec.Command("mpv", url, "--no-video")

	err := cmd.Run()
	if err != nil {
		if err.Error() == "exit status 4" {
			c <- true
			return
		}
		utils.HandleError(err, "Cannot run MPV")
	}
	c <- true
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
	fmt.Println(utils.ColorGreen, "Tuner - To exit, search for `:q`")
	for {
		searchFor()
	}
}
