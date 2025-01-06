package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/vakhrushevk/chat-server-service/internal/logger/slogpretty"
	"log/slog"
	"os"
)

var globalSlogger *slog.Logger
var globalSloggerPretty *slog.Logger

func New(level slog.Level) {
	if globalSlogger != nil && globalSloggerPretty != nil {
		return
	}
	switch level {
	case slog.LevelDebug:
		globalSlogger = setupJsonSlogWithFile(level)
		globalSloggerPretty = setupPrettySlog(level)
	default:
		globalSlogger = setupJsonSlogWithFile(level)
		globalSloggerPretty = setupPrettySlog(level)
	}
}

func setupPrettySlog(level slog.Level) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}

func setupJsonSlogWithFile(level slog.Level) *slog.Logger {
	opts := slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	}
	handler := slog.NewJSONHandler(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}, &opts)
	return slog.New(handler)
}

func Debug(msg string, args ...interface{}) {
	globalSlogger.Debug(msg, args...)
	globalSloggerPretty.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	globalSlogger.Info(msg, args...)
	globalSloggerPretty.Info(msg, args...)
}

func Error(msg string, args ...interface{}) {
	globalSlogger.Error(msg, args...)
	globalSloggerPretty.Error(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	globalSlogger.Warn(msg, args...)
	globalSloggerPretty.Warn(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	globalSlogger.Error(msg, args...)
	globalSloggerPretty.Error(msg, args...)
	os.Exit(1)
}

func ErrAttr(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
