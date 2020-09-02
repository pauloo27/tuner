package command

type CommandHandler func(input string) string

type Command struct {
	Name, Description string
	Aliases           []string
	Handle            CommandHandler
}
