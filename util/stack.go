package util

type _StackNode struct {
	value interface{}
	prev  *_StackNode
}

type Stack struct {
	size int
	curr *_StackNode
}

/*
Push a value to the top of the stack.
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
Get value on the top of the stack, and remove it.
*/
func (s *Stack) Pop() (value interface{}) {
	if s.curr != nil {
		curr := s.curr
		value, s.curr = curr.value, curr.prev
		s.size -= 1
		curr.prev, curr = nil, nil
	}
	return
}

func (s *Stack) Values() []interface{} {
	values, i := make([]interface{}, s.size), s.size-1
	for node := s.curr; node != nil; node = node.prev {
		values[i] = node.value
		i -= 1
	}
	return values
}

/*
Get the value on the top.
*/
func (s *Stack) Top() (value interface{}) {
	if s.curr != nil {
		value = s.curr.value
	}
	return
}

/*
Clear the stack.
*/
func (s *Stack) Clear() {
	for s.curr != nil {
		curr := s.curr
		s.curr = curr.prev
		s.size -= 1
		curr.prev, curr = nil, nil
	}
}

/*
Reset the stack will new values, the last value in the list will be on the top.
*/
func (s *Stack) Reset(values ...interface{}) {
	s.Clear()
	for _, value := range values {
		s.Push(value)
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func NewStack() *Stack {
	return &Stack{}
}
