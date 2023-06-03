// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import "io"

// Close 快捷关闭
func Close(f io.Closer) {
	_ = f.Close()
}
