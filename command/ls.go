package command

import (
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/core"
	"go.dead.blue/cli115/table"
	"go.dead.blue/cli115/util"
	"os"
	"regexp"
	"strings"
)

type LsCommand struct {
	ArgsCommand
}

func (c *LsCommand) Name() string {
	return "ls"
}

func (c *LsCommand) Exec(ctx *core.Context, args string) (err error) {
	dirId := "0"
	if !ctx.Path.IsEmpty() {
		dirId = (ctx.Path.Top()).(string)
	}
	reg := c.parsePattern(args)
	// Clear cache, we will update it
	ctx.Cache = make(map[string]*elevengo.File)
	// Print file list
	tbl := table.New().
		AddColumn("Size", table.AlignRight).
		AddColumn("Name", table.AlignLeft)
	for cursor := elevengo.FileCursor(); cursor.HasMore(); cursor.Next() {
		files, err := ctx.Agent.FileList(dirId, cursor)
		if err != nil {
			return err
		} else {
			for _, file := range files {
				// Update cache
				ctx.Cache[file.Name] = file
				if reg != nil && !reg.MatchString(file.Name) {
					continue
				}
				if file.IsDirectory {
					tbl.AppendRow([]string{"<DIR>", file.Name})
				} else {
					tbl.AppendRow([]string{util.FormatSize(file.Size), file.Name})
				}

			}
		}
	}
	tbl.Render(os.Stdout)
	return
}

func (c *LsCommand) parsePattern(pattern string) *regexp.Regexp {
	if pattern == "" {
		return nil
	}
	rp := strings.NewReplacer(".", "\\.", "*", ".*")
	return regexp.MustCompile("^" + rp.Replace(pattern) + "$")
}
