package cli115

import (
	"dead.blue/cli115/command"
	"dead.blue/cli115/container"
	"dead.blue/cli115/context"
	"github.com/deadblue/elevengo"
	"github.com/peterh/liner"
)

func Run() (err error) {
	opts := FromCommandLine()
	return createTerminal(opts).Run()
}

func createTerminal(opts *Options) *Terminal {
	agent := elevengo.Default()
	ctx := &context.Context{
		Agent:  agent,
		Prefix: "115",
		Path:   container.NewStack(),
	}
	t := &Terminal{
		context: ctx,
		state:   createLinerState(),
	}
	t.Register(&command.LsCommand{},
		&command.PwdCommand{})
	return t
}

func createLinerState() *liner.State {
	state := liner.NewLiner()
	state.SetCtrlCAborts(true)
	state.SetTabCompletionStyle(liner.TabPrints)
	return state
}
