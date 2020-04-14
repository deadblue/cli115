package command

import "go.dead.blue/cli115/core"

type ExitCommand struct {
	NoArgsCommand
}

func (c *ExitCommand) Name() string {
	return "exit"
}

func (c *ExitCommand) Exec(ctx *core.Context, args string) (err error) {
	ctx.Alive = false
	return nil
}
