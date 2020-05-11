package command

import (
	"go.dead.blue/cli115/internal/impl/context"
)

type ExitCommand struct {
	NoArgsCommand
}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) ImplExec(ctx *context.Impl, _ []string) error {
	ctx.Die()
	return nil
}
