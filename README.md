# goset

A simple wrapper around golangs map[T]struct{} for easier set methods  
Accepts data types that satisfy the builtin comparable generic.

## how to use

```go
import "github.com/be-nice/goset"
```

**Initialize a new set**

```go
s := goset.New[type]()
```

**Build set from slice**

```go
s := goset.FromSlice(slice)
```

**Make a clone from set A**

```go
s := a.Clone()
```

**Clear set**

```go
s.Clear()
```

### set methods

**add item**

```go
func (s *Set[T]) Add(val T)
```

**Remove item**

```go
func (s *Set[T]) Del(val T)
```

**Check if item is in the set**

```go
func (s *Set[T]) Contains(val T) bool
```

**Get number of items in the set**

```go
func (s *Set[T]) Len() int
```

**Return set values in a slice**

```go
func (s *Set[T]) Values() []T
```

**Return union of A and B**

```go
func (a *Set[T]) Union(b *Set[T]) *Set[T]
```

**Return intersection of A and B**

```go
func (a *Set[T]) Inter(b *Set[T]) *Set[T]
```

**Return difference of A and B**

```go
func (a *Set[T]) Diff(b *Set[T]) *Set[T]
```

**Return symmetric difference between A and B**

```go
func (a *Set[T]) SymDiff(b *Set[T]) *Set[T]
```

**Compare if set A and B are equal**

```go
func (a *Set[T]) IsEqual(b *Set[T]) bool
```

**Check if A is subset of B**

```go
func (a *Set[T]) IsSubset(b *Set[T]) bool
```

**Check if A is proper subset of B**

```go
func (a *Set[T]) IsProperSubset(b *Set[T]) bool
```
