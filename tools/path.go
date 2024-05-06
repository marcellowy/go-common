// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"os"
	"path/filepath"
)

// PathExists check path
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetCurrentDirectory return current execute file directory
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

func RemoveLastSeparator(path string) string {
	var length = len(path) - 1
	if length == -1 {
		return ""
	}
	if path[length] == '\\' {
		return path[:length]
	}
	if path[length] == '/' {
		return path[:length]
	}
	return path
}
