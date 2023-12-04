// Package db
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package db

import (
	"context"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vlog"
	gLogger "gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	SlowThreshold time.Duration // 慢查询时间,超过 SlowThreshold 就算慢查询
	Name          []string
}

func (gl GormLogger) LogMode(level gLogger.LogLevel) gLogger.Interface {
	return gl
}

func (gl GormLogger) Info(ctx context.Context, s string, v ...interface{}) {
	vlog.Infof(ctx, s, v...)
}

func (gl GormLogger) Warn(ctx context.Context, s string, v ...interface{}) {
	vlog.Warningf(ctx, s, v...)
}

func (gl GormLogger) Error(ctx context.Context, s string, v ...interface{}) {
	vlog.Errorf(ctx, s, v...)
}

func (gl GormLogger) Trace(ctx context.Context, begin time.Time,
	fc func() (sql string, rowsAffected int64), err error) {

	sql, _ := fc()
	elapsed := time.Since(begin)

	if elapsed > gl.SlowThreshold {
		fmt.Println(elapsed.String(), gl.SlowThreshold.String())
	}

	var title = "SQL"
	if elapsed > gl.SlowThreshold && gl.SlowThreshold > 0 {
		// 慢查询
		title = "SLOW SQL"
	}
	vlog.Infof(ctx, "%s: %s [%.3fms]", title, sql, float64(elapsed.Nanoseconds())/1e6)
}

type GormLogOptions func(*GormLogger)

// GormLogWithSlowThreshold 慢查询
func GormLogWithSlowThreshold(st time.Duration) GormLogOptions {
	return func(gormLogger *GormLogger) {
		gormLogger.SlowThreshold = st
	}
}

// GormLogWithName 日志名
func GormLogWithName(name ...string) GormLogOptions {
	return func(gormLogger *GormLogger) {
		gormLogger.Name = name
	}
}

// NewGormLog 定义数据库日志打印输出
func NewGormLog(option ...GormLogOptions) gLogger.Interface {

	gl := &GormLogger{
		SlowThreshold: 200 * time.Millisecond, // 和官方默认值保持一致
	}

	for _, opt := range option {
		opt(gl)
	}

	return gl
}
