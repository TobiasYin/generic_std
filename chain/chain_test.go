package chain

import (
	"fmt"
	"testing"

	. "github.com/TobiasYin/generic_std/container/result"
)

func f1() Result[A] {
	return Ok(A{})
}

type A struct{}

func (a A) f2(err bool) Result[B] {
	if err {
		return Err[B](fmt.Errorf("error f2"))
	}
	return Ok(B{})
}

type B struct{}

func (b B) f3(err bool) Result[int] {
	if err {
		return Err[int](fmt.Errorf("error"))
	}
	return Ok(123)
}

func TestChain(t *testing.T) {
	v := Call(func() int {
		return f1().Unwrap().f2(false).Unwrap().f3(false).Unwrap()
	})
	if v.IsErr() {
		t.Fatal(v)
	}
	if v.UnwrapOrNil() != 123 {
		t.Fatal(v)
	}

	v = Call(func() int {
		return f1().Unwrap().f2(true).Unwrap().f3(false).Unwrap()
	})
	if !v.IsErr() {
		t.Fatal(v)
	}
	if v.Err().Error() != "error f2" {
		t.Fatal(v.Err())
	}
}

func of1() (oA, error) {
	return oA{}, nil
}

type oA struct{}

func (a oA) of2(err bool) (oB, error) {
	if err {
		return oB{}, fmt.Errorf("error f2")
	}
	return oB{}, nil
}

type oB struct{}

func (b oB) of3(err bool) (int, error) {
	if err {
		return 0, fmt.Errorf("error")
	}
	return 123, nil
}

func regularDemo() (int, error) {
	a, err := of1()
	if err != nil {
		return 0, err
	}
	b, err := a.of2(false)
	if err != nil {
		return 0, err
	}
	return b.of3(false)
}

func newDemo() (int, error) {
	return Call(func() int { return f1().Unwrap().f2(true).Unwrap().f3(false).Unwrap() }).Unzip()
}
