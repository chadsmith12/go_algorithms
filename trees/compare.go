package main

func compare[T comparable](a *binarynode[T], b *binarynode[T]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compare[T](a.left, b.left) && compare[T](a.right, b.right)
}
