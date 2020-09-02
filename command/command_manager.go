package command

import (
	"strings"

	"github.com/Pauloo27/tuner/utils"
)

var commands = map[string]Command{}

func RegisterCommand(command Command) {
	commands[command.Name] = command

	for _, alias := range command.Aliases {
		commands[alias] = command
	}
}

func InvokeCommand(input string) (found bool) {
	lowerCased := strings.ToLower(input)
	label := strings.Split(lowerCased, " ")[0]
	command, ok := commands[label]
	if !ok {
		found = false
		return
	}

	utils.MoveCursorTo(1, 1)
	utils.ClearScreen()

	commandInput := strings.TrimPrefix(lowerCased, label)
	command.Handle(commandInput)

	_, _ = utils.AskFor("Press enter to close")

	found = true
	return
}
