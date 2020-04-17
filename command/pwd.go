package command

import (
	"fmt"
	"go.dead.blue/cli115/context"
)

type PwdCommand struct {
	NoArgsCommand
}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) Exec(ctx *context.Impl, _ []string) error {
	fmt.Println(ctx.Curr.Path("/"))
	return nil
}
