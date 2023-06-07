// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"os"
	"path/filepath"
	"strings"
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

// GetCurrentDirectory return current work directory
func GetCurrentDirectory() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}

	if IsWindows() {
		// only windows replace
		return strings.Replace(dir, "\\", "/", -1) // use / replace \\
	}
	return dir
}
