package main

import (
	"fmt"
	"os"

	"github.com/Pauloo27/tuner/album"
	"github.com/Pauloo27/tuner/display"
	"github.com/Pauloo27/tuner/img"
	"github.com/Pauloo27/tuner/integrations"
	"github.com/Pauloo27/tuner/keybind"
	"github.com/Pauloo27/tuner/mode"
	"github.com/Pauloo27/tuner/mpris"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/version"
)

const VERSION = "0.0.3-dev"

func exitWithInvalidUsage() {
	fmt.Println("Invalid mode. Valid modes are:")
	for mode := range modes {
		fmt.Println(mode)
	}
	os.Exit(-1)
}

var modes = map[string]mode.Mode{
	"play":        mode.PlayMode,
	"p":           mode.PlayMode,
	"simple-play": mode.SimplePlayMode,
	"sp":          mode.SimplePlayMode,
	"default":     mode.DefaultMode,
}

func main() {
	var mode string

	if len(os.Args) == 1 {
		mode = "default"
	} else {
		mode = os.Args[1]
	}

	selectedMode, ok := modes[mode]
	if !ok {
		exitWithInvalidUsage()
	}

	player.Initialize()

	if player.State.Data.ShowInDiscord {
		integrations.ConnectToDiscord()
	}

	// load mpv-mpris
	mpris.LoadScript()

	version.Migrate(VERSION)
	if selectedMode.Displayed {
		keybind.RegisterDefaultKeybinds()
		display.RegisterHooks()
		album.RegisterHooks()
		if player.State.Data.FetchAlbum {
			img.StartDaemon()
		}
	}

	selectedMode.Handler()
}
