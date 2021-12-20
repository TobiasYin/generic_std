package ring

import (
	std_ring "container/ring"
)

type Ring[T any] struct {
	ring std_ring.Ring
}

func (r *Ring[T]) SetValue(v T) {
	r.ring.Value = v
}

func (r *Ring[T]) GetValue() *T {
	var v *T
	if r.ring.Value != nil {
		vd := r.ring.Value.(T)
		v = &vd
	}
	return v
}

func wrapRing[T any](e *std_ring.Ring) *Ring[T] {
	if e == nil {
		return nil
	}

	return &Ring[T]{ring: *e}
}

// Next returns the next ring element. r must not be empty.
func (r *Ring[T]) Next() *Ring[T] {
	return wrapRing[T](r.ring.Next())
}

// Prev returns the previous ring element. r must not be empty.
func (r *Ring[T]) Prev() *Ring[T] {
	return wrapRing[T](r.ring.Prev())
}

// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// in the ring and returns that ring element. r must not be empty.
//
func (r *Ring[T]) Move(n int) *Ring[T] {
	return wrapRing[T](r.ring.Move(n))
}

// New creates a ring of n elements.
func New[T any](n int) *Ring[T] {
	return wrapRing[T](std_ring.New(n))
}

// Link connects ring r with ring s such that r.Next()
// becomes s and returns the original value for r.Next().
// r must not be empty.
//
// If r and s point to the same ring, linking
// them removes the elements between r and s from the ring.
// The removed elements form a subring and the result is a
// reference to that subring (if no elements were removed,
// the result is still the original value for r.Next(),
// and not nil).
//
// If r and s point to different rings, linking
// them creates a single ring with the elements of s inserted
// after r. The result points to the element following the
// last element of s after insertion.
//
func (r *Ring[T]) Link(s *Ring[T]) *Ring[T] {
	return wrapRing[T](r.ring.Link(&s.ring))
}

// Unlink removes n % r.Len() elements from the ring r, starting
// at r.Next(). If n % r.Len() == 0, r remains unchanged.
// The result is the removed subring. r must not be empty.
//
func (r *Ring[T]) Unlink(n int) *Ring[T] {
	return wrapRing[T](r.ring.Unlink(n))
}

// Len computes the number of elements in ring r.
// It executes in time proportional to the number of elements.
//
func (r *Ring[T]) Len() int {
	return r.ring.Len()
}

// Do calls function f on each element of the ring, in forward order.
// The behavior of Do is undefined if f changes *r.
func (r *Ring[T]) Do(f func(T)) {
	r.ring.Do(func(a interface{}) {
		f(a.(T))
	})
}
