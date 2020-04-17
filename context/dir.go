package context

import (
	"strings"
	"time"
)

type DirNode struct {
	// Unique ID
	Id string
	// Display name
	Name string
	// Node update time
	Time time.Time

	Depth    int
	Parent   *DirNode
	Children map[string]*DirNode
}

// Return the full path of the node
func (n *DirNode) Path(sep string) string {
	depth := n.Depth
	buf := make([]string, n.Depth+1)
	for node := n; node != nil; node = node.Parent {
		buf[depth] = node.Name
		depth -= 1
	}
	if sep == "" {
		sep = "/"
	}
	return strings.Join(buf, sep)
}

// Append children under current node.
func (n *DirNode) Append(id, name string) *DirNode {
	if _, ok := n.Children[name]; !ok {
		node := MakeNode(id, name)
		node.Parent = n
		node.Depth = n.Depth + 1
		n.Children[name] = node
		n.Time = time.Now()
	}
	return n
}

func MakeNode(id, name string) *DirNode {
	return &DirNode{
		Id:   id,
		Name: name,
		Time: time.Now(),

		Depth:    0,
		Parent:   nil,
		Children: make(map[string]*DirNode),
	}
}
