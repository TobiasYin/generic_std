package chain

import (
	"fmt"

	. "github.com/TobiasYin/generic_std/container/result"
)

func Call[T any](f func() T) (res Result[T]) {
	defer func() {
		if e := recover(); e != nil {
			var err error
			if er, ok := e.(error); ok {
				err = er
			} else {
				err = fmt.Errorf("panic: %v", e)
			}
			res = Err[T](err)
		}
	}()
	v := f()
	res = Ok(v)
	return
}
