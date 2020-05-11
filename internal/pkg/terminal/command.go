package terminal

/*
Command declare an interface that a command should be implemented.
*/
type Command interface {
	// Command name.
	Name() string
	// Indicate whether the command has arguments.
	HasArgs() bool
	// Execute the command.
	Exec(ctx Context, args []string) (err error)
}

/*
ArgCompleter interface is to indicate whether command supoorts argument completer.
*/
type ArgCompleter interface {

	// To get complete choices of an argument.
	// Parameter:
	//   ctx:    Terminal context.
	//   index:  The index of the argument that need to be compelte.
	//   prefix: The prefix of the argument.
	// Return:
	//   head:    Head of choice.
	//   choices: Available choices.
	Completer(ctx Context, index int, prefix string) (head string, choices []string)
}
