package container

import (
	"errors"
)

var ErrorNotFound = errors.New("no found")

type Store[T comparable, U any] interface {
	Get(T) (U, error)
	Add(T, U) error
	Remove(T) error
}

// TODO: May need to protect this with a mutex if it turns out it will be accessed concurrently
type SimpleMapStore[T comparable, U any] struct {
	m map[T]U
}

func NewSimpleMapStore[T comparable, U any]() SimpleMapStore[T, U] {
	return SimpleMapStore[T, U]{
		m: make(map[T]U),
	}
}

func (s *SimpleMapStore[T, U]) Get(k T) (U, error) {
	var v U
	v, ok := s.m[k]
	if !ok {
		return v, ErrorNotFound
	}
	return v, nil
}

func (s *SimpleMapStore[T, U]) Add(k T, v U) error {
	s.m[k] = v
	return nil
}

func (s *SimpleMapStore[T, U]) Remove(k T) error {
	delete(s.m, k)
	return nil
}
