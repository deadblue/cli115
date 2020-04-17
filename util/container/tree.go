package container

type TreeNode struct {
	Id    string
	Value interface{}

	Parent   *TreeNode
	Children map[string]*TreeNode
}

func (n *TreeNode) Append(id string, value interface{}) *TreeNode {
	n.Children[id] = &TreeNode{
		Id:     id,
		Value:  value,
		Parent: n,
	}
	return n
}

type Tree struct {
	// The tree root
	root *TreeNode
	//
	nodes map[string]*TreeNode
}

func (t *Tree) Resolve(path string) *Tree {
	// TODO
	return t
}

func NewTree(id string, value interface{}) *Tree {
	root := &TreeNode{
		Id:     id,
		Value:  value,
		Parent: nil,
	}
	return &Tree{
		root: root,
	}
}
