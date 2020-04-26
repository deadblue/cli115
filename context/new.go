package context

import (
	"github.com/deadblue/elevengo"
	"go.dead.blue/cli115/core"
)

func New(agent *elevengo.Agent) (core.Context, error) {
	impl := &Impl{
		alive: true,
		// Agent
		Agent: agent,
		User:  agent.User(),
		// Remote file system
		Fs: NewFs(agent),
	}
	return impl, nil
}
