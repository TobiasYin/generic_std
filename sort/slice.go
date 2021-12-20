package sort

import (
	std_sort "sort"

	"github.com/TobiasYin/generic_std/constraint"
)

func SliceForBasic[T constraint.BasicOrder](x []T) {
	std_sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
}

func SliceForOrder[T constraint.Order[T]](x []T) {
	std_sort.Slice(x, func(i, j int) bool {
		return x[i].Less(x[j])
	})
}

func SliceStable[T any](x []T, less func(i, j int) bool) {
	std_sort.SliceStable(x, less)
}

func SliceStableForBasic[T constraint.BasicOrder](x []T) {
	std_sort.SliceStable(x, func(i, j int) bool {
		return x[i] < x[j]
	})
}

func SliceStableForOrder[T constraint.Order[T]](x []T) {
	std_sort.SliceStable(x, func(i, j int) bool {
		return x[i].Less(x[j])
	})
}

func SliceIsSorted[T any](x []T, less func(i, j int) bool) bool {
	return std_sort.SliceIsSorted(x, less)
}

func SliceIsSortedForBasic[T constraint.BasicOrder](x []T) bool {
	return std_sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i] < x[j]
	})
}

func SliceIsSortedForOrder[T constraint.Order[T]](x []T) bool {
	return std_sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i].Less(x[j])
	})
}
