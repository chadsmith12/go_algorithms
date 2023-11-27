package bst

type Comparator[T any] func(a, b T) int

type BinaryNode[T any] struct {
	value  T
	left   *BinaryNode[T]
	right  *BinaryNode[T]
	parent *BinaryNode[T]
}

func newBinaryNode[T any](value T, parent *BinaryNode[T]) *BinaryNode[T] {
	return &BinaryNode[T]{
		value:  value,
		parent: parent,
	}
}

func findMin[T any](n *BinaryNode[T]) *BinaryNode[T] {
	current := n
	for current.left != nil {
		current = current.left
	}

	return current
}

type BinarySearchTree[T any] struct {
	root       *BinaryNode[T]
	lenth      int
	comparator Comparator[T]
}

func (t *BinarySearchTree[T]) isEqual(left, right T) bool {
	return t.comparator(left, right) == 0
}

func (t *BinarySearchTree[T]) isGreaterThan(left, right T) bool {
	return t.comparator(left, right) > 0
}

func (t *BinarySearchTree[T]) isLessThan(left, right T) bool {
	return t.comparator(left, right) < 0
}

func NewBST[T any](c Comparator[T]) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		comparator: c,
	}
}

// Returns a slice of the tree using Inorder traversal
func (t *BinarySearchTree[T]) InorderSlice() []T {
	path := make([]T, 0, t.lenth)
	path = t.walk(t.root, path)

	return path
}

func (t *BinarySearchTree[T]) walk(node *BinaryNode[T], path []T) []T {
	if node == nil {
		return path
	}

	path = t.walk(node.left, path)
	path = append(path, node.value)
	path = t.walk(node.right, path)

	return path
}

// Inserts the value T into the Binary Search Tree and increases the length of the tre
func (t *BinarySearchTree[T]) Insert(value T) {
	t.lenth++

	if t.root == nil {
		node := newBinaryNode(value, nil)
		t.root = node
		return
	}

	t.insert(t.root, value)
}

func (t *BinarySearchTree[T]) insert(node *BinaryNode[T], value T) {
	isGreater := t.isGreaterThan(value, node.value)

	if isGreater {
		if node.right == nil {
			node.right = newBinaryNode(value, node)
			return
		}
		t.insert(node.right, value)
	} else {
		if node.left == nil {
			node.left = newBinaryNode(value, node)
			return
		}
		t.insert(node.left, value)
	}
}

// Uses DFS to attempt to find the value inside the BST
func (t *BinarySearchTree[T]) Find(value T) bool {
	node := t.find(t.root, value)
	return node != nil
}

func (t *BinarySearchTree[T]) findNode(value T) *BinaryNode[T] {
	return t.find(t.root, value)
}

func (t *BinarySearchTree[T]) find(node *BinaryNode[T], value T) *BinaryNode[T] {
	if node == nil {
		return nil
	}

	if t.isEqual(node.value, value) {
		return node
	}

	isGreater := t.isGreaterThan(value, node.value)
	if isGreater {
		return t.find(node.right, value)
	}

	return t.find(node.left, value)
}

func (t *BinarySearchTree[T]) Delete(value T) {
	t.deleteNode(t.root, value)
}

func (t *BinarySearchTree[T]) deleteNode(node *BinaryNode[T], value T) *BinaryNode[T] {
	if node == nil {
		return nil
	}

	if t.isLessThan(value, node.value) {
		node.left = t.deleteNode(node.left, value)
		return node
	}
	if t.isGreaterThan(value, node.value) {
		node.right = t.deleteNode(node.right, value)
		return node
	}

	t.lenth--
	return t.delete(node, value)
}

func (t *BinarySearchTree[T]) delete(node *BinaryNode[T], value T) *BinaryNode[T] {
	if node.left == nil && node.right == nil {
		if node == t.root {
			t.root = nil
		}
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	temp := findMin(node.right)
	node.value = temp.value
	node.right = t.delete(node.right, node.value)

	return node

}
