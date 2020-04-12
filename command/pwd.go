package command

import (
	"dead.blue/cli115/core"
)

type PwdCommand struct{}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) Exec(ctx *core.Context, args string) (err error) {
	return nil
}
