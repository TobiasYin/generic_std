// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ring

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := New[int](3)
	r.SetValue(123)
	r1 := r.Next()
	r1.SetValue(234)
	r2 := r1.Next()
	r2.SetValue(567)
	r2.Next().SetValue(910)
	fmt.Println(*r.GetValue())
}
