// Package tools
package tools

import (
	"bytes"
	"testing"
)

func TestReverseByte(t *testing.T) {
	var a = []byte("ab")
	if bytes.Compare(ReverseByte(a), []byte("ba")) != 0 {
		t.Error("err")
	}
}
