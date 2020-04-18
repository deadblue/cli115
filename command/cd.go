package command

import (
	"go.dead.blue/cli115/context"
	"strings"
)

type CdCommand struct {
	ArgsCommand
}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	// TODO
	return nil
}

func (c *CdCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (choices []string) {
	choices = make([]string, 0)
	// Only handle first arguments
	if index > 0 {
		return
	}
	// parse prefix first
	parts := strings.Split(prefix, "/")
	partCount := len(parts)
	if partCount > 1 {
		root, start := ctx.Curr, 0
		if parts[0] == "" {
			root, start = ctx.Root, 1
		}
		for i := start; i < partCount-1; i++ {
			// TODO: start directory from "root"
		}
	}
	return
}
