package funcgo

// List represents an immutable singly-linked list.
type List[T any] struct {
	head  T
	tail  *List[T]
	empty bool
}

// EmptyList creates an empty immutable list.
func EmptyList[T any]() List[T] {
	return List[T]{empty: true}
}

// Cons prepends a value to an existing list, returning a new list.
func Cons[T any](head T, tail List[T]) List[T] {
	return List[T]{head: head, tail: &tail, empty: false}
}

// Head returns the first element of the list. Panics if empty.
func (l List[T]) Head() T {
	if l.empty {
		panic("Head on empty list")
	}
	return l.head
}

// Tail returns the rest of the list. Panics if empty.
func (l List[T]) Tail() List[T] {
	if l.empty {
		panic("Tail on empty list")
	}
	return *l.tail
}

// IsEmpty checks if the list is empty.
func (l List[T]) IsEmpty() bool {
	return l.empty
}

// Map applies a function to each element, returning a new list.
func (l List[T]) Map(f func(T) T) List[T] {
	if l.empty {
		return EmptyList[T]()
	}
	return Cons(f(l.head), l.tail.Map(f))
}

// Filter keeps elements that satisfy the predicate, returning a new list.
func (l List[T]) Filter(f func(T) bool) List[T] {
	if l.empty {
		return EmptyList[T]()
	}
	if f(l.head) {
		return Cons(l.head, l.tail.Filter(f))
	}
	return l.tail.Filter(f)
}

// ToSlice converts the list to a slice.
func (l List[T]) ToSlice() []T {
	var result []T
	for curr := l; !curr.empty; curr = *curr.tail {
		result = append(result, curr.head)
	}
	return result
}
