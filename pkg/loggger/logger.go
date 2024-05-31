package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Level      string
	Directory  string
	FileName   string
	ToFile     bool
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

var Logger *zap.Logger

func NewLogger(config Config) error {
	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	atomicLevel := zap.NewAtomicLevelAt(level)

	var writer zapcore.WriteSyncer
	if config.ToFile {
		logPath := filepath.Join(config.Directory, config.FileName+".log")
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		})
	} else {
		writer = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, writer, atomicLevel)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))

	zap.ReplaceGlobals(Logger) // 设置全局日志实例

	return nil
}
