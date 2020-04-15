package core

import (
	"fmt"
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/util"
)

type Dir struct {
	Id   string
	Name string
}

type Context struct {
	// Flag to control the terminal lifecycle
	Alive bool

	Agent *elevengo.Agent
	User  *elevengo.UserInfo

	Path  *util.Stack
	Cache map[string]*elevengo.File
}

func (c *Context) PromptString() string {
	return fmt.Sprintf("%s:/ # ", c.User.Name)
}
