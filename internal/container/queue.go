package container

import (
	"errors"
	"sync"
)

var ErrorQueueEmpty = errors.New("queue empty")

type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Front() (T, error)
	Back() (T, error)
	Length() (int64, error)
	Filter(func(T) bool) []T
}

type SliceQueue[T any] struct {
	l sync.RWMutex
	s []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		s: make([]T, 0),
	}
}

func (sq *SliceQueue[T]) Enqueue(v T) error {
	sq.l.Lock()
	defer sq.l.Unlock()
	sq.s = append(sq.s, v)
	return nil
}

func (sq *SliceQueue[T]) Dequeue() (T, error) {
	sq.l.Lock()
	defer sq.l.Unlock()
	if len(sq.s) == 0 {
		var zero T
		return zero, ErrorQueueEmpty
	}
	v := sq.s[0]
	sq.s = sq.s[1:]
	return v, nil
}

func (sq *SliceQueue[T]) Front() (T, error) {
	sq.l.RLock()
	defer sq.l.RUnlock()
	if len(sq.s) == 0 {
		var zero T
		return zero, ErrorQueueEmpty
	}
	return sq.s[0], nil
}

func (sq *SliceQueue[T]) Back() (T, error) {
	sq.l.RLock()
	defer sq.l.RUnlock()
	if len(sq.s) == 0 {
		var zero T
		return zero, ErrorQueueEmpty
	}
	return sq.s[len(sq.s)-1], nil
}

func (sq *SliceQueue[T]) Length() (int64, error) {
	sq.l.RLock()
	defer sq.l.RUnlock()
	return int64(len(sq.s)), nil
}

func (sq *SliceQueue[T]) Filter(pred func(T) bool) []T {
	var filtered []T
	sq.l.RLock()
	defer sq.l.RUnlock()

	for _, v := range sq.s {
		if ok := pred(v); ok {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
