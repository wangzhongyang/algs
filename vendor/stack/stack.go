package stack

import "fmt"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

func New() *Stack {
	return &Stack{}
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Return the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (interface{}, bool) {
	var value interface{}
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value, true
	}
	return nil, false
}

// Foreach stack
func (s *Stack) Foreach() {
	for s.size > 0 {
		value := s.top.value
		s.top = s.top.next
		fmt.Println(value)
		s.size--
	}
}
