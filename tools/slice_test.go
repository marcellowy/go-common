// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import (
	"testing"
)

func TestRemoveSameFromStringSlice(t *testing.T) {
	RemoveSameFromStringSlice([]string{"1", "1", "2", "2", "3"})
}

func TestSliceTrimSame(t *testing.T) {
	SliceTrimSame([]string{"1", "1", "2", "2", "3", "4", "4"})
	SliceTrimSame([]string{})
}

func TestSliceRemove(t *testing.T) {
	SliceRemove([]string{"1"}, []string{"1", "1", "2", "2", "3", "4", "4"})
}
