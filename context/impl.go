package context

import (
	"fmt"
	"github.com/deadblue/elevengo"
)

/*
An implementation of core.Context.
*/
type Impl struct {
	// Flag to control the terminal lifecycle
	alive bool

	Agent *elevengo.Agent
	User  *elevengo.UserInfo

	// File-system for remote storage
	Fs *RemoteFs
	// Root directory.
	Root *DirNode
	// Current directory.
	Curr *DirNode
	// Files under current directory
	Files map[string]*elevengo.File
}

func (i *Impl) Prompt() string {
	return fmt.Sprintf("%s:%s/ # ", i.User.Name, i.Curr.Name)
}

func (i *Impl) Alive() bool {
	return i.alive
}

func (i *Impl) Die() {
	i.alive = false
}
