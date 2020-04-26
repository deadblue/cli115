package context

/*
Directory node
*/
type DirNode struct {
	// Unique ID
	Id string
	// Display name
	Name string
	// Depth from root
	Depth int

	// Parent node
	Parent *DirNode
	// Children nodes
	Children map[string]*DirNode
	// Is children cached
	ChildrenCached bool
}

// Append child to this node, do not replace exists entry.
func (n *DirNode) Append(id, name string) *DirNode {
	// Do not replace existing one
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
		Id:    id,
		Name:  name,
		Depth: 0,

		Parent:         nil,
		Children:       make(map[string]*DirNode),
		ChildrenCached: false,
	}
}
