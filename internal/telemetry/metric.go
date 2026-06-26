package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "lb_total_requests",
			Help: "Total numbers the requests http",
		},
		[]string{"backend", "status"},
	)

	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "lb_request_duration_seconds",
			Help:    "Latency the request in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"backend"},
	)
)
