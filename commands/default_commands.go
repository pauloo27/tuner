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
		Handle: func(input string) {
			fmt.Println("No help!")
		},
	}

	showVideo := command.Command{
		Name:        "video",
		Description: "Set the option to show the video",
		Aliases:     []string{"v"},
		Handle: func(input string) {
			options.Options.ShowVideo = !options.Options.ShowVideo
			fmt.Println("Show video set to", options.Options.ShowVideo)
		},
	}

	cmds := []command.Command{help, showVideo}

	for _, cmd := range cmds {
		command.RegisterCommand(cmd)
	}
}
