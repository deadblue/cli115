package context

import (
	"fmt"
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/core"
	"go.dead.blue/cli115/internal/app"
	"go.dead.blue/cli115/internal/pkg/aria2"
	"log"
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

	// Aria2 RPC client
	Aria2 *aria2.Client

	// File-system for remote storage
	Fs *RemoteFs
}

func (i *Impl) Startup() error {
	// Test aria2 rpc server
	if conf := i.Conf.Aria2; conf.Rpc {
		i.Aria2 = aria2.New(conf.Endpoint, conf.Token)
		if ver, err := i.Aria2.GetVersion(); err == nil {
			log.Printf("Aria2 version: %s", ver)
		} else {
			conf.Rpc = false
			log.Printf("Fail to connect aria2 RPC server: %s", err.Error())
		}
	}
	// TODO: Move 115 login into here
	return nil
}

func (i *Impl) Shutdown() error {
	return nil
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

func New(agent *elevengo.Agent, conf *app.Conf) (core.Context, error) {
	impl := &Impl{
		alive: true,
		// App config
		Conf: conf,
		// Agent
		Agent: agent,
		User:  agent.User(),
		// Remote file system
		Fs: NewFs(agent),
	}
	return impl, nil
}
