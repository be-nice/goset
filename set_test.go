package goset

import "testing"

func TestAddContains(t *testing.T) {
	type testCase struct {
		input    []int
		check    int
		expected bool
	}

	tests := []testCase{
		{input: []int{1, 2, 3}, check: 2, expected: true},
		{input: []int{1, 2, 3}, check: 4, expected: false},
	}

	for _, tc := range tests {
		t.Run("Contains", func(t *testing.T) {
			s := FromSlice(tc.input)
			got := s.Contains(tc.check)
			if got != tc.expected {
				t.Errorf("Contains(%v) = %v, want %v", tc.check, got, tc.expected)
			}
		})
	}
}

func TestDel(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	s.Del(2)
	if s.Contains(2) {
		t.Errorf("Expected 2 to be deleted")
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 1, 2}, 2},
	}

	for _, tc := range tests {
		s := FromSlice(tc.input)
		if s.Len() != tc.expected {
			t.Errorf("Len(%v) = %d, want %d", tc.input, s.Len(), tc.expected)
		}
	}
}

func TestClear(t *testing.T) {
	s := FromSlice([]int{1, 2})
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected empty set after Clear")
	}
}

func TestClone(t *testing.T) {
	orig := FromSlice([]int{1, 2})
	clone := orig.Clone()

	if !orig.IsEqual(clone) {
		t.Errorf("Clone mismatch: got %v, want %v", clone.Values(), orig.Values())
	}

	clone.Add(3)
	if orig.Contains(3) {
		t.Errorf("Clone modification affected original")
	}
}

func TestValues(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	vals := s.Values()
	if len(vals) != 3 {
		t.Errorf("Expected 3 values, got %d", len(vals))
	}
}

func TestUnion(t *testing.T) {
	a := FromSlice([]int{1, 2})
	b := FromSlice([]int{2, 3})
	expected := FromSlice([]int{1, 2, 3})

	result := a.Union(b)
	if !result.IsEqual(expected) {
		t.Errorf("Union = %v, want %v", result.Values(), expected.Values())
	}
}

func TestInter(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{2, 3, 4})
	expected := FromSlice([]int{2, 3})

	result := a.Inter(b)
	if !result.IsEqual(expected) {
		t.Errorf("Inter = %v, want %v", result.Values(), expected.Values())
	}
}

func TestDiff(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{2, 4})
	expected := FromSlice([]int{1, 3})

	result := a.Diff(b)
	if !result.IsEqual(expected) {
		t.Errorf("Diff = %v, want %v", result.Values(), expected.Values())
	}
}

func TestSymDiff(t *testing.T) {
	a := FromSlice([]int{1, 2, 3})
	b := FromSlice([]int{2, 4})
	expected := FromSlice([]int{1, 3, 4})

	result := a.SymDiff(b)
	if !result.IsEqual(expected) {
		t.Errorf("SymDiff = %v, want %v", result.Values(), expected.Values())
	}
}

func TestIsSubset(t *testing.T) {
	tests := []struct {
		name     string
		a, b     *Set[int]
		expected bool
	}{
		{"subset", FromSlice([]int{1, 2}), FromSlice([]int{1, 2, 3}), true},
		{"not subset", FromSlice([]int{1, 4}), FromSlice([]int{1, 2, 3}), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.a.IsSubset(tc.b); got != tc.expected {
				t.Errorf("IsSubset = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestIsProperSubset(t *testing.T) {
	tests := []struct {
		name     string
		a, b     *Set[int]
		expected bool
	}{
		{"proper subset", FromSlice([]int{1, 2}), FromSlice([]int{1, 2, 3}), true},
		{"equal", FromSlice([]int{1, 2}), FromSlice([]int{1, 2}), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.a.IsProperSubset(tc.b); got != tc.expected {
				t.Errorf("IsProperSubset = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     *Set[int]
		expected bool
	}{
		{"equal", FromSlice([]int{1, 2, 3}), FromSlice([]int{3, 2, 1}), true},
		{"not equal", FromSlice([]int{1, 2}), FromSlice([]int{1, 2, 3}), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.a.IsEqual(tc.b); got != tc.expected {
				t.Errorf("IsEqual = %v, want %v", got, tc.expected)
			}
		})
	}
}

// BENCHMARKS
func generateSet(size int) *Set[int] {
	s := New[int]()
	for i := 0; i < size; i++ {
		s.Add(i)
	}
	return s
}

func BenchmarkUnion(b *testing.B) {
	a := generateSet(10000)
	bSet := generateSet(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Union(bSet)
	}
}

func BenchmarkInter(b *testing.B) {
	a := generateSet(10000)
	bSet := generateSet(5000) // Partial overlap

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Inter(bSet)
	}
}

func BenchmarkDiff(b *testing.B) {
	a := generateSet(10000)
	bSet := generateSet(5000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Diff(bSet)
	}
}

func BenchmarkSymDiff(b *testing.B) {
	a := generateSet(10000)
	bSet := generateSet(5000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.SymDiff(bSet)
	}
}
