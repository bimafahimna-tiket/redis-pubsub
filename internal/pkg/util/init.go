package util

import (
	"poc-redis-pubsub/internal/config"
	"poc-redis-pubsub/internal/pkg/metric"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
)

var util utils

type utils struct {
	monitor metrics.MonitorStatsd
}

func Init(config *config.Config) {
	monitor := metric.NewMonitor(config)
	util.monitor = monitor
}
