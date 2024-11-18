package metric

import (
	"poc-redis-pubsub/internal/config"
	"poc-redis-pubsub/internal/pkg/logger"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
)

// type Holder struct {
// 	Monitor metrics.MonitorStatsd
// }

func NewMonitor(config *config.Config) metrics.MonitorStatsd {
	monitorStatsD, err := metrics.NewMonitor(
		config.Monitor.Domain,
		config.Monitor.Port,
		config.Monitor.ServiceName,
	)
	if err != nil {
		logger.Log.Errorf("failed connect to monitor stats d: ", err)
		return nil
	}
	// return Holder{Monitor: monitorStatsD}
	return monitorStatsD
}
