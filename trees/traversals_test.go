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

func makeTwoSameTrees() (*binarynode[int], *binarynode[int]) {
	return makeTree(), makeTree()
}

func makeTwoTreesDifferentStructures() (*binarynode[int], *binarynode[int]) {
	tree1 := &binarynode[int]{
		value: 1,
		left: &binarynode[int]{
			value: 2,
			left: &binarynode[int]{
				value: 3,
			},
		},
	}

	tree2 := &binarynode[int]{
		value: 1,
		left: &binarynode[int]{
			value: 2,
		},
		right: &binarynode[int]{
			value: 3,
		},
	}

	return tree1, tree2
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

func TestBFS(t *testing.T) {
	tree := makeTree()
	found := bfs[int](tree, 4)
	if !found {
		t.Fatalf("failed to find 4. Wanted true, actual false")
	}

	found = bfs[int](tree, 42)
	if found {
		t.Fatalf("searching for 42 returned true. Wanted false, actual true")
	}
}

func TestCompareBinaryTrees(t *testing.T) {
	a, b := makeTwoSameTrees()
	isSame := compare[int](a, b)

	if !isSame {
		t.Fatal("compare(a, b) same trees is false, wanted true")
	}

	a, b = makeTwoTreesDifferentStructures()
	isSame = compare[int](a, b)

	if isSame {
		t.Fatal("compare(a, b) different tree structure is false, wanted true")
	}
}
