package command

import (
	"go.dead.blue/cli115/core"
	"strings"
)

type CdCommand struct{}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) Exec(ctx *core.Context, args string) (err error) {
	// TODO
	return nil
}

func (c *CdCommand) parsePath(ctx *core.Context, target string) (abs bool, dirs []string) {
	abs = target[0] == '/'
	for _, dir := range strings.Split(target, "/") {
		if dir == "" || dir == "." {
			continue
		}

	}
	return
}
