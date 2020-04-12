package command

import (
	"dead.blue/cli115/core"
)

type PushCommand struct{}

func (c *PushCommand) Name() string {
	return "push"
}

func (c *PushCommand) Exec(ctx *core.Context, args string) (err error) {
	return nil
}
