package heap

import (
	std_heap "container/heap"

	"github.com/TobiasYin/generic_std/constraint"
)

type heapContainer struct {
	data []interface{}
	less func(d1 interface{}, d2 interface{}) bool
}

func (h *heapContainer) Push(x interface{}) {
	h.data = append(h.data, x)
}

func (h *heapContainer) Pop() interface{} {
	d := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return d
}

func (h *heapContainer) Len() int {
	return len(h.data)
}

func (h *heapContainer) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *heapContainer) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

type Heap[T any] interface {
	Push(x T)
	Pop() T
	Remove(i int) T
	Fix(i int)
	Len() int
}

func NewForOrder[T constraint.Order[T]]() Heap[T] {
	return &heap[T]{container: &heapContainer{
		data: nil,
		less: func(d1 interface{}, d2 interface{}) bool {
			return d1.(T).Less(d2.(T))
		},
	}}
}

func New[T constraint.BasicOrder]() Heap[T] {
	return &heap[T]{container: &heapContainer{
		data: nil,
		less: func(d1 interface{}, d2 interface{}) bool {
			return d1.(T) < d2.(T)
		},
	}}
}

type heap[T any] struct {
	container *heapContainer
}

func (h *heap[T]) Push(x T) {
	std_heap.Push(h.container, x)
}

func (h *heap[T]) Pop() T {
	return std_heap.Pop(h.container).(T)
}

func (h *heap[T]) Remove(i int) T {
	return std_heap.Remove(h.container, i).(T)
}

func (h *heap[T]) Fix(i int) {
	std_heap.Fix(h.container, i)
}
func (h *heap[T]) Len() int {
	return h.container.Len()
}
