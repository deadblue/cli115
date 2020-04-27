package command

import (
	"go.dead.blue/cli115/context"
	"go.dead.blue/cli115/internal/pkg/spinner"
	"strings"
)

type CdCommand struct {
	ArgsCommand
}

func (c *CdCommand) Name() string {
	return "cd"
}

func (c *CdCommand) ImplExec(ctx *context.Impl, args []string) error {
	if len(args) == 0 {
		return errArgsNotEnough
	}
	if dir := ctx.Fs.LocateDir(args[0]); dir != nil {
		s := spinner.NewBuilder().
			Suffix(" Enter directory...").
			Complete("Done!").Build()
		s.Start()
		ctx.Fs.SetCurr(dir)
		s.Stop()
	} else {
		return errDirNotExist
	}
	return nil
}

func (c *CdCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	choices = make([]string, 0)
	// Only handle first arguments
	if index > 0 {
		return
	}
	head, last, curr := "", prefix, ctx.Fs.Curr()
	if pos := strings.LastIndex(prefix, "/"); pos >= 0 {
		head, last = prefix[:pos+1], prefix[pos+1:]
		curr = ctx.Fs.LocateDir(head)
	}
	if curr != nil {
		choices = ctx.Fs.DirNames(curr, last)
	}
	return
}
