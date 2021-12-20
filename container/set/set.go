package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewFromSlice[T comparable](data []T) *Set[T] {
	s := New[T]()
	for _, d := range data {
		s.Put(d)
	}
	return s
}

func New[T comparable]() *Set[T] {
	s := Set[T]{
		set: make(map[T]struct{}),
	}
	return &s
}

func (s *Set[T]) Put(data T) {
	s.set[data] = struct{}{}
}

func (s *Set[T]) Remove(data T) {
	delete(s.set, data)
}

func (s *Set[T]) Exist(data T) bool {
	_, ok := s.set[data]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.set)
}

func (s *Set[T]) ToSlice() []T {
	res := make([]T, 0, len(s.set))
	for k := range s.set {
		res = append(res, k)
	}
	return res
}
