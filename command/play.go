package command

import (
	"go.dead.blue/cli115/core"
)

type PlayCommand struct{}

func (c *PlayCommand) Name() string {
	return "play"
}

func (c *PlayCommand) Exec(ctx *core.Context, args string) (err error) {
	return nil
}
