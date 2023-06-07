// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"testing"
)

func TestMd5(t *testing.T) {
	if Md5("1") != "c4ca4238a0b923820dcc509a6f75849b" {
		t.Error("md5 err")
		return
	}
}
func TestMd5File(t *testing.T) {

	if s, err := Md5File("md5_test_file.txt"); err != nil {
		t.Error(err)
	} else if s != "f487785942516134c63f3da573a1109f" {
		t.Error("md5 file val err")
	}
}
