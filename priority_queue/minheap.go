package minheap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	data []T
}

func New[T constraints.Ordered](capacity int) *MinHeap[T] {
	return &MinHeap[T]{
		data: make([]T, 0, capacity),
	}
}

func (h *MinHeap[T]) Append(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap[T]) Delete() T {
	out := h.data[0]
	if len(h.data) == 1 {
		h.data = h.data[:0]
		return out
	}

	fmt.Printf("removing item %v\n", out)
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[1:]
	fmt.Printf("new len of data is %d\n", len(h.data))

	h.heapifyDown(0)
	fmt.Printf("new fist item in data is %v\n", h.data[0])

	return out
}

func (h *MinHeap[T]) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parentIndex := h.parent(index)
	parentValue := h.data[parentIndex]
	currentValue := h.data[index]

	if parentValue > currentValue {
		h.data[index] = parentValue
		h.data[parentIndex] = currentValue
		h.heapifyUp(parentIndex)
	}
}

func (h *MinHeap[T]) heapifyDown(index int) {
	if index >= len(h.data) {
		return
	}

	leftIndex := h.left(index)
	rightIndex := h.right(index)
	fmt.Printf("left index: %d and right index: %d\n", leftIndex, rightIndex)

	if leftIndex >= len(h.data) {
		return
	}

	value := h.data[index]
	leftValue := h.data[leftIndex]
	if rightIndex >= len(h.data) {
		if value > leftValue {
			h.data[index] = leftValue
			h.data[leftIndex] = value
			h.heapifyDown(leftIndex)
			return
		}
	}
	rightValue := h.data[rightIndex]

	if leftValue > rightValue && value > rightValue {
		h.data[index] = rightValue
		h.data[rightIndex] = value
		h.heapifyDown(rightIndex)
	} else if rightValue > leftValue && value > leftValue {
		h.data[index] = leftValue
		h.data[leftIndex] = value
		h.heapifyDown(leftIndex)
	}
}

func (h *MinHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap[T]) left(i int) int {
	return 2*i + 1
}

func (h *MinHeap[T]) right(i int) int {
	return 2*i + 2
}
