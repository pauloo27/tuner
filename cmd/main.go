package main

import (
	"github.com/Pauloo27/tuner/internal/ui"

	// import all pages
	_ "github.com/Pauloo27/tuner/internal/ui/pages/home"
)

func main() {
	err := ui.StartApp("home")
	if err != nil {
		panic(err)
	}
}
