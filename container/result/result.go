package result

type Result[T any] struct {
	err  error
	item T
}

func New[T any](val T, err error) Result[T] {
	return Result[T]{err, val}
}

func Ok[T any](val T) Result[T] {
	return Result[T]{nil, val}
}

func Err[T any](err error) Result[T] {
	if err == nil {
		panic("can't create Err result use nil error")
	}
	return Result[T]{err, *new(T)}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(r.err)
	}
	return r.item
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) UnwrapOrElse(e func(err error) T) T {
	if r.IsErr() {
		return e(r.Err())
	}
	return r.item
}

func (r Result[T]) UnwrapOrDefault(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.item
}

func (r Result[T]) UnwrapOrNil() T {
	return r.item
}

func (r Result[T]) UnwrapWithOk() (T, bool) {
	return r.UnwrapOrNil(), r.IsOk()
}

func (r Result[T]) Unzip() (T, error) {
	return r.UnwrapOrNil(), r.Err()
}
