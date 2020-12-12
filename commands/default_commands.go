package commands

import (
	"fmt"
	"strings"

	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/state"
	"github.com/Pauloo27/tuner/storage"
)

func SetupDefaultCommands() {
	help := command.Command{
		Name:        "help",
		Description: "List all the commands",
		Aliases:     []string{"h"},
		Handle: func(input string) string {
			for label, cmd := range command.Commands {
				if cmd.Name == label {
					aliases := strings.Join(cmd.Aliases, " ")
					fmt.Printf("  -> /%s (or %v): %s\n", cmd.Name, aliases, cmd.Description)
				}
			}
			return ""
		},
	}

	keepLiveCache := command.Command{
		Name: "cache",
		Description: "Toggle the option to keep cache of the playing media." +
			" When true, you can seek back, but it will use more RAM " +
			"(default is false)",
		Aliases: []string{"c"},
		Handle: func(input string) string {
			state.Data.Cache = !state.Data.Cache
			storage.Save(state.Data)
			return fmt.Sprintf("Cache set to %v", state.Data.Cache)
		},
	}

	showVideo := command.Command{
		Name:        "video",
		Description: "Toggle the option to show the video (default is false)",
		Aliases:     []string{"v"},
		Handle: func(input string) string {
			state.Data.ShowVideo = !state.Data.ShowVideo
			storage.Save(state.Data)
			return fmt.Sprintf("Show video set to %v", state.Data.ShowVideo)
		},
	}

	fetchAlbum := command.Command{
		Name:        "album",
		Description: "Toggle the option to show album art (default is false)",
		Aliases:     []string{"a"},
		Handle: func(input string) string {
			state.Data.FetchAlbum = !state.Data.FetchAlbum
			storage.Save(state.Data)
			return fmt.Sprintf("Show album set to %v", state.Data.FetchAlbum)
		},
	}

	command.RegisterCommands(help, keepLiveCache, showVideo, fetchAlbum)
}
