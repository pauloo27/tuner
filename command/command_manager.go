package command

import (
	"strings"

	"github.com/Pauloo27/tuner/utils"
)

var Commands = map[string]Command{}

func RegisterCommand(command Command) {
	Commands[command.Name] = command

	for _, alias := range command.Aliases {
		Commands[alias] = command
	}
}

func RegisterCommands(commands ...Command) {
	for _, command := range commands {
		RegisterCommand(command)
	}
}

func InvokeCommand(input string) (found bool, out string) {
	lowerCased := strings.ToLower(input)
	label := strings.Split(lowerCased, " ")[0]
	command, ok := Commands[label]
	if !ok {
		found = false
		return
	}

	utils.MoveCursorTo(1, 1)
	utils.ClearScreen()

	commandInput := strings.TrimPrefix(lowerCased, label)
	out = command.Handle(commandInput)

	if out == "" {
		err := utils.WaitForEnter("Press enter to continue...")
		utils.HandleError(err, "Cannot read input")
	}

	found = true
	return
}
