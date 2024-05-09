package vconfig

import (
	"testing"
)

func TestGet(t *testing.T) {
	if Get("a").String() != "1" {
		t.Error("get error")
		return
	}
}
