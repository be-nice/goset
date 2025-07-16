package goset

import (
	"reflect"
	"testing"
)

func TestGoSet(t *testing.T) {
	type testCase[T comparable] struct {
		name       string
		operations func(s *Set[T])
		expected   []T
		contains   map[T]bool
	}

	tests := []testCase[string]{
		{
			name: "Add single element",
			operations: func(s *Set[string]) {
				s.Add("apple")
			},
			expected: []string{"apple"},
			contains: map[string]bool{"apple": true},
		},
		{
			name: "Add and remove",
			operations: func(s *Set[string]) {
				s.Add("banana")
				s.Del("banana")
			},
			expected: []string{},
			contains: map[string]bool{"banana": false},
		},
		{
			name: "Multiple adds",
			operations: func(s *Set[string]) {
				s.Add("a")
				s.Add("b")
			},
			expected: []string{"a", "b"},
			contains: map[string]bool{"a": true, "b": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New[string]()
			tt.operations(s)

			got := s.Values()
			if !equalUnordered(got, tt.expected) {
				t.Errorf("expected values %v, got %v", tt.expected, got)
			}

			for val, expected := range tt.contains {
				if s.Contains(val) != expected {
					t.Errorf("Has(%v) = %v; want %v", val, s.Contains(val), expected)
				}
			}
		})
	}
}

func equalUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	setA := make(map[T]int)
	setB := make(map[T]int)
	for _, v := range a {
		setA[v]++
	}
	for _, v := range b {
		setB[v]++
	}
	return reflect.DeepEqual(setA, setB)
}
