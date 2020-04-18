package command

import (
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/context"
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

func (c *LsCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	dirId := ctx.Curr.Id

	var filter *regexp.Regexp
	if len(args) > 0 {
		filter = c.parsePattern(args[0])
	}

	// Clear cache, we will update it
	ctx.Files = make(map[string]*elevengo.File)
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
				ctx.Files[file.Name] = file
				if filter != nil && !filter.MatchString(file.Name) {
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
