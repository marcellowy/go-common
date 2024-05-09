// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"unsafe"
)

// ReverseByte reverse []byte
func ReverseByte(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseString reverse string
func ReverseString(s string) string {
	s1 := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
	return string(s1)
}

// StringToBytes converts string to a byte slice.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesToString converts byte to a string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// SubString returns a substring of 's' starting from 'pos' and optionally ending at 'to'.
// If 'pos' is out of bounds, it returns the original string.
func SubString(s string, pos uint, to ...uint) string {
	if len(s) == 0 {
		return ""
	}
	if pos > uint(len(s)-1) {
		return s
	}
	var endNum = uint(len(s))
	if len(to) == 1 && to[0] < endNum {
		endNum = to[0]
	}
	if endNum < pos {
		return ""
	}
	return s[pos:endNum]
}
