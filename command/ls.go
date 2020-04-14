package command

import (
	"github.com/deadblue/elevengo"
	"github.com/olekukonko/tablewriter"
	"go.dead.blue/cli115/core"
	"os"
	"strconv"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Exec(ctx *core.Context, args string) (err error) {
	dirId := "0"
	if !ctx.Path.IsEmpty() {
		dirId = (ctx.Path.Top()).(string)
	}
	w := tablewriter.NewWriter(os.Stdout)
	w.SetAutoWrapText(false)
	w.SetRowLine(false)
	w.SetHeader([]string{
		"ID", "Size", "Name",
	})
	for cursor := elevengo.FileCursor(); cursor.HasMore(); cursor.Next() {
		files, err := ctx.Agent.FileList(dirId, cursor)
		if err != nil {
			return err
		} else {
			for _, file := range files {
				w.Append([]string{
					file.FileId,
					strconv.FormatInt(file.Size, 10),
					file.Name,
				})
			}
		}
	}
	w.Render()
	return
}
