package linkedlist

import "fmt"

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

// Structure that represents a generic doubly-linked list.
type LinkedList[T any] struct {
	length int
	head   *node[T]
	tail   *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

// Creates a new empty linked list with the length of 0
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		length: 0,
		head:   nil,
	}
}

// Len returns the number of elements currently in the list
func (l *LinkedList[T]) Len() int {
	return l.length
}

// Prepends the item to the beginning of the linked list.
// This will be the new head of the linked list.
func (l *LinkedList[T]) Prepend(item T) {
	node := newNode[T](item)
	l.length++
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	// prepend to the start, so this new node will be the head.
	node.next = l.head
	l.head.prev = node
	l.head = node
}

// Inserts item at the index specified in the Linkedlist.
// If index is < 0 or > lenth of the linked list, it will return an IndexOutOfRangeError
// If the insert was successful then error will be nil
func (l *LinkedList[T]) InsertAt(item T, index int) error {
	if index < 0 || index > l.length {
		return newIndexOutOfRangeError(index)
	} else if index == 0 {
		l.Prepend(item)
		return nil
	} else if index == l.length {
		l.Append(item)
	}

	foundNode := l.head
	for i := index; i < index; i++ {
		foundNode = foundNode.next
	}

	node := newNode[T](item)
	l.length++
	node.next = foundNode
	node.prev = foundNode.prev
	foundNode.prev.next = node
	foundNode.prev = node

	return nil
}

// Appends the item to the end of the linked list
func (l *LinkedList[T]) Append(item T) {
	l.length++
	node := newNode[T](item)
	if l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) RemoveItem(item T, compareFunc func(a, b T) bool) {

}

// Structure that represents an error for when the index is out of range.
type IndexOutOfRangeError struct {
	index int
}

func newIndexOutOfRangeError(index int) *IndexOutOfRangeError {
	return &IndexOutOfRangeError{index: index}
}

func (e *IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("index %v is out of range.", e.index)
}
