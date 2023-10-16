package queue

import (
	"testing"
)

func ensureEmpty[T any](t *testing.T, q *Queue[T]) {
	if q.Len() != 0 {
		t.Errorf("q.Len() = %d, want %d", q.Len(), 0)
	}

	_, hasValue := q.Peek()
	if hasValue {
		t.Errorf("q.Peek() has value, want false")
	}
}

func ensureSingleValue[T any](t *testing.T, q *Queue[T]) {
	if q.Len() != 1 {
		t.Errorf("q.Len() = %d, want %d", q.Len(), 1)
	}
}

func TestNew(t *testing.T) {
	q := Init[int]()
	ensureEmpty[int](t, q)
}

func TestEnqueue1(t *testing.T) {
	q := Init[int]()
	q.Enqueue(42)

	ensureSingleValue[int](t, q)
}

func TestDeque(t *testing.T) {
	q := Init[int]()
	ensureEmpty[int](t, q)
	q.Enqueue(42)
	ensureSingleValue[int](t, q)

	q.Enqueue(69)
	if q.Len() != 2 {
		t.Errorf("q.Len() = %d, want %d", q.Len(), 2)
	}
}
