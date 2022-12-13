package main

import (
	"github.com/Pauloo27/tuner/internal/ui"

	_ "github.com/Pauloo27/tuner/internal/ui/pages/home"
	_ "github.com/Pauloo27/tuner/internal/ui/pages/playing"
	_ "github.com/Pauloo27/tuner/internal/ui/pages/searching"

	_ "github.com/Pauloo27/tuner/internal/providers/player/mpv"
)

func main() {
	err := ui.StartApp("home")
	if err != nil {
		panic(err)
	}
}
