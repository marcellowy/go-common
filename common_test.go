package common

import (
	"fmt"
	"testing"
)

func TestIsNumber(t *testing.T) {

	if IsNumber("1234") == false {
		t.Error("error")
		return
	}

	if IsNumber("a123") == true {
		t.Error("error")
		return
	}
}

func TestRandStr(t *testing.T) {

	if len(RandStr(16)) != 16 {
		t.Error(fmt.Sprintf("rand str"))
		return
	}
}
