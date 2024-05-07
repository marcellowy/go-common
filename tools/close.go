// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import "io"

// Close is a Go function that takes an io.Closer as a parameter and closes it.
//
// The parameter `f` is the io.Closer that needs to be closed.
// This function does not return anything.
func Close(f io.Closer) {
	_ = f.Close()
}
