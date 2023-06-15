// Package log
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package log

import (
	"context"
	"fmt"
	"github.com/marcellowy/go-common/config"
	"github.com/marcellowy/go-common/consts"
	"github.com/marcellowy/go-common/server/ginctx"
	"github.com/marcellowy/go-common/tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var (
	defaultTraceName = "trace_id"
	logger           *zap.Logger
)

func init() {

	hook := lumberjack.Logger{
		Filename: getExecuteLogFilename(),
	}
	if v := config.GetString("logger.filename"); v != "" {
		hook.Filename = v
	}
	if v := config.GetInt("logger.maxSize"); v > 0 {
		hook.MaxSize = v
	}

	if v := config.GetInt("logger.maxAge"); v > 0 {
		hook.MaxAge = v
	}

	if v := config.GetInt("logger.maxBackups"); v > 0 {
		hook.MaxBackups = v
	}

	hook.LocalTime = config.GetBool("logger.localTime")
	hook.Compress = config.GetBool("logger.compress")

	// zap
	level, err := zapcore.ParseLevel(config.GetString("logger.level"))
	if err != nil {

	}
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})
	var writeSyncer = []zapcore.WriteSyncer{
		zapcore.AddSync(&hook),
	}
	if config.GetBool("logger.stdout") {
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	multiWriter := zapcore.NewMultiWriteSyncer(writeSyncer...)
	logger = zap.New(
		zapcore.NewTee(zapcore.NewCore(consoleEncoder, multiWriter, highPriority)),
	).
		WithOptions(zap.AddCaller()).
		WithOptions(zap.AddCallerSkip(1))
}

func ReadTraceId(ctx context.Context) string {
	return ginctx.ReadTrace(ctx)
}

// getExecuteLogFilename 获取当前运行时的文件名
func getExecuteLogFilename() string {
	var path = os.Args[0]
	if tools.IsWindows() {
		path = strings.ReplaceAll(path, "\\", "/")
		if len(path) > 4 && path[len(path)-4:] == ".exe" {
			path = path[:len(path)-4]
		}
	}
	paths := strings.Split(path, "/")
	path = paths[len(paths)-1]
	return path + ".log"
}

func toString(msg ...any) string {
	var token = ""
	for range msg {
		token += "%v "
	}
	return fmt.Sprintf(token, msg...)
}

func AddTraceId(ctx context.Context) context.Context {
	return context.WithValue(ctx, consts.DefaultTraceKey, "")
}

func Debug(ctx context.Context, msg ...any) {
	logger.Debug(toString(msg...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}

func Info(ctx context.Context, msg ...any) {
	logger.Info(toString(msg...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Warn(ctx context.Context, msg ...any) {
	logger.Warn(toString(msg...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Error(ctx context.Context, msg ...any) {
	logger.Error(toString(msg...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Fatal(ctx context.Context, msg ...any) {
	logger.Fatal(toString(msg...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}

func Debugf(ctx context.Context, msg string, v ...any) {
	logger.Debug(fmt.Sprintf(msg, v...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Infof(ctx context.Context, msg string, v ...any) {
	logger.Info(fmt.Sprintf(msg, v...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Warnf(ctx context.Context, msg string, v ...any) {
	logger.Warn(fmt.Sprintf(msg, v...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Errorf(ctx context.Context, msg string, v ...any) {
	logger.Error(fmt.Sprintf(msg, v...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
func Fatalf(ctx context.Context, msg string, v ...any) {
	logger.Fatal(fmt.Sprintf(msg, v...), zap.String(defaultTraceName, ReadTraceId(ctx)))
}
