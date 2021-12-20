package sync

import (
	std_sync "sync"
)

type Pool[T any] struct {
	pool std_sync.Pool

	New func() T
}

// Put adds x to the pool.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}

// Get selects an arbitrary item from the Pool, removes it from the
// Pool, and returns it to the caller.
// Get may choose to ignore the pool and treat it as empty.
// Callers should not assume any relation between values passed to Put and
// the values returned by Get.
//
// If Get would otherwise return nil and p.New is non-nil, Get returns
// the result of calling p.New.
func (p *Pool[T]) Get() (value T) {
	data := p.pool.Get()
	if data == nil {
		if p.New != nil {
			return p.New()
		}
		return value
	}
	return data.(T)
}
