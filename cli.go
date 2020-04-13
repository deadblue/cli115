package cli115

import (
	"dead.blue/cli115/command"
	"dead.blue/cli115/container"
	"dead.blue/cli115/core"
	"github.com/peterh/liner"
)

func Run() error {
	opts := FromCommandLine()
	if t, err := initTerminal(opts); err == nil {
		return t.Run()
	} else {
		return err
	}
}

func initTerminal(opts *Options) (t *Terminal, err error) {
	ctx, err := initContext(opts)
	if err != nil {
		return
	}
	t = &Terminal{
		ctx:   ctx,
		state: createLinerState(),
	}
	// register commands
	t.Register(&command.LsCommand{},
		&command.PwdCommand{},
		&command.PullCommand{},
		&command.PushCommand{},
		&command.PlayCommand{})
	return
}

func initContext(opts *Options) (ctx *core.Context, err error) {
	agent, err := initAgent(opts)
	if err == nil {
		ctx = &core.Context{
			Agent:  agent,
			Prefix: "115",
			Path:   container.NewStack(),
		}
	}
	return
}

func createLinerState() *liner.State {
	state := liner.NewLiner()
	state.SetCtrlCAborts(true)
	state.SetTabCompletionStyle(liner.TabPrints)
	return state
}
