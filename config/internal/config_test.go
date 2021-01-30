package internal

import "testing"

func TestIsSupportedExt(t *testing.T) {
	if IsSupportedExt("yaml") == false {
		t.Error("supported err")
	}
}
