// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

// IsNumber check string is number?
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
