package context

import (
	"github.com/deadblue/elevengo"
	"strings"
)

/*
A file system wrapper for all file operations on cloud storage.
*/
type RemoteFs struct {
	agent *elevengo.Agent

	// root dir
	root *DirNode
	// current dir
	curr *DirNode
	// files cache under current dir
	files map[string]*elevengo.File
}

func (f *RemoteFs) SetCurr(dir *DirNode) {
	f.curr = dir
	// TODO: update file caches
}

func (f *RemoteFs) GetCurr() *DirNode {
	return f.curr
}

/*
Locate directory by path.
"path" starts with "/" means an absolute path, otherwise a relative path.
*/
func (f *RemoteFs) LocateDir(path string) (dir *DirNode) {
	dir = f.curr
	dirs := strings.Split(path, "/")
	depth, start := len(dirs), 0
	if depth > 1 && dirs[0] == "" {
		// Starts from root
		dir = f.root
		start = 1
	}
	//
	for i := start; i < depth; i += 1 {
		if !dir.IsCached {
			f.createChildrenCache(dir)
		}
		dirName := dirs[i]
		if dirName == "." || dirName == "" {
			// "." means current dir
			continue
		} else if dirName == ".." {
			// ".." means upstairs dir
			if dir != f.root {
				dir = dir.Parent
			}
		} else {
			dir = dir.Children[dirName]
		}
		if dir == nil {
			break
		} else {
			if !dir.IsCached {
				f.createChildrenCache(dir)
			}
		}
	}
	return dir
}

func (f *RemoteFs) createChildrenCache(dir *DirNode) {
	for cur := elevengo.FileCursor(); cur.HasMore(); cur.Next() {
		if files, err := f.agent.FileList(dir.Id, cur); err != nil {
			break
		} else {
			for _, file := range files {
				if !file.IsDirectory {
					continue
				}
				dir.Append(file.FileId, file.Name)
			}
		}
	}
	dir.IsCached = true
}

// Get a file from current directory with specific
// name, or return nil when not found.
func (f *RemoteFs) File(name string) {

}

func NewFs(agent *elevengo.Agent) *RemoteFs {
	root := MakeNode("0", "")
	return &RemoteFs{
		agent: agent,
		root:  root,
		curr:  root,
	}
}
