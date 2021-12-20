package atomic

import (
	std_atomic "sync/atomic"
)

type Value[T any] struct {
	v std_atomic.Value
}

func (v *Value[T]) Load() T {
	data := v.v.Load()
	if data == nil {
		return *new(T)
	}
	d := data.(T)
	return d
}

func (v *Value[T]) Store(val T) {
	v.v.Store(val)
}

func (v *Value[T]) Swap(new T) (old T) {
	r := v.v.Swap(new)
	if r == nil {
		return old
	}
	o := r.(T)
	return o
}

func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.v.CompareAndSwap(old, new)
}
