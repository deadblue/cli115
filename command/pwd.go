package command

import "dead.blue/cli115/context"

type PwdCommand struct{}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) Exec(context *context.Context, args string) (err error) {
	return nil
}
