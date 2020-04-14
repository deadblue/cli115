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

type CompletePhase int

const (
	phaseName CompletePhase = iota
	phaseArgs
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

func (t *Terminal) handle(line string) (err error) {
	name, args := line, ""
	pos := strings.IndexRune(line, ' ')
	if pos > 0 {
		name = strings.TrimSpace(line[:pos])
		args = strings.TrimSpace(line[:pos+1])
	}
	if c, ok := t.cmds[name]; !ok {
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

func (t *Terminal) Completer(line string) (choices []string) {
	// parse line
	np, ap, phase := line, "", phaseName
	pos := strings.IndexRune(line, ' ')
	if pos > 0 {
		phase = phaseArgs
		np = strings.TrimSpace(line[:pos])
		ap = strings.TrimSpace(line[:pos+1])
	}
	// fill choices
	if phase == phaseName {
		// Search command names
		choices = make([]string, 0)
		for name, cmd := range t.cmds {
			if !strings.HasPrefix(name, np) {
				continue
			}
			if cmd.HasArgs() {
				choices = append(choices, name+" ")
			} else {
				choices = append(choices, name)
			}
		}
	} else if phase == phaseArgs {
		// Call command complete
		cmd, ok := t.cmds[np]
		if !ok {
			return
		}
		cc, ok := cmd.(CommandCompleter)
		if !ok {
			return
		}
		return cc.Completer(t.ctx, ap)
	}
	return
}
