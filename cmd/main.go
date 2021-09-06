package main

import (
	"github.com/Pauloo27/tuner/internal/ui"

	// import all pages
	_ "github.com/Pauloo27/tuner/internal/ui/pages/home"
	_ "github.com/Pauloo27/tuner/internal/ui/pages/playing"
	_ "github.com/Pauloo27/tuner/internal/ui/pages/searching"
	// import providers
	_ "github.com/Pauloo27/tuner/internal/providers/player"
)

func main() {
	err := ui.StartApp("home")
	if err != nil {
		panic(err)
	}
}
