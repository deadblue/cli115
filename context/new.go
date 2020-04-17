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
		// Directory tree
		Root: nil,
		Curr: nil,
		// File cache
		Files: make(map[string]*elevengo.File),
	}
	root := MakeNode("0", "")
	impl.Root, impl.Curr = root, root
	// Cache files and directory under root
	for cur := elevengo.FileCursor(); cur.HasMore(); cur.Next() {
		if files, err := agent.FileList(impl.Curr.Id, cur); err != nil {
			return nil, err
		} else {
			for _, file := range files {
				if file.IsDirectory {
					impl.Curr.Append(file.FileId, file.Name)
				}
				if file.IsFile {
					impl.Files[file.Name] = file
				}
			}
		}
	}
	return impl, nil
}
