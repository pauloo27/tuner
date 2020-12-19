package commands

import (
	"fmt"
	"strings"

	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/player"
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
			player.State.Data.Cache = !player.State.Data.Cache
			storage.Save(player.State.Data)
			return fmt.Sprintf("Cache set to %v", player.State.Data.Cache)
		},
	}

	fetchAlbum := command.Command{
		Name:        "album",
		Description: "Toggle the option to show album art (default is false)",
		Aliases:     []string{"a"},
		Handle: func(input string) string {
			player.State.Data.FetchAlbum = !player.State.Data.FetchAlbum
			storage.Save(player.State.Data)
			return fmt.Sprintf("Show album set to %v", player.State.Data.FetchAlbum)
		},
	}

	loadMPRIS := command.Command{
		Name:        "mpris",
		Description: "Toggle the option to load mpv-mpris (default is false)",
		Aliases:     []string{"m"},
		Handle: func(input string) string {
			player.State.Data.LoadMPRIS = !player.State.Data.LoadMPRIS
			storage.Save(player.State.Data)
			return fmt.Sprintf("Load MPRIS set to %v", player.State.Data.LoadMPRIS)
		},
	}

	searchSoundCloud := command.Command{
		Name:        "soundcloud",
		Description: "Toggle the option to search on SoundCloud (default is false)",
		Aliases:     []string{"sc"},
		Handle: func(input string) string {
			player.State.Data.SearchSoundCloud = !player.State.Data.SearchSoundCloud
			storage.Save(player.State.Data)
			return fmt.Sprintf("Search SoundCloud set to %v", player.State.Data.SearchSoundCloud)
		},
	}

	command.RegisterCommands(help, keepLiveCache, fetchAlbum, loadMPRIS, searchSoundCloud)
}
