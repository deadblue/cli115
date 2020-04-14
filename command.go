package cli115

import "go.dead.blue/cli115/core"

/*
Command declare an interface that a command should be implemented.
*/
type Command interface {
	// Command name.
	Name() string
	// Indicate whether the command has arguments.
	HasArgs() bool
	// Execute the command.
	Exec(ctx *core.Context, args string) (err error)
}

/*
CommandCompleter is an additional interface to indicate if command supoorts tab complete.
*/
type CommandCompleter interface {
	Completer(ctx *core.Context, args string) []string
}
