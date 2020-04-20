package command

import (
	"errors"
	"github.com/deadblue/elevengo"
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
	// Do nothing by default
	target := "."
	if len(args) > 0 {
		target = args[0]
	}
	if target == "/" {
		// Go to root
		ctx.Curr = ctx.Root
	} else if target == ".." {
		// Go to parent dir
		if ctx.Curr != ctx.Root {
			ctx.Curr = ctx.Curr.Parent
		}
	} else {
		// Search target dir
		var targetDir = (*elevengo.File)(nil)
		for cur := elevengo.FileCursor(); cur.HasMore() && targetDir == nil; cur.Next() {
			if files, err := ctx.Agent.FileList(ctx.Curr.Id, cur); err != nil {
				break
			} else {
				for _, file := range files {
					if file.IsDirectory && target == file.Name {
						targetDir = file
						break
					}
				}
			}
		}
		// Go to dir
		if targetDir != nil {
			node := context.MakeNode(targetDir.FileId, targetDir.Name)
			node.AppendTo(ctx.Curr)
			ctx.Curr = node
		} else {
			err = errors.New("no such dir")
		}
	}
	return
}

func (c *CdCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (choices []string) {
	choices = make([]string, 0)
	// Only handle first arguments
	if index > 0 {
		return
	}
	//
	if strings.HasPrefix(prefix, "/") || strings.HasPrefix(prefix, "..") {
		return
	}
	// Search directories under current dir
	for cur := elevengo.FileCursor(); cur.HasMore(); cur.Next() {
		if files, err := ctx.Agent.FileList(ctx.Curr.Id, cur); err != nil {
			break
		} else {
			for _, file := range files {
				if !file.IsDirectory {
					continue
				}
				if prefix == "" || strings.HasPrefix(file.Name, prefix) {
					choices = append(choices, file.Name)
				}
			}
		}
	}
	return
}
