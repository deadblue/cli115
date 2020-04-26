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

	// Remote API agent
	Agent *elevengo.Agent
	User  *elevengo.UserInfo

	// File-system for remote storage
	Fs *RemoteFs
}

func (i *Impl) Prompt() string {
	return fmt.Sprintf("%s:%s/ # ", i.User.Name, i.Fs.curr.Name)
}

func (i *Impl) Alive() bool {
	return i.alive
}

func (i *Impl) Die() {
	i.alive = false
}
