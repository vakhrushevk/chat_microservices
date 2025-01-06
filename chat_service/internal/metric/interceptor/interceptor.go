package interceptor

import (
	"context"
	"github.com/vakhrushevk/chat-server-service/internal/metric"
	"google.golang.org/grpc"
)

func MetricInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	metric.IncRequestCounter()

	return handler(ctx, req)
}
