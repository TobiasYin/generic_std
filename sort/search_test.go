package sort

import (
	"fmt"
	std_sort "sort"
	"testing"
)

func TestSearch(t *testing.T) {
	data := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println(Search(data, 5))
	fmt.Println(std_sort.SearchInts(data, 5))

	{
		data := []float64{1, 3, 5, 7, 9, 11, 13}
		fmt.Println(Search(data, 5))
		fmt.Println(std_sort.SearchFloat64s(data, 5))
	}
}
