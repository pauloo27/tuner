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

	showVideo := command.Command{
		Name:        "video",
		Description: "Set the option to show the video",
		Aliases:     []string{"v"},
		Handle: func(input string) string {
			options.Options.ShowVideo = !options.Options.ShowVideo
			return fmt.Sprintf("Show video set to %v", options.Options.ShowVideo)
		},
	}

	cmds := []command.Command{help, showVideo}

	for _, cmd := range cmds {
		command.RegisterCommand(cmd)
	}
}
