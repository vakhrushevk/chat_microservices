package main

import (
	"context"
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vakhrushevk/chat-server-service/internal/logger"
	"github.com/vakhrushevk/chat-server-service/internal/metric"
	"log"
	"net/http"

	"github.com/vakhrushevk/chat-server-service/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {

	flag.Parse()
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}

	metric.Init(context.Background())

	go runPrometheusServer()

	if err := a.Run(); err != nil {
		logger.Fatal("failed to run app: ", logger.ErrAttr(err))
	}
}

// TODO: Export to App file
func runPrometheusServer() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	log.Println("HELLO WORLD!")
	logger.Info("Prometheus server is running on localhost:2112")
	log.Fatal(http.ListenAndServe(":2112", mux))
}
