package env

import (
	"errors"
	"net"
	"os"

	"github.com/vakhrushevk/chat-server-service/internal/config"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

// Address - Объединяет значение хоста и порта из конфигурации
func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}

// NewGRPCConfig - Создает и возвращает конфигурацию для gRPC, извлекая значения хоста и порта
func NewGRPCConfig() (config.GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}
