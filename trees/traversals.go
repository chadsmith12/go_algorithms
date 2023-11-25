package main

type binarynode[T any] struct {
	value T
	left  *binarynode[T]
	right *binarynode[T]
}

func preOrderSearch[T any](head *binarynode[T]) []T {
	path := make([]T, 0)
	path = walkPreOrder[T](head, path)

	return path
}

func walkPreOrder[T any](head *binarynode[T], path []T) []T {
	if head == nil {
		return path
	}

	path = append(path, head.value)
	path = walkPreOrder[T](head.left, path)
	path = walkPreOrder[T](head.right, path)

	return path
}

func inOrderSearch[T any](head *binarynode[T]) []T {
	path := make([]T, 0)
	path = walkInOrder[T](head, path)

	return path
}

func walkInOrder[T any](head *binarynode[T], path []T) []T {
	if head == nil {
		return path
	}

	path = walkInOrder[T](head.left, path)
	path = append(path, head.value)
	path = walkInOrder[T](head.right, path)

	return path
}

func postOrderSearch[T any](head *binarynode[T]) []T {
	path := make([]T, 0)
	path = walkPostOrder[T](head, path)

	return path
}

func walkPostOrder[T any](head *binarynode[T], path []T) []T {
	if head == nil {
		return path
	}

	path = walkPostOrder[T](head.left, path)
	path = walkPostOrder[T](head.right, path)
	path = append(path, head.value)

	return path
}
