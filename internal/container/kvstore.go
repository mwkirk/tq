package container

import (
	"errors"
	"sync"
)

var ErrorNotFound = errors.New("not found")

// KVStore might be implemented by a database or external service that could return an error
type KVStore[T comparable, U any] interface {
	Exists(T) (bool, error)
	Get(T) (U, error)
	GetAndDelete(T) (U, error)
	Put(T, U) error
	Delete(T) error
	Update(T, func(v U) U) error
	Filter(func(U) bool) []U
}

type SimpleMapStore[T comparable, U any] struct {
	l sync.RWMutex
	m map[T]U
}

func NewSimpleMapStore[T comparable, U any]() *SimpleMapStore[T, U] {
	return &SimpleMapStore[T, U]{
		m: make(map[T]U),
	}
}

func (s *SimpleMapStore[T, U]) Exists(k T) (bool, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	_, ok := s.m[k]
	if !ok {
		return ok, ErrorNotFound
	}
	return ok, nil
}

func (s *SimpleMapStore[T, U]) Get(k T) (U, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	var v U
	v, ok := s.m[k]
	if !ok {
		return v, ErrorNotFound
	}
	return v, nil
}

func (s *SimpleMapStore[T, U]) GetAndDelete(k T) (U, error) {
	s.l.RLock()
	defer s.l.RUnlock()
	var v U
	v, ok := s.m[k]
	if !ok {
		return v, ErrorNotFound
	}
	delete(s.m, k)
	return v, nil
}

func (s *SimpleMapStore[T, U]) Put(k T, v U) error {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[k] = v
	return nil
}

func (s *SimpleMapStore[T, U]) Delete(k T) error {
	s.l.Lock()
	defer s.l.Unlock()
	delete(s.m, k)
	return nil
}

func (s *SimpleMapStore[T, U]) Update(k T, f func(v U) U) error {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[k] = f(s.m[k]) // todo: f() might fail, should handle this
	return nil
}

func (s *SimpleMapStore[T, U]) Filter(pred func(U) bool) []U {
	var filtered []U
	s.l.RLock()
	defer s.l.RUnlock()

	for _, v := range s.m {
		if ok := pred(v); ok {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
