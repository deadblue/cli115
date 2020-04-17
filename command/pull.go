package command

import (
	"go.dead.blue/cli115/context"
)

type PullCommand struct {
	ArgsCommand
}

func (c *PullCommand) Name() string {
	return "pull"
}

func (c *PullCommand) Exec(ctx *context.Impl, args []string) error {
	return nil
}
