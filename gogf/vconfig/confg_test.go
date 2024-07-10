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

func TestNewConfigFromFile(t *testing.T) {
	config, err := NewConfigFromFile("new_config.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	val := config.Get("test_new")
	if val.Int() != 123 {
		t.Error("get error")
		return
	}
}
