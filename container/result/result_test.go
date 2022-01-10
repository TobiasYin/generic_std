package result

import (
	"fmt"
	"testing"
)

func shouldPanic(t *testing.T, f func()) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal(":shouldPanic")
		}
	}()
	f()
}

func TestResult(t *testing.T) {
	shouldPanic(t, func() {
		Err[int](nil)
	})

	e := Err[int](fmt.Errorf("Err Test"))
	shouldPanic(t, func() {
		e.Unwrap()
	})
	v := e.UnwrapOrDefault(3)
	if v != 3 {
		t.Fatal(v)
	}

	v = e.UnwrapOrElse(func(err error) int {
		fmt.Println(err)
		return 4
	})
	if v != 4 {
		t.Fatal(v)
	}

	ok := Ok(123)

	v = ok.Unwrap()
	if v != 123 {
		t.Fatal(v)
	}

	if !ok.IsOk() {
		t.Fatal()
	}

	if ok.IsErr() {
		t.Fatal()
	}

	if e.IsOk() {
		t.Fatal()
	}

	if !e.IsErr() {
		t.Fatal()
	}

	if e.Err() == nil {
		t.Fatal()
	}
}
