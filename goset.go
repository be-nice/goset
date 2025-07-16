package goset

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func FromSlice[T comparable](vals []T) *Set[T] {
	s := New[T]()

	for _, v := range vals {
		s.Add(v)
	}

	return s
}

func (s *Set[T]) Clone() *Set[T] {
	clone := New[T]()

	for k := range s.data {
		clone.Add(k)
	}

	return clone
}

func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
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

func (s *Set[T]) SymDiff(other *Set[T]) *Set[T] {
	diff := New[T]()

	for k := range s.data {
		if !other.Contains(k) {
			diff.Add(k)
		}
	}

	for k := range other.data {
		if !s.Contains(k) {
			diff.Add(k)
		}
	}

	return diff
}

func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	diff := New[T]()

	for k := range s.data {
		if !other.Contains(k) {
			diff.Add(k)
		}
	}

	return diff
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := New[T]()

	for k := range s.data {
		union.Add(k)
	}

	for k := range other.data {
		union.Add(k)
	}

	return union
}

func (s *Set[T]) Inter(other *Set[T]) *Set[T] {
	inter := New[T]()

	for k := range s.data {
		if other.Contains(k) {
			inter.Add(k)
		}
	}

	return inter
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}

	for k := range s.data {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}

func (s *Set[T]) IsProperSubset(other *Set[T]) bool {
	return s.IsSubset(other) && s.Len() != other.Len()
}

func (s *Set[T]) IsEqual(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}

	for k := range s.data {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}
