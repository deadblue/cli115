package context

import (
	"fmt"
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/internal/app/conf"
	"go.dead.blue/cli115/internal/pkg/aria2"
	"go.dead.blue/cli115/internal/pkg/fs"
	"go.dead.blue/cli115/internal/pkg/terminal"
	"log"
)

/*
An implementation of terminal.Context.
*/
type Impl struct {
	// Flag to control the terminal lifecycle
	alive bool

	// options and conf
	Opts *conf.Options
	Conf *conf.Conf

	// Login user name
	user string

	// Remote API agent
	Agent *elevengo.Agent

	// File-system for remote storage
	Fs *fs.RemoteFs

	// Aria2 RPC client
	Aria2 *aria2.Client
}

func (i *Impl) Startup() error {
	// Init elevengo agent
	if agent, err := initAgent(i.Opts); err != nil {
		return err
	} else {
		i.Agent = agent
		i.user = agent.User().Name
		i.Fs = fs.New(agent)
	}
	// Test aria2 rpc server
	if cnf := i.Conf.Aria2; cnf.Rpc {
		i.Aria2 = aria2.New(cnf.Endpoint, cnf.Token)
		if ver, err := i.Aria2.GetVersion(); err == nil {
			log.Printf("Aria2 version: %s", ver)
		} else {
			cnf.Rpc = false
			log.Printf("Fail to connect aria2 RPC server: %s", err.Error())
		}
	}
	return nil
}

func (i *Impl) Shutdown() error {
	return nil
}

func (i *Impl) Prompt() string {
	return fmt.Sprintf("%s:%s/ # ", i.user, i.Fs.Curr().Name)
}

func (i *Impl) Alive() bool {
	return i.alive
}

func (i *Impl) Die() {
	i.alive = false
}

func New(opts *conf.Options, conf *conf.Conf) terminal.Context {
	return &Impl{
		alive: true,
		// App config
		Opts: opts,
		Conf: conf,
	}
}
