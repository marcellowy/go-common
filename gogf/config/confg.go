// Package config
package config

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

// Deprecated: use vconfig.Get instead.
func Get(s string, def ...interface{}) *gvar.Var {
	vr, err := g.Config().Get(context.Background(), s, def...)
	if err != nil {
		return &gvar.Var{}
	}
	return vr
}
