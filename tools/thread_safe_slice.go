// Package tools
// Copyright 2016-2024 volibear<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"fmt"
	"golang.org/x/exp/slices"
	"sync"
)

// TSSlice thread safe slice
type TSSlice[T any] struct {
	value        []T
	defaultValue T
	locker       sync.RWMutex
}

// NewTSSlice new thread safe slice object
func NewTSSlice[T any](def T) *TSSlice[T] {
	return &TSSlice[T]{
		value:        make([]T, 0),
		defaultValue: def,
		locker:       sync.RWMutex{},
	}
}

// Remove remove element by index
func (s *TSSlice[T]) Remove(i, j int) error {
	s.locker.Lock()
	s.value = slices.Delete(s.value, i, j)
	s.locker.Unlock()
	return nil
}

// Insert element
func (s *TSSlice[T]) Insert(i int, v ...T) {
	s.locker.Lock()
	s.value = slices.Insert(s.value, i, v...)
	s.locker.Unlock()
}

// Append append
func (s *TSSlice[T]) Append(v ...T) {
	s.locker.Lock()
	s.value = append(s.value, v...)
	s.locker.Unlock()
}

// Reverse reverses the elements of the slice in place.
func (s *TSSlice[T]) Reverse() {
	s.locker.Lock()
	slices.Reverse(s.value)
	s.locker.Unlock()
}

// Index get value
func (s *TSSlice[T]) Index(index int) (T, error) {
	s.locker.Lock()
	if len(s.value)-1 < index || index < 0 {
		s.locker.Unlock()
		return s.defaultValue, fmt.Errorf("index out range of slice length")
	}
	var value = s.value[index]
	s.locker.Unlock()
	return value, nil
}

func (s *TSSlice[T]) Clone() *TSSlice[T] {
	s.locker.Lock()
	newValue := slices.Clone(s.value)
	s.locker.Unlock()
	return &TSSlice[T]{
		value:        newValue,
		defaultValue: s.defaultValue,
		locker:       sync.RWMutex{},
	}
}

// Get slice
func (s *TSSlice[T]) Get() []T {
	return s.value
}

// Len get slice length
func (s *TSSlice[T]) Len() int {
	return len(s.value)
}
