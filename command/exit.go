package command

import (
	"go.dead.blue/cli115/context"
)

type ExitCommand struct {
	NoArgsCommand
}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) Exec(ctx *context.Impl, _ []string) error {
	ctx.Die()
	return nil
}
