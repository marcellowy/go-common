package tools

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	a := RandomStr(16, RandomAll)
	if len(a) != 16 {
		t.Errorf("length err")
		return
	}
}
