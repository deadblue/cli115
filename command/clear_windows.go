package command

import (
	"fmt"
)

type ClearCommand struct {
	NoArgsCommand
}

func (c *ClearCommand) Name() string {
	return "clear"
}

func (c *ClearCommand) Exec(_ *context.Impl, _ []string) error {
	fmt.Println("Sorry, this command does not support your OS.")
	return nil
}
