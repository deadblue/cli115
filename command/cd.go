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
	if len(args) == 0 {
		return
	}
	dir := c.locate(ctx, args[0])
	if dir != nil {
		ctx.Curr = dir
	} else {
		return errors.New("no such dir")
	}
	return
}

func (c *CdCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	choices = make([]string, 0)
	// Only handle first arguments
	if index > 0 {
		return
	}
	// TODO: use c.locate()
	return
}

func (c *CdCommand) locate(ctx *context.Impl, path string) (dir *context.DirNode) {
	dirs := strings.Split(path, "/")
	depth, curr, start := len(dirs), ctx.Curr, 0
	if depth > 1 && dirs[0] == "" {
		// Starts from root
		curr = ctx.Root
		start = 1
	}
	//
	for i := start; i < depth; i += 1 {
		if !curr.IsCached {
			c.fillCache(curr, ctx.Agent)
		}
		dirName := dirs[i]
		if dirName == "." || dirName == "" {
			continue
		} else if dirName == ".." {
			if curr != ctx.Root {
				curr = curr.Parent
			}
		} else {
			curr = curr.Children[dirName]
		}
		if curr == nil {
			break
		}
	}
	return curr
}

func (c *CdCommand) fillCache(dir *context.DirNode, agent *elevengo.Agent) {
	for cur := elevengo.FileCursor(); cur.HasMore(); cur.Next() {
		if files, err := agent.FileList(dir.Id, cur); err != nil {
			break
		} else {
			for _, file := range files {
				if !file.IsDirectory {
					continue
				}
				node := context.MakeNode(file.FileId, file.Name)
				node.AppendTo(dir)
			}
		}
	}
	dir.IsCached = true
}
