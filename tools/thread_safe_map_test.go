package tools

import (
	"fmt"
	"testing"
)

func TestTSMap_NewTSMap(t *testing.T) {
	NewTSMap[int, string]()
}

func TestTSMap_Store(t *testing.T) {
	a := NewTSMap[int, string]()
	a.Store(0, "1")
}

func TestTSMap_Load(t *testing.T) {
	a := NewTSMap[int, string]()
	a.Store(0, "1")
	s, ok := a.Load(0)
	if !ok {
		t.Error("value not found")
		return
	}
	// output: 1
	fmt.Println(s)
	if s != "1" {
		t.Error("value error")
		return
	}
}

func TestTSMap_LoadAndDelete(t *testing.T) {
	a := NewTSMap[int, string]()
	a.Store(0, "1")
	s, ok := a.LoadAndDelete(0)
	if !ok {
		t.Error("value not found")
		return
	}
	if s != "1" {
		t.Error("value error")
		return
	}
	if a.Len() > 0 {
		t.Error("len error")
		return
	}
}
