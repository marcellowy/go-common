// Package tools
// Copyright 2016-2023 volibear<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"sync"
	"sync/atomic"
)

// TSMap thread safe map
type TSMap[K, T any] struct {
	count atomic.Int64
	value sync.Map
}

func NewTSMap[K, T any]() *TSMap[K, T] {
	return &TSMap[K, T]{}
}

func (m *TSMap[K, T]) Store(key K, value T) {
	if _, ok := m.value.Swap(key, value); !ok {
		m.count.Add(1)
	}
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *TSMap[K, T]) Load(key K) (T, bool) {
	var (
		value     any
		realValue T
		ok        bool
	)
	if value, ok = m.value.Load(key); ok {
		if realValue, ok = value.(T); ok {
			return realValue, ok
		}
	}
	return realValue, ok
}

func (m *TSMap[K, T]) Len() int64 {
	return m.count.Load()
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *TSMap[K, T]) Range(f func(key K, value T) bool) {
	m.value.Range(func(key, value any) bool {
		var (
			realKey   K
			realValue T
			ok        bool
		)

		if realKey, ok = key.(K); !ok {
			return false
		}

		if realValue, ok = value.(T); !ok {
			return false
		}

		return f(realKey, realValue)
	})
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *TSMap[K, T]) LoadAndDelete(key K) (T, bool) {
	var realValue T
	if v, ok := m.value.LoadAndDelete(key); ok {
		m.count.Add(-1)
		if realValue, ok = v.(T); ok {
			return realValue, ok
		}
		return realValue, false
	}
	return realValue, false
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *TSMap[K, T]) LoadOrStore(key K, value T) (T, bool) {
	var (
		v         any
		realValue T
		ok        bool
	)
	if v, ok = m.value.LoadOrStore(key, value); !ok {
		m.count.Add(1)
	}
	if realValue, ok = v.(T); ok {
		return realValue, ok
	}
	return realValue, false
}
