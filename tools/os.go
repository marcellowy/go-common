// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import "runtime"

// All possible GOOS value
// const goosList = "android darwin dragonfly freebsd linux nacl \
// netbsd openbsd plan9 solaris windows "

// IsWindows 判断是不是windows
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

// IsDarwin 判断是不是Mac
func IsDarwin() bool {
	if runtime.GOOS == "darwin" {
		return true
	}
	return false
}

// IsLinux 判断是否linux系统
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
