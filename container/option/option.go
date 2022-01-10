package option

type Option[T any] struct {
	notNil bool // default nil
	item   T
}

func New[T any](data T) Option[T] {
	return Option[T]{true, data}
}

func Nil[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsNil() bool {
	return !o.notNil
}

func (o Option[T]) Unwrap() T {
	if o.IsNil() {
		panic("empty option")
	}
	return o.item
}

func (o Option[T]) UnwrapOrElse(e func() T) T {
	if o.IsNil() {
		return e()
	}
	return o.item
}

func (o Option[T]) UnwrapOrDefault(defaultValue T) T {
	if o.IsNil() {
		return defaultValue
	}
	return o.item
}

func (o Option[T]) UnwrapOrError(err error) (T, error) {
	if o.IsNil() {
		return *new(T), err
	}
	return o.item, nil
}

func (o Option[T]) UnwrapOrNil() T {
	return o.item
}

func (o Option[T]) UnwrapWithOk() (T, bool) {
	return o.UnwrapOrNil(), !o.IsNil()
}
