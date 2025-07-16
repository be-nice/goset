package goset

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

func (s *Set[T]) Del(val T) {
	delete(s.data, val)
}

func (s *Set[T]) Contains(val T) bool {
	_, isInSet := s.data[val]
	return isInSet
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) Values() []T {
	items := make([]T, 0, len(s.data))

	for item := range s.data {
		items = append(items, item)
	}

	return items
}
