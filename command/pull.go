package command

import (
	"go.dead.blue/cli115/core"
)

type PullCommand struct {
	ArgsCommand
}

func (c *PullCommand) Name() string {
	return "pull"
}

func (c *PullCommand) Exec(ctx *core.Context, args []string) (err error) {
	return nil
}
