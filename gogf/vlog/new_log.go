// Package vlog
package vlog

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// New logger
func New(key string) *glog.Logger {
	return g.Log(key).Line(true)
}
