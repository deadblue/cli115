package command

import (
	"fmt"
	"go.dead.blue/cli115/core"
)

type ClearCommand struct{}

func (c *ClearCommand) Name() string {
	return "clear"
}

func (c *ClearCommand) Exec(ctx *core.Context, _ string) (err error) {
	fmt.Println("Sorry, this command does not support your OS.")
	return nil
}
