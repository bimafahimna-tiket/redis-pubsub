package metric

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "statusCode"},
	)
	ActiveConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Sum of active connections",
		},
	)
	LatencyHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response latency (seconds) for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "uri", "status"},
	)
)

func init() {
	prometheus.MustRegister(RequestCount, ActiveConnections, LatencyHistogram)
}
