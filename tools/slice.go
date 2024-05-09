// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import "fmt"

// SliceTrimSame remove same value
func SliceTrimSame[T comparable](s []T) []T {
	var (
		tmp = make(map[T]byte, len(s))
		cc  = make([]T, len(s))
	)
	for _, el := range s {
		l := len(tmp)
		tmp[el] = 0
		if len(tmp) > l {
			cc[len(tmp)-1] = el
		}
	}
	cc = cc[:len(tmp)]
	return cc
}

// SliceRemove remove value
func SliceRemove[T comparable](remove []T, s []T) []T {
	var (
		tmp = make(map[T]byte, len(remove))
		cc  = make([]T, len(s))
	)
	for _, v := range remove {
		tmp[v] = 0
	}
	n := 0
	for _, v := range s {
		fmt.Println(v)
		if _, ok := tmp[v]; !ok {
			cc[n] = v
			n++
		}
	}
	return cc[:n]
}

// RemoveSameFromStringSlice remove []string same element
// Deprecated: use SliceTrimSame replace
func RemoveSameFromStringSlice(slice []string) []string {

	var result []string
	tmp := map[string]byte{}
	for _, s := range slice {
		l := len(tmp)
		tmp[s] = 0
		if len(tmp) != l {
			result = append(result, s)
		}
	}
	return result
}

// InSlice check value in []T
func InSlice[T comparable](need T, arr []T) bool {
	for _, v := range arr {
		if v == need {
			return true
		}
	}
	return false
}
