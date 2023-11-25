package main

type binarynode[T comparable] struct {
	value T
	left  *binarynode[T]
	right *binarynode[T]
}

func preOrderSearch[T comparable](head *binarynode[T]) []T {
	path := make([]T, 0)
	path = walkPreOrder[T](head, path)

	return path
}

func walkPreOrder[T comparable](head *binarynode[T], path []T) []T {
	if head == nil {
		return path
	}

	path = append(path, head.value)
	path = walkPreOrder[T](head.left, path)
	path = walkPreOrder[T](head.right, path)

	return path
}

func inOrderSearch[T comparable](head *binarynode[T]) []T {
	path := make([]T, 0)
	path = walkInOrder[T](head, path)

	return path
}

func walkInOrder[T comparable](head *binarynode[T], path []T) []T {
	if head == nil {
		return path
	}

	path = walkInOrder[T](head.left, path)
	path = append(path, head.value)
	path = walkInOrder[T](head.right, path)

	return path
}

func postOrderSearch[T comparable](head *binarynode[T]) []T {
	path := make([]T, 0)
	path = walkPostOrder[T](head, path)

	return path
}

func walkPostOrder[T comparable](head *binarynode[T], path []T) []T {
	if head == nil {
		return path
	}

	path = walkPostOrder[T](head.left, path)
	path = walkPostOrder[T](head.right, path)
	path = append(path, head.value)

	return path
}
