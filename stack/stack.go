package stack

import "errors"

var ErrEmptyStack = errors.New("empty Stack")

type stack[T comparable] struct {
	size int64
	top  *Node[T]
}

func NewStack[T comparable]() Stacker[T] {
	return &stack[T]{}
}

func (s *stack[T]) Size() int64 {
	return s.size
}

func (s *stack[T]) Find(value T) bool {
	if s.size == 0 {
		return false
	}
	node := s.top
	for node.value != value {
		node = node.next
		if node == nil {
			return false
		}
	}
	return true
}

func (s *stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}

	value := s.top.value
	s.top = s.top.next
	s.size = s.size - 1
	return value, nil
}

func (s *stack[T]) Push(value T) {
	s.top = &Node[T]{value, s.top}
	s.size = s.Size() + 1
}

func (s *stack[T]) Top() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}
	return s.top.value, nil
}

func (s *stack[T]) Button() (T, error) {
	if s.size == 0 {
		var t T
		return t, ErrEmptyStack
	}
	node := s.top
	botton := node

	for node != nil {
		botton = node
		node = node.next
	}

	return botton.value, nil
}
