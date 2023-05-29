// Package verr
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package verr

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// New 对 gerror 的封装
func New(code int64, message string) error {
	return gerror.NewCode(gcode.New(int(code), message, ""))
}
