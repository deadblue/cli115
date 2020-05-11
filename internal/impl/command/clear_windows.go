package command

import (
	"fmt"
	"go.dead.blue/cli115/internal/impl/context"
)

type ClearCommand struct {
	NoArgsCommand
}

func (c *ClearCommand) Name() string {
	return "clear"
}

func (c *ClearCommand) ImplExec(_ *context.Impl, _ []string) error {
	fmt.Println("Sorry, this command does not support your OS.")
	return nil
}
