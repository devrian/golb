package monitor

import "github.com/prometheus/client_golang/prometheus"

var (
	httpLatencyHistogram      *prometheus.HistogramVec
	httpResponsesTotalCounter *prometheus.CounterVec
	httpMetricLabels          = []string{"handler", "method", "httpcode", "env"}

	grpcLatencyHistogram      *prometheus.HistogramVec
	grpcResponsesTotalCounter *prometheus.CounterVec
	grpcMetricLabels          = []string{"handler", "grpccode", "grpcstatus", "env"}

	consumerLatencyHistogram      *prometheus.HistogramVec
	consumerResponsesTotalCounter *prometheus.CounterVec
	consumerMetricLabels          = []string{"topic", "group", "status", "env"}
)
