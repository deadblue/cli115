package cli115

import (
	"errors"
	"fmt"
	"github.com/peterh/liner"
	"go.dead.blue/cli115/core"
	"strings"
)

var (
	errCommandNotExist = errors.New("no such command")
)

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
		if input, err := t.state.Prompt(t.ctx.PromptString()); err != nil {
			if err == liner.ErrPromptAborted {
				return err
			} else {
				t.handleErr(err)
			}
		} else {
			input = strings.TrimSpace(input)
			t.handleErr(t.handle(input))
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
	if c, ok := t.cmds[cmd]; !ok {
		return errCommandNotExist
	} else {
		return c.Exec(t.ctx, args)
	}
}

func (t *Terminal) handleErr(err error) {
	if err == nil {
		return
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func (t *Terminal) Completer(line string) []string {

	// TODO
	return nil
}
