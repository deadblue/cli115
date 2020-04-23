package command

import (
	"bytes"
	"go.dead.blue/cli115/context"
	"go.dead.blue/cli115/util"
	"os/exec"
	"strings"
)

type PlayCommand struct {
	ArgsCommand
}

func (c *PlayCommand) Name() string {
	return "play"
}

func (c *PlayCommand) ImplExec(ctx *context.Impl, args []string) (err error) {
	if len(args) == 0 {
		return errFileNotExist
	}
	// search mpv
	exe, err := exec.LookPath("mpv")
	if err != nil {
		return errMpvNotExist
	}
	// search file
	file, ok := ctx.Files[args[0]]
	if !ok {
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
	cmd := exec.Command(exe, "-")
	cmd.Stdin = bytes.NewReader(hls)
	return cmd.Run()
}

func (c *PlayCommand) ImplCplt(ctx *context.Impl, index int, prefix string) (choices []string) {
	choices = make([]string, 0)
	// "play" command only handle first argument
	if index > 0 {
		return
	}
	// Search file from file cache
	for name := range ctx.Files {
		if len(prefix) == 0 || strings.HasPrefix(name, prefix) {
			choices = append(choices, util.InputFieldEscape(name))
		}
	}
	return
}
