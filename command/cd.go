package command

import (
	"go.dead.blue/cli115/core"
	"log"
)

type CdCommand struct {
	ArgsCommand
}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) Exec(ctx *core.Context, args []string) (err error) {
	rootId := "0"
	if len(args) > 0 {
		target := args[0]
		if target[0] != '/' && !ctx.Path.IsEmpty() {
			// target is a relative path, search from current dir
			root := ctx.Path.Top().(*core.Dir)
			rootId = root.Id
		}
	}
	// TODO: parse path
	log.Printf("Search from => %s", rootId)
	return nil
}
