// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import "runtime"

// All possible GOOS value
// const goosList = "android darwin dragonfly freebsd linux nacl \
// netbsd openbsd plan9 solaris windows "

// IsWindows check system windows
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

// IsDarwin  check system mac
func IsDarwin() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

// IsLinux  check system linux
func IsLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

// IsAndroid 判断是否linux系统
func IsAndroid() bool {
	if runtime.GOOS == "android" {
		return true
	}
	return false
}
