// Package config
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package config

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

// Get 获取配置
func Get(s string, def ...interface{}) *gvar.Var {
	vr, err := g.Config().Get(context.Background(), s, def...)
	if err != nil {
		return &gvar.Var{}
	}
	return vr
}
