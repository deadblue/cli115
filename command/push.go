package command

import (
	"go.dead.blue/cli115/context"
)

type PushCommand struct {
	ArgsCommand
}

func (c *PushCommand) Name() string {
	return "push"
}

func (c *PushCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	return nil
}
