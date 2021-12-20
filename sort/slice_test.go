package sort

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	data := []int{9, 4, 1, 2, 8, 6, 5}
	SliceForBasic(data)
	fmt.Println(data)
}
