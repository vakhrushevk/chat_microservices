package app

import (
	"context"
	"github.com/vakhrushevk/chat-server-service/internal/logger"
	"log"

	"github.com/vakhrushevk/local-platform/closer"
	"github.com/vakhrushevk/local-platform/db"
	"github.com/vakhrushevk/local-platform/db/pg"
	"github.com/vakhrushevk/local-platform/db/transaction"

	"github.com/vakhrushevk/chat-server-service/internal/api/chat"
	"github.com/vakhrushevk/chat-server-service/internal/config"
	"github.com/vakhrushevk/chat-server-service/internal/config/env"
	"github.com/vakhrushevk/chat-server-service/internal/repository"
	"github.com/vakhrushevk/chat-server-service/internal/repository/postgres"
	"github.com/vakhrushevk/chat-server-service/internal/service"
	"github.com/vakhrushevk/chat-server-service/internal/service/chatservice"
)

type serviceProvider struct {
	pgConfig     config.PgConfig
	grpcConfig   config.GRPCConfig
	loggerConfig config.LogConfig

	dbClient  db.Client
	txManager db.TxManager

	chatRepository repository.ChatRepository
	chatService    service.ChatService

	chatImplementation *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) InitializeLogger() {
	logger.New(s.LogConfig().LoggerLevel())
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := pg.New(context.Background(), s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}
		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(client.Close)
		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) PGConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("Failed to get pg config: %v", err)
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) LogConfig() config.LogConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %v", err)
		}
		s.loggerConfig = cfg
	}
	return s.loggerConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		repo := postgres.NewChatRepository(s.DBClient(ctx))
		s.chatRepository = repo
	}
	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		serv := chatservice.New(s.ChatRepository(ctx), s.TxManager(ctx))
		s.chatService = serv
	}
	return s.chatService
}

func (s *serviceProvider) ChatImplementation(ctx context.Context) *chat.Implementation {
	if s.chatImplementation == nil {
		s.chatImplementation = chat.NewChatImplementation(s.ChatService(ctx))
	}
	return s.chatImplementation
}
