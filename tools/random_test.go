package tools

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	var length uint = 5
	a := RandomString(length, RandomMajuscule)
	fmt.Println(a)
	if len(a) != int(length) {
		t.Errorf("length err")
		return
	}
}

func TestRandomString2(t *testing.T) {
	var length uint = 15
	a := RandomString(length, RandomAll)
	fmt.Println(a)
	if len(a) != int(length) {
		t.Errorf("length err")
		return
	}
}

func TestRandomString3(t *testing.T) {
	var length uint = 17
	a := RandomString(length, RandomLowercase)
	fmt.Println(a)
	if len(a) != int(length) {
		t.Errorf("length err")
		return
	}
}

func TestRandomString4(t *testing.T) {
	var length uint = 15
	a := RandomString(length, RandomSymbol)
	fmt.Println(a)
	if len(a) != int(length) {
		t.Errorf("length err")
		return
	}
}
