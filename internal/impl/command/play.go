package command

import (
	"bytes"
	"fmt"
	"go.dead.blue/cli115/internal/impl/context"
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
	// search mpv
	exe, err := exec.LookPath("mpv")
	if err != nil {
		return errMpvNotExist
	}
	// search file
	file := ctx.Fs.File(args[0])
	if file == nil {
		return errFileNotExist
	}
	if !file.IsFile {
		return errNotFile
	}
	// play video via mpv
	hls, err := ctx.Agent.VideoHlsContent(file.PickCode)
	if err != nil {
		return
	}
	cmd := exec.Command(exe,
		fmt.Sprintf("--title=%s", file.Name), "-")
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
