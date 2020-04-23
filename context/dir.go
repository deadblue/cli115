package context

import (
	"strings"
)

type DirNode struct {
	// Unique ID
	Id string
	// Display name
	Name string

	// Is children cached
	IsCached bool

	Depth    int
	Parent   *DirNode
	Children map[string]*DirNode
}

// Return the full path of the node
func (n *DirNode) Path(sep string) string {
	buf, depth := make([]string, n.Depth+1), n.Depth
	for node := n; node != nil; node = node.Parent {
		buf[depth] = node.Name
		depth -= 1
	}
	if sep == "" {
		sep = "/"
	}
	path := strings.Join(buf, sep)
	if path == "" {
		path = "/"
	}
	return path
}

// Append child to this node, do not replace exists entry.
func (n *DirNode) Append(id, name string) *DirNode {
	if _, ok := n.Children[name]; !ok {
		node := MakeNode(id, name)
		node.Parent = n
		node.Depth = n.Depth + 1
		n.Children[name] = node
	}
	return n
}

func MakeNode(id, name string) *DirNode {
	return &DirNode{
		Id:   id,
		Name: name,

		Depth:    0,
		Parent:   nil,
		Children: make(map[string]*DirNode),
		IsCached: false,
	}
}
