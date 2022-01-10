package option

import "testing"

func TestOption(t *testing.T) {
	opt := New(123)
	v := opt.Unwrap()
	if v != 123 {
		t.Fatal()
	}

	n := Nil[int]()
	v = n.UnwrapOrDefault(3)
	if v != 3 {
		t.Fatal(v)
	}

	v = n.UnwrapOrElse(func() int {
		return 4
	})

	if v != 4 {
		t.Fatal()
	}
	shouldPanic(t, func() {
		n.Unwrap()
	})
}

func shouldPanic(t *testing.T, f func()) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal(":shouldPanic")
		}
	}()
	f()
}
