package constraint

type Order[T any] interface {
	Less(b T) bool
}

type BasicOrder interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~byte
}

func Less[T BasicOrder](a T, b T) bool {
	return a < b
}
