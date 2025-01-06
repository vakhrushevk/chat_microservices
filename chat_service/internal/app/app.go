package app

import (
	"errors"
	"github.com/vakhrushevk/chat-server-service/internal/logger"
	"github.com/vakhrushevk/chat-server-service/internal/metric/interceptor"
	"log/slog"
	"net"

	"github.com/vakhrushevk/local-platform/closer"

	"github.com/vakhrushevk/chat-server-service/internal/config"
	"github.com/vakhrushevk/chat-server-service/pkg/chat_v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App - Структура ‘App’ представляет собой приложение.
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// New - Функция ‘New’ создает новый экземпляр приложения.
func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, errors.New("failed to start App, " + err.Error())
	}

	return a, nil
}

// Run - ‘Run’ метод запускает сервер gRPC.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initGRPCService,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	a.serviceProvider.InitializeLogger()
	return nil
}

func (a *App) initGRPCService(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.MetricInterceptor),
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(a.grpcServer)
	chat_v1.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImplementation(ctx))
	return nil
}

func (a *App) runGRPCServer() error {

	logger.Info("GRPC server is running on ",
		slog.Any("addres:", a.serviceProvider.GRPCConfig().Address()))

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}
	return nil
}
