package core

import (
	"errors"
	"fmt"
	"github.com/peterh/liner"
	"go.dead.blue/cli115/util"
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
	ctx   *Context
	cmds  map[string]Command
}

/*
Register one or more commands into terminal.
*/
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
			t.handleErr(t.handle(input))
			t.state.AppendHistory(input)
		}
	}
	return nil
}

func (t *Terminal) handle(line string) (err error) {
	// Split input by space
	fields := util.SplitInput(line)
	if len(fields) == 0 {
		return
	}
	name, args := fields[0], fields[1:]
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

func (t *Terminal) completer(line string) (choices []string) {
	// parse line
	np, ap, phase := line, "", phaseName
	pos := strings.IndexRune(line, ' ')
	if pos > 0 {
		phase = phaseArgs
		np = strings.TrimSpace(line[:pos])
		ap = strings.TrimSpace(line[pos+1:])
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
		// Find command
		cmd, ok := t.cmds[np]
		if !ok {
			return
		}
		// Check whether command supports completer
		cc, ok := cmd.(CommandCompleter)
		if !ok {
			return
		}
		choices = cc.Completer(t.ctx, ap)
	}
	return
}

func NewTerminal(ctx *Context) *Terminal {
	// Create state
	state := liner.NewLiner()
	state.SetCtrlCAborts(true)
	state.SetTabCompletionStyle(liner.TabPrints)
	// Create terminal
	t := &Terminal{
		state: state,
		ctx:   ctx,
		cmds:  make(map[string]Command),
	}
	t.state.SetCompleter(t.completer)
	return t
}
