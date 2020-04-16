package core

/*
Command declare an interface that a command should be implemented.
*/
type Command interface {
	// Command name.
	Name() string
	// Indicate whether the command has arguments.
	HasArgs() bool
	// Execute the command.
	Exec(ctx *Context, args []string) (err error)
}

/*
CommandCompleter is an additional interface to indicate if command supoorts completer.
*/
type CommandCompleter interface {

	// The cursor is at the end of the last args, command SHOULD only give choices for it.
	Completer(ctx *Context, args []string) (choices []string)
}
