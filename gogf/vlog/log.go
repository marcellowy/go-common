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
	logger.Skip(1).Line(true).Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line(true).Infof(ctx, format, v...)
}

func Error(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line(true).Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line(true).Errorf(ctx, format, v...)
}

func Debug(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line(true).Debug(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line(true).Debugf(ctx, format, v...)
}

func Warning(ctx context.Context, v ...interface{}) {
	logger.Skip(1).Line(true).Warning(ctx, v...)
}

func Warningf(ctx context.Context, format string, v ...interface{}) {
	logger.Skip(1).Line(true).Warningf(ctx, format, v...)
}
