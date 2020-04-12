package core

import (
	"dead.blue/cli115/container"
	"github.com/deadblue/elevengo"
)

type Context struct {
	Agent  *elevengo.Agent
	Path   *container.Stack
	Prefix string
}

func New() *Context {
	return &Context{
		Agent:  elevengo.Default(),
		Path:   container.NewStack(),
		Prefix: "115",
	}
}
