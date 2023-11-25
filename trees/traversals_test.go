package main

import (
	"testing"
)

func testSliceSame(expected []int, actual []int, t *testing.T) {
	if len(expected) != len(actual) {
		t.Fatalf("len(expected) != len(actual) - got %d, want %d", len(actual), len(expected))
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("expected[i] != actual[i] - got %d, want %d", actual[i], expected[i])
		}
	}
}

func makeTree() *binarynode[int] {
	head := &binarynode[int]{
		value: 1,
		left: &binarynode[int]{
			value: 2,
			left: &binarynode[int]{
				value: 3,
			},
			right: &binarynode[int]{
				value: 4,
			},
		},
		right: &binarynode[int]{
			value: 5,
			left: &binarynode[int]{
				value: 6,
			},
			right: &binarynode[int]{
				value: 7,
			},
		},
	}

	return head
}

func TestPreOrderTraversal(t *testing.T) {
	tree := makeTree()
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	actual := preOrderSearch(tree)

	testSliceSame(expected, actual, t)
}

func TestInOrderTraversal(t *testing.T) {
	tree := makeTree()
	expected := []int{3, 2, 4, 1, 6, 5, 7}
	actual := inOrderSearch[int](tree)

	testSliceSame(expected, actual, t)
}

func TestPostOrderTraversal(t *testing.T) {
	tree := makeTree()
	expected := []int{3, 4, 2, 6, 7, 5, 1}
	actual := postOrderSearch[int](tree)

	testSliceSame(expected, actual, t)
}