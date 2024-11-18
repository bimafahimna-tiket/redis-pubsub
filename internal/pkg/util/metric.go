package util

import (
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/logger"
	"time"
)

func SendMetricLatency(dto dto.MetricDto) {
	latencyTime := time.Since(dto.StartTime)
	err := util.monitor.CustomMonitorLatency(dto.Entity, dto.ServiceGroup, dto.ErrorCode, dto.HttpCode, dto.CustomTag, latencyTime)
	if err != nil {
		logger.Log.Errorf("failed to send metric: %v", err)
		return
	}
	logger.Log.Info("Send Metric to monitor via influxdb")
}
