package minheap

import (
	"testing"
)

func assertLen(h *MinHeap[int], expected int, t *testing.T) {
	length := len(h.data)

	if length != expected {
		t.Errorf("len(h.data) was %d, expected %d", length, expected)
	}
}

func assertFirstItem(h *MinHeap[int], exepcted int, t *testing.T) {
	item := h.data[0]

	if item != exepcted {
		t.Errorf("first item of data was %d, expected %d", item, exepcted)
	}
}

func assertDeletedItem(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("retrieved wrong item on deletion: got %d, expected %d", actual, expected)
	}
}

func TestMinHeap(t *testing.T) {
	heap := New[int](0)
	assertLen(heap, 0, t)

	heap.Append(10)
	assertLen(heap, 1, t)
	assertFirstItem(heap, 10, t)

	heap.Append(20)
	assertLen(heap, 2, t)
	assertFirstItem(heap, 10, t)

	heap.Append(1)
	assertLen(heap, 3, t)
	assertFirstItem(heap, 1, t)

	deletedItem := heap.Delete()
	assertDeletedItem(1, deletedItem, t)
	assertLen(heap, 2, t)

	deletedItem = heap.Delete()
	assertDeletedItem(10, deletedItem, t)
	assertLen(heap, 1, t)

	deletedItem = heap.Delete()
	assertDeletedItem(20, deletedItem, t)
	assertLen(heap, 0, t)
}
