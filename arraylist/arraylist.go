package arraylist

type ArrayList[T any] struct {
	data []T
}

func NewList[T any](capacity int) *ArrayList[T] {
	list := &ArrayList[T]{
		data: make([]T, 0, capacity),
	}

	return list
}

func (l *ArrayList[T]) Push(value T) {
	l.data = append(l.data, value)
}

func (l *ArrayList[T]) Pop() T {
	value := l.data[len(l.data)-1]
	l.data = l.data[0 : len(l.data)-1]

	return value
}

func (l *ArrayList[T]) Get(index int) T {
	return l.data[index]
}

func (l *ArrayList[T]) Len() int {
	return len(l.data)
}

func (l *ArrayList[T]) Cap() int {
	return cap(l.data)
}
