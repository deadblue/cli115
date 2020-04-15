package cli115

import (
	"go.dead.blue/cli115/command"
	"go.dead.blue/cli115/core"
)

func Run() error {
	opts := FromCommandLine()
	if t, err := createTerminal(opts); err == nil {
		return t.Run()
	} else {
		return err
	}
}

func createTerminal(opts *Options) (t *core.Terminal, err error) {
	agent, err := initAgent(opts)
	if err != nil {
		return
	}
	ctx, err := createContext(agent)
	if err != nil {
		return
	}
	t = core.NewTerminal(ctx)
	// Register commands
	t.Register(
		&command.ExitCommand{},
		&command.ClearCommand{},
		&command.CdCommand{},
		&command.PwdCommand{},
		&command.LsCommand{},
		&command.PullCommand{},
		&command.PushCommand{},
		&command.PlayCommand{})
	return
}
