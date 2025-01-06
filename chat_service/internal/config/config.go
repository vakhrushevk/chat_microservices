package config

import (
	"github.com/joho/godotenv"
	"log/slog"
)

// PgConfig - Конфиг для подключения к базе данных
type PgConfig interface {
	DSN() string
}

// GRPCConfig - Конфиг для подключения к grpc
type GRPCConfig interface {
	Address() string
}

// LogConfig - Конфиг для логгера
// LoggerLevel - уровень логгирования
type LogConfig interface {
	LoggerLevel() slog.Level
}

// Load - Load config from path
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
