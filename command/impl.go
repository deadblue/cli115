package command

import (
	"errors"
	"go.dead.blue/cli115/context"
	"go.dead.blue/cli115/core"
)

var (
	errIllegalContext = errors.New("illegal context")
)

type Impl interface {
	// The command name.
	Name() string
	// Does the command have arguments.
	HasArgs() bool
	// Execute command.
	ImplExec(ctx *context.Impl, args []string) error
}

type ImplCompleter interface {
	ImplCplt(ctx *context.Impl, index int, prefix string) (string, []string)
}

type wrapper struct {
	Impl
}

// Wrap the Exec method, convert the ctx into the type we want.
func (w *wrapper) Exec(ctx core.Context, args []string) error {
	if ictx, ok := ctx.(*context.Impl); ok {
		return w.Impl.ImplExec(ictx, args)
	} else {
		return errIllegalContext
	}
}

type wraperEx struct {
	wrapper
}

func (we *wraperEx) Completer(ctx core.Context, index int, prefix string) (string, []string) {
	if ictx, ok := ctx.(*context.Impl); !ok {
		// return empty choice for illegal context
		return "", []string{}
	} else {
		// We are sure that Impl in wrapperEx always implements ArgCompleter
		ic, _ := we.Impl.(ImplCompleter)
		return ic.ImplCplt(ictx, index, prefix)
	}
}

func Wrap(cmd Impl) core.Command {
	wp := wrapper{cmd}
	if _, ok := cmd.(ImplCompleter); ok {
		return &wraperEx{wp}
	} else {
		return &wp
	}
}
