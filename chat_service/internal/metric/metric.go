package metric

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "my_space"
	appName   = "my_app"
)

type Metrics struct {
	requestCounter prometheus.Counter
}

var metrics *Metrics

func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "grpc",
			Name:      appName + "_request_total",
			Help:      "The total number of requests",
		}),
	}

	return nil
}

func IncRequestCounter() {
	metrics.requestCounter.Inc()
}
