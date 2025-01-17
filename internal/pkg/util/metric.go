package util

import (
	"log"
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/logger"
	"time"
)

func SendMetricLatency(dto dto.MetricDto) {
	log.Println("SEND")
	latencyTime := time.Since(dto.StartTime)
	CustomTag := make(map[string]interface{})
	err := util.monitor.CustomMonitorLatency(dto.Entity, dto.ServiceGroup, dto.ErrorCode, dto.HttpCode, CustomTag, latencyTime)
	if err != nil {
		logger.Log.Errorf("failed to send metric: %v", err)
	}
	logger.Log.Info("Send Metric to monitor via influxdb")
}
