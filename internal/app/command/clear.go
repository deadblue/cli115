// +build !windows

package command

import (
	"fmt"
	"go.dead.blue/cli115/internal/app/context"
)

type ClearCommand struct {
	NoArgsCommand
}

func (c *ClearCommand) Name() string {
	return "clear"
}

func (c *ClearCommand) ImplExec(_ *context.Impl, _ []string) error {
	fmt.Print("\x1b[H\x1b[2J")
	return nil
}
