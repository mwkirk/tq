package container

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Front() T
	Back() T
	Length() int64
}

type SliceQueue[T any] struct {
	s []T
}

func NewSliceQueue[T any]() SliceQueue[T] {
	return SliceQueue[T]{
		s: make([]T, 0),
	}
}

func (sq *SliceQueue[T]) Enqueue(v T) {
	sq.s = append(sq.s, v)
}

func (sq *SliceQueue[T]) Dequeue() (T, bool) {
	if len(sq.s) == 0 {
		var zero T
		return zero, false
	}
	v := sq.s[0]
	sq.s = sq.s[1:]
	return v, true
}

func (sq *SliceQueue[T]) Front() (T, bool) {
	if len(sq.s) == 0 {
		var zero T
		return zero, false
	}
	return sq.s[0], true
}

func (sq *SliceQueue[T]) Back() (T, bool) {
	if len(sq.s) == 0 {
		var zero T
		return zero, false
	}
	return sq.s[len(sq.s)-1], true
}

func (sq *SliceQueue[T]) Length() int64 {
	return int64(len(sq.s))
}
