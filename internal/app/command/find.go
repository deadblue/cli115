package command

import (
	"fmt"
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/internal/app/context"
	"go.dead.blue/cli115/internal/pkg/table"
	"go.dead.blue/cli115/internal/pkg/util"
	"os"
	"strconv"
	"strings"
)

type FindCommand struct {
	ArgsCommand
}

func (c *FindCommand) Name() string {
	return "find"
}

func (c *FindCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	if len(args) == 0 {
		return errArgsNotEnough
	}

	rootId, keyword := ctx.Fs.Curr().Id, args[0]
	cache, index := make(map[string]string), 1
	tbl := table.New().
		AddColumn("#", table.AlignRight).
		AddColumn("Size", table.AlignRight).
		AddColumn("Name", table.AlignLeft)
	for cursor := elevengo.FileCursor(); cursor.HasMore(); cursor.Next() {
		files, err := ctx.Agent.FileSearch(rootId, keyword, cursor)
		if err != nil {
			break
		}
		for _, file := range files {
			filePath := fmt.Sprintf("./%s", file.Name)
			if file.ParentId != rootId {
				// Search the relative path
				dirPath, ok := cache[file.ParentId]
				if !ok {
					dirPath = c.relativeDirPath(ctx, rootId, file.ParentId)
					cache[file.ParentId] = dirPath
				}
				filePath = fmt.Sprintf("./%s%s", dirPath, file.Name)
			}
			if file.IsDirectory {
				tbl.AppendRow([]string{
					strconv.Itoa(index), "<DIR>", filePath,
				})
			} else {
				tbl.AppendRow([]string{
					strconv.Itoa(index), util.FormatSize(file.Size), filePath,
				})
			}

			index += 1
		}
	}
	tbl.Render(os.Stdout)
	return nil
}

func (c *FindCommand) relativeDirPath(ctx *context.Impl, rootId string, targetId string) string {
	info, err := ctx.Agent.FileStat(targetId)
	if err != nil {
		return ""
	}
	buf, start := strings.Builder{}, false
	for _, di := range info.Parents {
		if rootId == di.Id {
			start = true
		} else {
			if start {
				_, _ = buf.WriteString(di.Name + "/")
			}
		}
	}
	buf.WriteString(info.Name + "/")
	return buf.String()
}
