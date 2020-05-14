package command

import (
	"bytes"
	"fmt"
	"go.dead.blue/cli115/internal/app/context"
	"os/exec"
)

type PlayCommand struct {
	ArgsCommand
}

func (c *PlayCommand) Name() string {
	return "play"
}

func (c *PlayCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	if len(args) == 0 {
		return errArgsNotEnough
	}
	// Check mpv
	if ctx.Conf.Mpv == nil {
		return errMpvNotExist
	}
	// Search file
	file := ctx.Fs.File(args[0])
	if file == nil {
		return errFileNotExist
	}
	if !file.IsFile {
		return errNotFile
	}
	// Play video via mpv
	hls, err := ctx.Agent.VideoHlsContent(file.PickCode)
	if err != nil {
		return
	}
	cmd := exec.Command(ctx.Conf.Mpv.Path,
		fmt.Sprintf("--title=%s", file.Name))
	if ctx.Conf.Mpv.Fs {
		cmd.Args = append(cmd.Args, "--fullscreen")
	}
	// Read HLS content from stdin
	cmd.Args = append(cmd.Args, "-")
	cmd.Stdin = bytes.NewReader(hls)
	// TODO: handle interrupt signal.
	return cmd.Run()
}

func (c *PlayCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (head string, choices []string) {
	head = ""
	if index == 0 {
		choices = ctx.Fs.FileNames(prefix)
	} else {
		choices = make([]string, 0)
	}
	return
}
