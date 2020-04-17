package command

import (
	"bytes"
	"errors"
	"go.dead.blue/cli115/core"
	"go.dead.blue/cli115/util"
	"os"
	"os/exec"
	"strings"
)

type PlayCommand struct {
	ArgsCommand
}

func (c *PlayCommand) Name() string {
	return "play"
}

func (c *PlayCommand) Exec(ctx *core.Context, args []string) (err error) {
	if len(args) == 0 {
		return errors.New("no file to play")
	}
	// search mpv
	exe, err := exec.LookPath("mpv")
	if err != nil {
		return errors.New("can not find mpv executable file")
	}
	// search file
	file, ok := ctx.Cache[args[0]]
	if !ok {
		return os.ErrNotExist
	}
	if !file.IsFile {
		return errors.New("not a regular file")
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

func (c *PlayCommand) Completer(ctx *core.Context, index int, prefix string) (choices []string) {
	choices = make([]string, 0)
	// "play" command only handle first argument
	if index > 0 {
		return
	}
	// Search file from file cache
	for name, file := range ctx.Cache {
		// Skip directory
		if file.IsDirectory {
			continue
		}
		if len(prefix) == 0 || strings.HasPrefix(name, prefix) {
			choices = append(choices, util.InputFieldEscape(name))
		}
	}
	return
}
