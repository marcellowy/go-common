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
	tmp := map[string]byte{} // 存放不重复主键
	for _, s := range slice {
		l := len(tmp)
		tmp[s] = 0
		if len(tmp) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, s)
		}
	}
	return result
}
