// Package vconfig
package vconfig

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

// Get 从默认配置文件中获取配置
func Get(s string, def ...interface{}) *gvar.Var {
	vr, err := g.Config().Get(context.Background(), s, def...)
	if err != nil {
		return &gvar.Var{}
	}
	return vr
}

type CustomFileConfig struct {
	cfg *gcfg.Config
}

func (cc *CustomFileConfig) Get(s string, def ...interface{}) *gvar.Var {
	vr, err := cc.cfg.Get(context.Background(), s, def...)
	if err != nil {
		return &gvar.Var{}
	}
	return vr
}

// NewConfigFromFile 从不同的文件中获取配置对象
func NewConfigFromFile(filename string) (*CustomFileConfig, error) {
	adapterFile, err := gcfg.NewAdapterFile(filename)
	if err != nil {
		return nil, err
	}
	cfg := gcfg.Config{}
	cfg.SetAdapter(adapterFile)
	return &CustomFileConfig{cfg: &cfg}, nil
}
