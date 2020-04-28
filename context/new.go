package context

import (
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/core"
	"go.dead.blue/cli115/internal/app"
)

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
