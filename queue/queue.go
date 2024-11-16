package main

type Queue[T any] struct {
	data []T
}

func NewQueueWithValue[T any](init ...T) *Queue[T] {
	return &Queue[T]{data: init}
}

func (q *Queue[T]) Enqueue(value ...T) {
	q.data = append(q.data, value...)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	firstElement := q.data[0]
	q.data = q.data[1:]
	return firstElement
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return q.data[len(q.data)-1]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Length() int {
	return len(q.data)
}
