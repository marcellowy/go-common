package db

import (
	"testing"
)

func TestNewDbConfig(t *testing.T) {

	config := NewConfig("", 0, "", "", "", "")
	if _, ok := interface{}(config).(*Config); !ok {
		t.Error("config not instanceof DBConfig")
	}
}
