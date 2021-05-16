package log

import (
	"context"
	"fmt"

	"gitee.com/marcellos/wyi-common/access"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func Init(filename string) {

	hook := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1024 * 2, // megabytes
		MaxBackups: 10,
		MaxAge:     7,     // days
		Compress:   false, // disabled by default
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	fileWriter := zapcore.AddSync(&hook)

	Logger = zap.New(zapcore.NewTee(zapcore.NewCore(consoleEncoder, fileWriter, highPriority))).
		WithOptions(zap.AddCaller()).
		WithOptions(zap.AddCallerSkip(1))
}

// traceId 打日志用的trace id
func traceId(ctx context.Context) zap.Field {
	var header *access.Header
	if c, ok := ctx.(*gin.Context); ok {
		header = c.MustGet("header").(*access.Header)
	} else {
		header = &access.Header{}
	}
	return zap.Field{
		Key:    "trace_id",
		Type:   zapcore.StringType,
		String: header.TraceID,
	}
}

// connectArgs 将要打印的参数连接起来
func connectArgs(args ...interface{}) string {
	var f string
	for range args {
		f += "%v "
	}
	return fmt.Sprintf(f, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	Logger.Error(connectArgs(args...), traceId(ctx))
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	Logger.Error(fmt.Sprintf(format, args...), traceId(ctx))
}

func Info(ctx context.Context, args ...interface{}) {
	Logger.Info(connectArgs(args...), traceId(ctx))
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	Logger.Info(fmt.Sprintf(format, args...), traceId(ctx))
}
