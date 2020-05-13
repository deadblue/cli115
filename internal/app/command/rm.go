package command

import (
	"fmt"
	"go.dead.blue/cli115/internal/app/context"
)

type RmCommand struct {
	ArgsCommand
}

func (c *RmCommand) Name() string {
	return "rm"
}

func (c *RmCommand) ImplExec(ctx *context.Impl, args []string) error {
	if len(args) == 0 {
		return errArgsNotEnough
	}
	files := ctx.Fs.Files(args[0])
	if fileCount := len(files); fileCount == 0 {
		return nil
	} else {
		fileIds := make([]string, fileCount)
		for i, file := range files {
			fileIds[i] = file.FileId
		}
		err := ctx.Agent.FileDelete(ctx.Fs.Curr().Id, fileIds...)
		if err == nil {
			fmt.Printf("Removed files: %d\n", fileCount)
		}
		return err
	}
}

func (c *RmCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	head = ""
	if index == 0 {
		choices = ctx.Fs.FileNames(prefix)
	} else {
		choices = make([]string, 0)
	}
	return
}
