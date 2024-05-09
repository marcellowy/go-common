// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import "unsafe"

// ReverseByte reverse []byte
func ReverseByte(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// StringToBytes converts string to a byte slice.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesToString converts byte to a string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
