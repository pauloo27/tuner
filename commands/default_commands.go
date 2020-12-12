package commands

import (
	"fmt"
	"strings"

	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/options"
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
			options.Options.Cache = !options.Options.Cache
			return fmt.Sprintf("Cache set to %v", options.Options.Cache)
		},
	}

	showVideo := command.Command{
		Name:        "video",
		Description: "Toggle the option to show the video (default is false)",
		Aliases:     []string{"v"},
		Handle: func(input string) string {
			options.Options.ShowVideo = !options.Options.ShowVideo
			return fmt.Sprintf("Show video set to %v", options.Options.ShowVideo)
		},
	}

	command.RegisterCommands(help, keepLiveCache, showVideo)
}
