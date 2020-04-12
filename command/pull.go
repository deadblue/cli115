package command

import (
	"dead.blue/cli115/core"
)

type PullCommand struct{}

func (c *PullCommand) Name() string {
	return "pull"
}

func (c *PullCommand) Exec(ctx *core.Context, args string) (err error) {
	return nil
}
