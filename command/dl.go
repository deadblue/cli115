package command

import (
	"fmt"
	"go.dead.blue/cli115/context"
	"os"
	"os/exec"
	"strings"
)

type DlCommand struct {
	ArgsCommand
}

func (c *DlCommand) Name() string {
	return "dl"
}

func (c *DlCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	if len(args) == 0 {
		return errArgsNotEnough
	}
	// Check is curl exist.
	exe, err := exec.LookPath("curl")
	if err != nil {
		return errCurlNotExist
	}
	// Find file from cache.
	file, ok := ctx.Files[args[0]]
	if !ok {
		return errFileNotExist
	}
	if !file.IsFile {
		return errNotFile
	}
	// Get download ticket
	ticket, err := ctx.Agent.CreateDownloadTicket(file.PickCode)
	if err != nil {
		return
	}
	// Download via curl
	cmd := exec.Command(exe, "-#", ticket.Url)
	for name, value := range ticket.Headers {
		cmd.Args = append(cmd.Args, "-H", fmt.Sprintf("%s: %s", name, value))
	}
	cmd.Args = append(cmd.Args, "-o", ticket.FileName)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	fmt.Printf("Downloading file: %s\n", ticket.FileName)
	// TODO: handle interrupt signal.
	return cmd.Run()
}

func (c *DlCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	head, choices = "", make([]string, 0)
	if index > 0 {
		return
	}
	// Search file from file cache
	for name := range ctx.Files {
		if len(prefix) == 0 || strings.HasPrefix(name, prefix) {
			choices = append(choices, name)
		}
	}
	return
}
