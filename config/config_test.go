// Package config
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package config

import (
	"testing"
)

func TestGetString(t *testing.T) {
	if GetString("server.listen") == "" {
		t.Error("server.listen err")
		return
	}
}

func TestGetStringSlice(t *testing.T) {
	if len(GetStringSlice("server.arr")) == 0 {
		t.Error("server.arr err")
		return
	}
}
