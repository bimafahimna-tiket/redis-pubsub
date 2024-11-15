package metric

import (
	"poc-redis-pubsub/internal/config"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
)

// type Holder struct {
// 	Monitor metrics.MonitorStatsd
// }

func NewMonitor(config *config.Config) metrics.MonitorStatsd {
	monitorStatsD, _ := metrics.NewMonitor(
		"localhost",
		"8125",
		"REDIS-PUBSUB",
	)
	// return Holder{Monitor: monitorStatsD}
	return monitorStatsD
}
