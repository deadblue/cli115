package app

import (
	"go.dead.blue/cli115/internal/app/command"
	"go.dead.blue/cli115/internal/app/config"
	"go.dead.blue/cli115/internal/app/context"
	"go.dead.blue/cli115/internal/pkg/terminal"
)

func Run() error {
	// Create context
	opts, conf := config.ParseCommandLine(), config.Load()
	ctx := context.New(opts, conf)
	// Create and setup terminal
	t := terminal.New(ctx)
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
		// Remove file
		command.Wrap(&command.RmCommand{}),
	)
	// Run the terminal
	return t.Run()
}
