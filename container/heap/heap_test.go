package heap

import (
	"fmt"
	"testing"
)

type demo struct {
	data int
}

func (d demo) Less(b demo) bool {
	return d.data < b.data
}

func TestHeap(t *testing.T) {
	heap := New[int]()
	heap.Push(1)
	heap.Push(3)
	heap.Push(2)
	heap.Push(4)
	for heap.Len() > 0 {
		fmt.Println(heap.Pop())
	}

	heap2 := NewForOrder[demo]()
	heap2.Push(demo{4})
	heap2.Push(demo{2})
	heap2.Push(demo{5})
	heap2.Push(demo{3})
	for heap2.Len() > 0 {
		fmt.Println(heap2.Pop())
	}
}
