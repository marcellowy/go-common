package vstrings

import (
	"testing"
)

func TestReplaceLineBreaks(t *testing.T) {
	var result = "abcd"
	var s = `a
b
c
d`
	s = ReplaceLineBreaks(s)
	if s != result {
		t.Error("result not match")
		return
	}
}
