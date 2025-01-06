package env

import (
	"github.com/vakhrushevk/chat-server-service/internal/config"
	"log/slog"
	"os"
	"strings"
)

var _ config.LogConfig = (*loggerConfig)(nil)

const (
	loggerLevel = "LOGGER_LEVEL"
	LEVEL_INFO  = slog.LevelInfo
	LEVEL_DEBUG = slog.LevelDebug
	LEVEL_WARN  = slog.LevelWarn
	LEVEL_ERROR = slog.LevelError
)

type loggerConfig struct {
	loggerLevel slog.Level
}

func (l *loggerConfig) LoggerLevel() slog.Level {
	return l.loggerLevel
}

// NewLoggerConfig - Создает и инициализирует конфигурацию для логгера
func NewLoggerConfig() (config.LogConfig, error) {
	levelStr := os.Getenv(loggerLevel)
	var level slog.Level
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "info":
		level = LEVEL_INFO
	case "debug":
		level = LEVEL_DEBUG
	case "warn":
		level = LEVEL_WARN
	case "error":
		level = LEVEL_ERROR
	default:
		level = LEVEL_INFO
	}

	return &loggerConfig{
		loggerLevel: level,
	}, nil
}
