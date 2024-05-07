// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

// IsNumber checks if the given string is a number.
//
// Parameters:
// - s: the string to be checked.
//
// Returns:
// - bool: true if the string is a number, false otherwise.
func IsNumber(s string) bool {

	if s == "" {
		return false
	}

	var a = []byte(s)
	for _, v := range a {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
