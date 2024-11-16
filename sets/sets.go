package sets

type Set[T comparable] map[T]struct{}

func Make[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Adds(elements ...T) {
	for _, v := range elements {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Deletes(elements ...T) {
	for _, v := range elements {
		delete(s, v)
	}
}

func (s Set[T]) Contain(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}
