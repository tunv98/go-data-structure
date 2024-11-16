package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any](init ...T) *Stack[T] {
	return &Stack[T]{data: init}
}

func (s *Stack[T]) Push(value ...T) {
	s.data = append(s.data, value...)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	lastElement := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return lastElement
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Length() int {
	return len(s.data)
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return s.data[len(s.data)-1]
}
