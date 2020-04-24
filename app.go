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
		// Exit the terminal
		command.Wrap(&command.ExitCommand{}),
		// Clear screen
		command.Wrap(&command.ClearCommand{}),
		// Change work directory
		command.Wrap(&command.CdCommand{}),
		// Print work directory
		command.Wrap(&command.PwdCommand{}),
		// List files under work directory
		command.Wrap(&command.LsCommand{}),
		// Download file to local
		command.Wrap(&command.DlCommand{}),
		// Upload file from local
		command.Wrap(&command.UlCommand{}),
		// Play a remote video
		command.Wrap(&command.PlayCommand{}),
	)
	return
}
