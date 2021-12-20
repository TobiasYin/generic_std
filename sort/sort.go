package sort

import (
	std_sort "sort"

	"github.com/TobiasYin/generic_std/constraint"
)

type Slice[T constraint.BasicOrder] []T

func (x Slice[T]) Len() int           { return len(x) }
func (x Slice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x Slice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x Slice[T]) Sort()              { std_sort.Sort(x) }
func (x Slice[T]) IsSorted()          { SliceIsSortedForBasic(x) }

func IsSorted[T constraint.BasicOrder](x []T) bool { return std_sort.IsSorted(Slice[T](x)) }

type OrderSlice[T constraint.Order[T]] []T

func (x OrderSlice[T]) Len() int           { return len(x) }
func (x OrderSlice[T]) Less(i, j int) bool { return x[i].Less(x[j]) }
func (x OrderSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x OrderSlice[T]) Sort()              { std_sort.Sort(x) }
func (x OrderSlice[T]) IsSorted()          { SliceIsSortedForOrder(x) }

func IsSortedForOrder[T constraint.Order[T]](x []T) bool { return std_sort.IsSorted(OrderSlice[T](x)) }
