package sort

import (
	std_sort "sort"

	"github.com/TobiasYin/generic_std/constraint"
)

func Search[T constraint.BasicOrder](a []T, x T) int {
	return std_sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func SearchForOrder[T constraint.Order[T]](a []T, x T) int {
	return std_sort.Search(len(a), func(i int) bool { return !a[i].Less(x) })
}
