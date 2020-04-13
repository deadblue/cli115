package cli115

import (
	"dead.blue/cli115/core"
	"errors"
	"github.com/peterh/liner"
	"strings"
)

type Command interface {
	Name() string
	Exec(ctx *core.Context, args string) (err error)
}

type Terminal struct {
	state *liner.State
	ctx   *core.Context
	cmds  map[string]Command
}

func (t *Terminal) Register(cmds ...Command) {
	if t.cmds == nil {
		t.cmds = make(map[string]Command)
	}
	for _, cmd := range cmds {
		t.cmds[cmd.Name()] = cmd
	}
}

func (t *Terminal) Run() (err error) {
	for t.ctx.Alive {
		if input, err := t.state.Prompt(t.ctx.Prompt()); err != nil {
			return err
		} else {
			input = strings.TrimSpace(input)
			err = t.handle(input)
			t.state.AppendHistory(input)
		}
	}
	return nil
}

func (t *Terminal) handle(input string) (err error) {
	cmd, args := input, ""
	pos := strings.IndexRune(input, ' ')
	if pos > 0 {
		cmd, args = input[:pos], input[pos+1:]
	}
	if c, ok := t.cmds[cmd]; ok {
		return c.Exec(t.ctx, args)
	} else {
		return errors.New("no such command")
	}
}
