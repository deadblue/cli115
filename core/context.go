package core

import (
	"fmt"
	"github.com/deadblue/elevengo"
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

	Path  *Stack
	Cache map[string]*elevengo.File
}

func (c *Context) PromptString() string {
	return fmt.Sprintf("%s:/ # ", c.User.Name)
}
