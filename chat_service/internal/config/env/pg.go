package env

import (
	"errors"
	"os"

	"github.com/vakhrushevk/chat-server-service/internal/config"
)

var _ config.PgConfig = (*pgConfig)(nil)

const (
	dnsEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

// NewPGConfig создает и инициализирует конфигурацию для подключения к бд
func NewPGConfig() (config.PgConfig, error) {
	dsn := os.Getenv(dnsEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}
	return &pgConfig{
		dsn: dsn,
	}, nil
}

// DSN - Возвращает строку подключения к бд
func (p *pgConfig) DSN() string {
	return p.dsn
}
