package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
)

func searchFor() {
	searchTerm := utils.AskFor("Search term")
	c := make(chan bool)
	go utils.PrintWithLoadIcon(fmt.Sprintf("Searching for %s", searchTerm), c)
	results := search.SearchYouTube(searchTerm, 10)

	c <- true
	<-c

	if len(results) == 0 {
		fmt.Printf("%sNo results found.\n", utils.ColorRed)
	}

	for index, result := range results {
		fmt.Printf(" %s-> %d: %s\n", utils.ColorGreen, index+1, result.Title)
	}
	index := utils.AskFor("Your pick ID")
	parsedIndex, err := strconv.ParseInt(index, 10, 32)
	utils.HandleError(err, "Invalid ID")
	realIndex := int(parsedIndex) - 1

	if len(results) <= realIndex || realIndex <= -1 {
		fmt.Printf("%sIndex too big\n", utils.ColorRed)
		return
	}
	result := results[realIndex]
	url := fmt.Sprintf("https://youtube.com/watch?v=%s", result.ID)
	fmt.Printf("%sPlaying %s // %s%s\n", utils.ColorGreen, result.Title, url, utils.ColorReset)

	cmd := exec.Command("mpv", url, "--no-video")

	err = cmd.Run()
	utils.HandleError(err, "Cannot run MPV")
}

func searchLocal() {
	fmt.Println("Not implemented yet")
}

func main() {
	if len(os.Args) >= 2 {
		operation := strings.ToLower(os.Args[1])
		switch operation {
		case "s":
			searchFor()
		case "l":
			searchLocal()
		default:
			searchFor()
		}
	} else {
		searchFor()
	}
}
