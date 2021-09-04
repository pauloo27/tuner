package main

import (
	"github.com/Pauloo27/tuner/ui"

	// import all pages
	_ "github.com/Pauloo27/tuner/ui/pages/home"
)

func main() {
	err := ui.StartApp("home")
	if err != nil {
		panic(err)
	}
}
