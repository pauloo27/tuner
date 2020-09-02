package commands

import (
	"fmt"

	"github.com/Pauloo27/tuner/command"
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

	cmds := []command.Command{help}

	for _, cmd := range cmds {
		command.RegisterCommand(cmd)
	}
}
