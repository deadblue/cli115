package command

import (
	"fmt"
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/internal/app/config"
	"go.dead.blue/cli115/internal/app/context"
	"os"
	"os/exec"
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
	// Find file
	file := ctx.Fs.File(args[0])
	if file == nil {
		return errFileNotExist
	}
	if !file.IsFile {
		return errNotFile
	}
	fmt.Printf("Downloading file: %s\n", file.Name)
	// Create download ticket
	ticket, err := ctx.Agent.CreateDownloadTicket(file.PickCode)
	if err != nil {
		return
	}
	// Prefer to use aria2 if available.
	if ctx.Conf.Aria2 != nil {
		return c.aria2Download(ctx, ticket, file.Sha1)
	} else if ctx.Conf.Curl != nil {
		return c.curlDownload(ctx.Conf.Curl, ticket)
	} else {
		return errNoDownloader
	}
}

func (c *DlCommand) aria2Download(ctx *context.Impl, ticket *elevengo.DownloadTicket, sha1 string) error {
	if conf := ctx.Conf.Aria2; conf.Rpc {
		fmt.Println("Downloading file via aira2 RPC ...")
		// headers
		headers, i := make([]string, len(ticket.Headers)), 0
		for name, value := range ticket.Headers {
			headers[i] = fmt.Sprintf("%s: %s", name, value)
			i += 1
		}
		// options
		options := map[string]interface{}{
			"max-connection-per-server": 2,
			"split":                     16,
			"min-split-size":            "1M",
			"header":                    headers,
			"out":                       ticket.FileName,
			"checksum":                  fmt.Sprintf("sha-1=%s", sha1),
		}
		return ctx.Aria2.AddTask(ticket.Url, options)
	} else {
		fmt.Println("Downloading file via aira2 ...")
		cmd := exec.Command(conf.Path,
			"--max-connection-per-server=2",
			"--split=16", "--min-split-size=1M",
			fmt.Sprintf("--out=%s", ticket.FileName),
			fmt.Sprintf("--checksum=sha-1=%s", sha1),
		)
		for name, value := range ticket.Headers {
			cmd.Args = append(cmd.Args, fmt.Sprintf("--header=%s: %s", name, value))
		}
		cmd.Args = append(cmd.Args, ticket.Url)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		return cmd.Run()
	}
}

func (c *DlCommand) curlDownload(conf *config.CurlConf, ticket *elevengo.DownloadTicket) error {
	fmt.Println("Downloading file via curl ...")
	cmd := exec.Command(conf.Path, "-#", ticket.Url)
	for name, value := range ticket.Headers {
		cmd.Args = append(cmd.Args, "-H", fmt.Sprintf("%s: %s", name, value))
	}
	cmd.Args = append(cmd.Args, "-o", ticket.FileName)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func (c *DlCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	head = ""
	if index == 0 {
		choices = ctx.Fs.FileNames(prefix)
	} else {
		choices = make([]string, 0)
	}
	return
}
