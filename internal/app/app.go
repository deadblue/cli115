package app

import (
	"go.dead.blue/cli115/internal/impl/command"
	"go.dead.blue/cli115/internal/impl/context"
	"go.dead.blue/cli115/internal/pkg/terminal"
)

const (
	appName = "cli115"
	appVer  = "0.0.1"
)

func Run() error {
	opts, conf := ParseOptions(), LoadConf()
	if t, err := createTerminal(opts, conf); err == nil {
		return t.Run()
	} else {
		return err
	}
}

func createTerminal(opts *Options, conf *Conf) (t *terminal.Terminal, err error) {
	agent, err := initAgent(opts)
	if err != nil {
		return
	}
	ctx, err := context.New(agent, conf)
	if err != nil {
		return
	}
	t = terminal.New(ctx)
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
