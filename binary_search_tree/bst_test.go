package bst

import (
	"testing"

	"golang.org/x/exp/slices"
)

func IntComparator(a, b int) int {
	return a - b
}

func TestInsertingBST(t *testing.T) {
	tree := NewBST[int](IntComparator)
	if tree.root != nil {
		t.Fatalf("newly initialized tree is not empty. root is %v", tree.root.value)
	}

	tree.Insert(50)
	if tree.root == nil {
		t.Fatal("root is nil after inserting first node, expected 50")
	}

	tree.Insert(25)
	if tree.root == nil {
		t.Fatal("root is nil after inserting new leaf, expected 50")
	}
	if tree.root.left == nil {
		t.Fatalf("roots left node is nil after inserting new leaf less than root")
	}
	if tree.root.left.value != 25 {
		t.Fatalf("roots left node is %d, expected %d", tree.root.left.value, 25)
	}

	tree.Insert(75)
	if tree.root == nil {
		t.Fatal("root is nil after inserting new leaf, expected 50")
	}
	if tree.root.right == nil {
		t.Fatalf("roots right node is nil after inserting new leaf greater than root")
	}
	if tree.root.right.value != 75 {
		t.Fatalf("roots right node is %d, expected %d", tree.root.right.value, 75)
	}
}

func TestInOrder(t *testing.T) {
	tree := NewBST[int](IntComparator)
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)

	expected := []int{25, 50, 75}
	path := tree.InorderSlice()

	if !slices.IsSorted(path) {
		t.Fatalf("tree.InorderSlice() is not in sorted order, %v, expected %v", expected, path)
	}
}

func TestFind(t *testing.T) {
	tree := NewBST[int](IntComparator)
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)

	expected := true
	actual := tree.Find(25)
	if actual != expected {
		t.Fatal("tree.Find(25) expected to return true, got false")
	}

	expected = false
	actual = tree.Find(420)
	if actual != expected {
		t.Fatal("tree.Find(420) expected to return false, got true")
	}
}
