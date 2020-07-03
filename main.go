package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Pauloo27/tuner/search"
	"github.com/Pauloo27/tuner/utils"
)

func searchFor() {
	searchTerm := utils.AskFor("Search term")
	c := make(chan bool)
	go utils.PrintWithLoadIcon(fmt.Sprintf("Searching for %s", searchTerm), c)
	search.SearchYouTube(searchTerm)
	c <- true
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
