package cli115

import (
	"dead.blue/cli115/context"
	"errors"
	"github.com/peterh/liner"
	"strings"
)

type Command interface {
	Name() string
	Exec(context *context.Context, args string) (err error)
}

type Terminal struct {
	state    *liner.State
	context  *context.Context
	commands map[string]Command
}

func (t *Terminal) Register(cmds ...Command) {
	if t.commands == nil {
		t.commands = make(map[string]Command)
	}
	for _, cmd := range cmds {
		t.commands[cmd.Name()] = cmd
	}
}

func (t *Terminal) Run() (err error) {
	for alive := true; alive; {
		if input, err := t.state.Prompt(t.context.Prefix); err != nil {
			alive = false
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
	if c, ok := t.commands[cmd]; ok {
		return c.Exec(t.context, args)
	} else {
		return errors.New("no such command")
	}
}
