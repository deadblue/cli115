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
	if len(args) == 0 {
		return
	}
	dir := ctx.Fs.LocateDir(args[0])
	if dir != nil {
		ctx.Fs.SetCurr(dir)
	} else {
		return errDirNotExist
	}
	return
}

func (c *CdCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	choices = make([]string, 0)
	// Only handle first arguments
	if index > 0 {
		return
	}
	head, last, curr := "", prefix, ctx.Curr
	pos := strings.LastIndex(prefix, "/")
	if pos >= 0 {
		head, last = prefix[:pos+1], prefix[pos+1:]
		if pos == 0 {
			curr = ctx.Root
		} else {
			curr = ctx.Fs.LocateDir(head)
		}
	}
	if curr == nil {
		return
	}
	for name := range curr.Children {
		if last == "" || strings.HasPrefix(name, last) {
			choices = append(choices, name+"/")
		}
	}
	return
}
