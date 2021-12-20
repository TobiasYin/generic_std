package sync

import (
	std_sync "sync"
)

type Map[K any, V any] struct {
	realMap std_sync.Map
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.realMap.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), true
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.realMap.Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	ac, loaded := m.realMap.LoadOrStore(key, value)
	return ac.(V), loaded

}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.realMap.LoadAndDelete(key)
	if !loaded {
		return value, false
	}
	return v.(V), true
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.realMap.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.realMap.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
