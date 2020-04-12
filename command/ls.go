package command

import (
	"dead.blue/cli115/context"
	"github.com/deadblue/elevengo"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type LsCommand struct{}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Exec(context *context.Context, args string) (err error) {
	dirId := "0"
	if !context.Path.IsEmpty() {
		dirId = (context.Path.Top()).(string)
	}
	w := tablewriter.NewWriter(os.Stdout)
	w.SetAutoWrapText(false)
	w.SetRowLine(false)
	w.SetHeader([]string{
		"ID", "Size", "Name",
	})
	for cursor := elevengo.FileCursor(); cursor.HasMore(); cursor.Next() {
		files, err := context.Agent.FileList(dirId, cursor)
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
