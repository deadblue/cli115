package util

type _TreeNode struct {
	value    interface{}
	parent   *_TreeNode
	children map[string]*_TreeNode
}

type Tree struct {
	root  *_TreeNode
	curr  *_TreeNode
	depth int
}

func (t *Tree) Curr() interface{} {
	return t.curr.value
}

func (t *Tree) Depth() int {
	return t.depth
}

func (t *Tree) Children() (children map[string]interface{}) {
	children = make(map[string]interface{})
	if t.curr.children != nil {
		for id, node := range t.curr.children {
			children[id] = node.value
		}
	}
	return
}

func (t *Tree) Append(id string, value interface{}) *Tree {
	if t.curr.children == nil {
		t.curr.children = make(map[string]*_TreeNode)
	}
	t.curr.children[id] = &_TreeNode{
		value:  value,
		parent: t.curr,
	}
	return t
}

func (t *Tree) Backward(depth int) *Tree {
	// TODO
	return t
}

func (t *Tree) Forward(id ...string) *Tree {
	// TODO
	return t
}

func (t *Tree) Goto(path string) *Tree {
	// TODO
	return t
}

func NewTree(root interface{}) *Tree {
	node := &_TreeNode{
		value:  root,
		parent: nil,
	}
	return &Tree{
		root:  node,
		curr:  node,
		depth: 0,
	}
}
