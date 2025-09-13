// Package vlog
package vlog

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

var logger *glog.Logger

func init() {
	logger = g.Log()
}

func SetLogger(log *glog.Logger) {
	logger = log
}

func Info(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line().Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line().Infof(ctx, format, v...)
}

func Error(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line().Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line().Errorf(ctx, format, v...)
}

func Debug(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line().Debug(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line().Debugf(ctx, format, v...)
}

func Warning(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line().Warning(ctx, v...)
}

func Warningf(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line().Warningf(ctx, format, v...)
}
