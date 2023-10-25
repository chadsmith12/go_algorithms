package stack

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	value    T
	previous *Node[T]
}

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func New[T any]() *Stack[T] {
	stack := &Stack[T]{
		length: 0,
		head:   nil,
	}

	return stack
}

func (s *Stack[T]) String() string {
	if s.length == 0 {
		return "[]"
	}

	builder := strings.Builder{}
	builder.WriteByte('[')
	currentNode := s.head
	for currentNode != nil {
		fmt.Fprintf(&builder, "%v", currentNode.value)
		if currentNode.previous != nil {
			builder.WriteByte(' ')
		}
		currentNode = currentNode.previous
	}

	builder.WriteByte(']')

	return builder.String()
}

func (s *Stack[T]) Push(item T) {
	node := &Node[T]{
		value: item,
	}
	s.length++
	if s.head == nil {
		s.head = node
		return
	}

	node.previous = s.head
	s.head = node
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.length == 0 {
		return *new(T), false
	}

	s.length--
	if s.length == 0 {
		head := s.head
		s.head = nil

		return head.value, true
	}

	head := s.head
	s.head = head.previous

	return head.value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.length == 0 {
		zeroValue := *new(T)
		return zeroValue, false
	}

	return s.head.value, true
}
