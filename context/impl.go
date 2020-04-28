package context

import (
	"fmt"
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/internal/app"
)

/*
An implementation of core.Context.
*/
type Impl struct {
	// Flag to control the terminal lifecycle
	alive bool

	// Config
	Conf *app.Conf

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
