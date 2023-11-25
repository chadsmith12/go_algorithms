package main

import "github.com/chadsmith12/go_algorithms/queue"

func bfs[T comparable](head *binarynode[T], needle T) bool {
	q := queue.Init[*binarynode[T]]()
	q.Enqueue(head)

	for q.Len() != 0 {
		current, _ := q.Deque()
		if current.value == needle {
			return true
		}

		if current.left != nil {
			q.Enqueue(current.left)
		}
		if current.right != nil {
			q.Enqueue(current.right)
		}
	}

	return false
}
