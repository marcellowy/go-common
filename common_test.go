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

func TestRemoveSameFromStringSlice(t *testing.T) {
	var s = []string{"1", "2", "1"}
	var r = RemoveSameFromStringSlice(s)
	if len(r) != 2 {
		t.Error("error")
	}
	fmt.Println(r)

}
