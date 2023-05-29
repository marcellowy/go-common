// Package verr
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package verr

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(-1, "test")
	if _, ok := err.(*gerror.Error); !ok {
		t.Errorf("object err")
		return
	}
}
