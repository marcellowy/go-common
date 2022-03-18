package log

import (
	"context"
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
	Level                zapcore.Level
	StdOut               bool
	FileOut              bool
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
		Level:                zap.DebugLevel,
		StdOut:               false,
		FileOut:              true,
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
		return lvl >= defaultConfig.Level
	})

	var writeSyncer []zapcore.WriteSyncer

	if defaultConfig.StdOut {
		// 同时打印到标准输出
		writeSyncer = append(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	if defaultConfig.FileOut {
		writeSyncer = append(writeSyncer, zapcore.AddSync(&hook))
	}

	multiWriter := zapcore.NewMultiWriteSyncer(writeSyncer...)
	Logger = zap.New(
		zapcore.NewTee(zapcore.NewCore(consoleEncoder, multiWriter, highPriority)),
	).
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

func lastLog(ctx context.Context, f func(ctx context.Context) []zap.Field) []zap.Field {
	if f != nil {
		return f(ctx)
	}
	return nil
}

// Debug Debug日志
func Debug(ctx context.Context, msg string) {
	Logger.Debug(msg, lastLog(ctx, defaultConfig.LastLog)...)
}

// Info 信息日志
func Info(ctx context.Context, msg string) {
	Logger.Info(msg, lastLog(ctx, defaultConfig.LastLog)...)
}

// Warn 警告日志
func Warn(ctx context.Context, msg string) {
	Logger.Warn(msg, lastLog(ctx, defaultConfig.LastLog)...)
}

// Error 错误日志
func Error(ctx context.Context, msg string) {
	Logger.Error(msg, lastLog(ctx, defaultConfig.LastLog)...)
}
