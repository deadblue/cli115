package core

import (
	"dead.blue/cli115/container"
	"fmt"
	"github.com/deadblue/elevengo"
)

type Context struct {
	//
	Alive bool

	Agent *elevengo.Agent
	User  *elevengo.UserInfo

	Path *container.Stack
}

func (c *Context) Prompt() string {
	return fmt.Sprintf("%s@115:/ # ", c.User.Name)
}
