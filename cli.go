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
	agent, err := initAgent(opts)
	if err != nil {
		return
	}
	t = &Terminal{
		ctx: &core.Context{
			Alive: true,
			Agent: agent,
			Path:  container.NewStack(),
			User:  agent.User(),
		},
		state: createLinerState(),
	}
	t.state.SetCompleter(t.Completer)
	// register commands
	t.Register(&command.LsCommand{},
		&command.PwdCommand{},
		&command.PullCommand{},
		&command.PushCommand{},
		&command.PlayCommand{},
		&command.ExitCommand{})
	return
}

func createLinerState() *liner.State {
	state := liner.NewLiner()
	state.SetCtrlCAborts(true)
	state.SetTabCompletionStyle(liner.TabPrints)
	return state
}
