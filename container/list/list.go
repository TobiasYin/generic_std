package list

import (
	std_list "container/list"
)

type Element[T any] struct {
	item *std_list.Element
}

func (e Element[T]) SetValue(v T) {
	e.item.Value = v
}

func (e Element[T]) GetValue() T {
	return e.item.Value.(T)
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	n := e.item.Next()
	return wrapElement[T](n)
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	n := e.item.Prev()
	return wrapElement[T](n)
}

func wrapElement[T any](e *std_list.Element) *Element[T] {
	if e == nil {
		return nil
	}
	return &Element[T]{item: e}
}

type List[T any] struct {
	list std_list.List
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] {
	l.list.Init()
	return l
}

// New returns an initialized list.
func New[T any]() *List[T] { return new(List[T]).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[T]) Len() int { return l.list.Len() }

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Element[T] {
	return wrapElement[T](l.list.Front())
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Element[T] {
	return wrapElement[T](l.list.Back())
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List[T]) Remove(e *Element[T]) T {
	v := l.list.Remove(e.item)
	return v.(T)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] {
	r := l.list.PushFront(v)
	return wrapElement[T](r)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] {
	r := l.list.PushBack(v)
	return wrapElement[T](r)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	r := l.list.InsertBefore(v, mark.item)
	return wrapElement[T](r)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	r := l.list.InsertAfter(v, mark.item)
	return wrapElement[T](r)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToFront(e *Element[T]) {
	l.list.MoveToFront(e.item)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToBack(e *Element[T]) {
	l.list.MoveToBack(e.item)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.list.MoveBefore(e.item, mark.item)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.list.MoveAfter(e.item, mark.item)
}

// PushBackList[T] inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[T]) PushBackList(other *List[T]) {
	l.list.PushBackList(&other.list)
}

// PushFrontList[T] inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) {
	l.list.PushFrontList(&other.list)
}
