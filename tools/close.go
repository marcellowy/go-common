// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import "io"

// Close fast ptr close
func Close(f io.Closer) {
	_ = f.Close()
}
