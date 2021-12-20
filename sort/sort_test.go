// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{9, 8, 2, 1, 5, 3, 6, 7}
	sl := Slice[int](data)
	sl.Sort()
	fmt.Println(data)
}
