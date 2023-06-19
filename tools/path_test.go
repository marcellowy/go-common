// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import (
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	dir := GetCurrentDirectory()
	if dir == "" {
		t.Error("GetCurrentDirectory err")
		return
	}
}

func TestPathExists(t *testing.T) {
	ok, err := PathExists(".")
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("path not exists")
		return
	}

	ok, err = PathExists("/aaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Error(err)
		return
	}

	if ok {
		// directory not exists and ok is false it's true
		t.Error(err)
		return
	}

}
