package dto

import (
	"time"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
)

type MetricDto struct {
	Entity       string
	ServiceGroup metrics.ServiceGroup
	ErrorCode    metrics.ErrorCode
	HttpCode     int
	StartTime    time.Time
}
