// Package vlog
package vlog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func Info(ctx context.Context, v ...interface{}) {
	g.Log().Skip(1).Line(true).Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(1).Line(true).Infof(ctx, format, v...)
}

func Error(ctx context.Context, v ...interface{}) {
	g.Log().Skip(1).Line(true).Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(1).Line(true).Errorf(ctx, format, v...)
}

func Debug(ctx context.Context, v ...interface{}) {
	g.Log().Skip(1).Line(true).Debug(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(1).Line(true).Debugf(ctx, format, v...)
}

func Warning(ctx context.Context, v ...interface{}) {
	g.Log().Skip(1).Line(true).Warning(ctx, v...)
}

func Warningf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Skip(1).Line(true).Warningf(ctx, format, v...)
}
