package cli115

import (
	"dead.blue/cli115/core"
	"github.com/deadblue/elevengo"
)

func createContext(agent *elevengo.Agent) (ctx *core.Context, err error) {
	ctx = &core.Context{
		Alive: true,
		Agent: agent,
		User:  agent.User(),
		Path:  core.NewStack(),
		Cache: make(map[string]*elevengo.File),
	}
	// Cache files under root
	for cursor := elevengo.FileCursor(); cursor.HasMore(); cursor.Next() {
		files, err := agent.FileList("0", cursor)
		if err != nil {
			return nil, err
		} else {
			for _, file := range files {
				ctx.Cache[file.Name] = file
			}
		}
	}
	return
}
