# goset

A simple wrapper around golangs map[T]struct{} for easier set methods  
Accepts data types that satisfy the builtin comparable generic.

## how to use

```go
import "github.com/be-nice/goset"
```

**initialize a new set**

```go
s := goset.New[type]()
```

#### methods

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
