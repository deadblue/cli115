package cli115

import (
	"github.com/peterh/liner"
	"go.dead.blue/cli115/command"
)

func Run() error {
	opts := FromCommandLine()
	if t, err := createTerminal(opts); err == nil {
		return t.Run()
	} else {
		return err
	}
}

func createTerminal(opts *Options) (t *Terminal, err error) {
	agent, err := initAgent(opts)
	if err != nil {
		return
	}
	ctx, err := createContext(agent)
	if err != nil {
		return
	}
	t = &Terminal{
		ctx:   ctx,
		state: createLinerState(),
	}
	t.state.SetCompleter(t.Completer)
	// Register commands
	t.Register(&command.ExitCommand{},
		&command.LsCommand{},
		&command.PwdCommand{},
		&command.PullCommand{},
		&command.PushCommand{},
		&command.PlayCommand{})
	return
}

func createLinerState() *liner.State {
	state := liner.NewLiner()
	state.SetCtrlCAborts(true)
	state.SetTabCompletionStyle(liner.TabPrints)
	return state
}
