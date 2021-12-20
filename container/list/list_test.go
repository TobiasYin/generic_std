package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := New[int]()
	list.PushFront(1)
	list.PushFront(4)
	list.PushFront(6)
	list2 := New[int]()
	list2.PushFront(3)
	list2.PushFront(9)
	list2.PushFront(10)

	list.PushBackList(list2)
	list.PushBack(12)

	for list.Len() > 0 {
		f := list.Front()
		fmt.Println(f.Value)
		list.Remove(f)
	}
}
