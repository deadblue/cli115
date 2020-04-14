package command

import (
	"bytes"
	"errors"
	"go.dead.blue/cli115/core"
	"os"
	"os/exec"
)

type PlayCommand struct{}

func (c *PlayCommand) Name() string {
	return "play"
}

func (c *PlayCommand) Exec(ctx *core.Context, args string) (err error) {
	if args == "" {
		return errors.New("no file to play")
	}
	// search mpv
	exe, err := exec.LookPath("mpv")
	if err != nil {
		return errors.New("can not find mpv executable file")
	}
	// search file
	file, ok := ctx.Cache[args]
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
