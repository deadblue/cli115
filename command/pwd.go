package command

import (
	"fmt"
	"go.dead.blue/cli115/core"
	"strings"
)

/*
"pwd" command print full path of current directory.
*/
type PwdCommand struct{}

func (c *PwdCommand) Name() string {
	return "pwd"
}

func (c *PwdCommand) Exec(ctx *core.Context, args string) (err error) {
	sb := strings.Builder{}
	sb.WriteString("/")
	for i, v := range ctx.Path.Values() {
		dir, ok := v.(*core.Dir)
		if !ok {
			continue
		}
		if i > 0 {
			sb.WriteString("/")
		}
		sb.WriteString(dir.Name)
	}
	fmt.Println(sb.String())
	return nil
}
