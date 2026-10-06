package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type MetricsConfig struct {
	Host string `envconfig:"METRICS_HOST"`
	Port string `envconfig:"METRICS_PORT"`
	Path string `envconfig:"METRICS_PATH"`
}

// HTTP
var (
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_http_requests_total",
			Help: "Total HTTP requests processed, labeled by method, path and status",
		},
		[]string{"method", "path", "status"},
	)

	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "app_http_request_duration_seconds",
			Help:    "HTTP request latency distributions",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)
)

func safeRegister(c prometheus.Collector) {
	if err := prometheus.Register(c); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			_ = are.ExistingCollector
		} else {
			panic(err)
		}
	}
}

func init() {
	safeRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	safeRegister(collectors.NewGoCollector())
}

func StartMetricsServer(cfg *MetricsConfig, logger *zap.SugaredLogger) {
	mux := http.NewServeMux()
	mux.Handle(cfg.Path, promhttp.Handler())
	server := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		logger.Infow("metrics server listening on " + cfg.Host + ":" + cfg.Port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatalw("metrics server crashed", "error", err)
		}
	}()
}
