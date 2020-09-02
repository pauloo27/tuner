package commands

import (
	"fmt"

	"github.com/Pauloo27/tuner/command"
	"github.com/Pauloo27/tuner/options"
)

func SetupDefaultCommands() {
	help := command.Command{
		Name:        "help",
		Description: "List all the commands",
		Aliases:     []string{"h"},
		Handle: func(input string) string {
			fmt.Println("No help!")
			return ""
		},
	}

	keepLiveCache := command.Command{
		Name: "cache",
		Description: "Toggle the option to keep cache of a live when one is playing (default is false)." +
			" When true, you can seek back a live but it use more ram overtime",
		Aliases: []string{"c"},
		Handle: func(input string) string {
			options.Options.KeepLiveCache = !options.Options.KeepLiveCache
			return fmt.Sprintf("Keep live cache %v", options.Options.KeepLiveCache)
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

	cmds := []command.Command{help, keepLiveCache, showVideo}

	for _, cmd := range cmds {
		command.RegisterCommand(cmd)
	}
}
