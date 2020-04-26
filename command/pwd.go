package command

import (
	"fmt"
	"go.dead.blue/cli115/context"
	"strings"
)

type PwdCommand struct {
	NoArgsCommand
}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) ImplExec(ctx *context.Impl, _ []string) error {
	currDir := ctx.Fs.Curr()
	dirs, depth := make([]string, currDir.Depth+2), currDir.Depth
	for dir := currDir; dir != nil; dir = dir.Parent {
		dirs[depth] = dir.Name
		depth -= 1
	}
	fmt.Println(strings.Join(dirs, "/"))
	return nil
}
