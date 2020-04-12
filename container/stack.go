package container

type _StackNode struct {
	value interface{}
	prev  *_StackNode
}

type Stack struct {
	size int
	curr *_StackNode
}

/*
Push
*/
func (s *Stack) Push(value interface{}) {
	node := &_StackNode{value: value, prev: nil}
	if s.curr == nil {
		s.curr = node
	} else {
		node.prev, s.curr = s.curr, node
	}
	s.size += 1
}

/*
Get the value of top node, and remove the node from stack.
*/
func (s *Stack) Pop() (value interface{}) {
	if s.curr != nil {
		value, s.curr = s.curr.value, s.curr.prev
		s.size -= 1
	}
	return
}

/*
Get the value of top node without removing the node.
*/
func (s *Stack) Top() (value interface{}) {
	if s.curr != nil {
		value = s.curr.value
	}
	return
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func NewStack() *Stack {
	return &Stack{}
}
