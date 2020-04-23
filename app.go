package cli115

import (
	"go.dead.blue/cli115/command"
	"go.dead.blue/cli115/context"
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
	ctx, err := context.New(agent)
	if err != nil {
		return
	}
	t = core.NewTerminal(ctx)
	// Register commands
	t.Register(
		command.Wrap(&command.CdCommand{}),
		command.Wrap(&command.ClearCommand{}),
		command.Wrap(&command.ExitCommand{}),
		command.Wrap(&command.LsCommand{}),
		command.Wrap(&command.PlayCommand{}),
		command.Wrap(&command.DlCommand{}),
		command.Wrap(&command.PushCommand{}),
		command.Wrap(&command.PwdCommand{}))
	return
}
