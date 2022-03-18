package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config log 配置
type Config struct {
	Filename             string
	RollingLogMaxSize    int
	RollingLogMaxBackups int
	RollingLogMaxAge     int
	Compress             bool
	LastLog              func(ctx context.Context) []zap.Field
}

var (
	// Logger 日志实例
	Logger        *zap.Logger
	defaultConfig = &Config{
		Filename:             getExecuteName(),
		RollingLogMaxSize:    1024 * 2,
		RollingLogMaxBackups: 10,
		RollingLogMaxAge:     7,
		Compress:             false,
		LastLog:              nil,
	}
)

func init() {
	initLog(defaultConfig)
}

func initLog(config *Config) {

	hook := lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.RollingLogMaxSize, // megabytes
		MaxBackups: config.RollingLogMaxBackups,
		MaxAge:     config.RollingLogMaxAge, // days
		Compress:   false,                   // disabled by default
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	multiWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	Logger = zap.New(zapcore.NewTee(zapcore.NewCore(consoleEncoder, multiWriter, highPriority))).
		WithOptions(zap.AddCaller()).
		WithOptions(zap.AddCallerSkip(1))
}

// getExecuteName 获取当前运行时的文件名
func getExecuteName() string {
	var path = os.Args[0]
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "\\", "/")
		if len(path) > 4 && path[len(path)-4:] == ".exe" {
			path = path[:len(path)-4]
		}
	}
	paths := strings.Split(path, "/")
	path = paths[len(paths)-1]
	return path + ".log"
}

// SetLogConfig 设置日志配置
func SetLogConfig(config *Config) {
	if config.Filename != "" {
		defaultConfig.Filename = config.Filename
	}
	if config.RollingLogMaxSize > 0 {
		defaultConfig.RollingLogMaxSize = config.RollingLogMaxSize
	}
	if config.RollingLogMaxAge > 0 {
		defaultConfig.RollingLogMaxAge = config.RollingLogMaxAge
	}
	if config.RollingLogMaxBackups > 0 {
		defaultConfig.RollingLogMaxBackups = config.RollingLogMaxBackups
	}
	if config.LastLog != nil {
		defaultConfig.LastLog = config.LastLog
	}
	initLog(defaultConfig)
}

// connectArgs 将要打印的参数连接起来
func connectArgs(args ...interface{}) string {
	var f string
	for range args {
		f += "%v "
	}
	return fmt.Sprintf(f, args...)
}

func lastLog(ctx context.Context, f func(ctx context.Context) []zap.Field) []zap.Field {
	if f != nil {
		return f(ctx)
	}
	return nil
}

func Debug(ctx context.Context, args ...interface{}) {
	Logger.Debug(connectArgs(args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Info(ctx context.Context, args ...interface{}) {
	Logger.Info(connectArgs(args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	Logger.Info(fmt.Sprintf(format, args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Warn(ctx context.Context, args ...interface{}) {
	Logger.Warn(connectArgs(args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	Logger.Warn(fmt.Sprintf(format, args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Error(ctx context.Context, args ...interface{}) {
	Logger.Error(connectArgs(args...), lastLog(ctx, defaultConfig.LastLog)...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	Logger.Error(fmt.Sprintf(format, args...), lastLog(ctx, defaultConfig.LastLog)...)
}
