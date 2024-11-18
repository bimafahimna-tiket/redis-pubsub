package middleware

import (
	"fmt"
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/util"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
)

// func Monitoring(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ec echo.Context) error {
// 		metric.ActiveConnections.Inc()
// 		startTime := time.Now()
// 		err := next(ec)
// 		latencyTime := time.Since(startTime)
// 		metric.ActiveConnections.Dec()
// 		reqMethod := ec.Request().Method
// 		statusCode := ec.Response().Status
// 		metric.LatencyHistogram.WithLabelValues(reqMethod, ec.Request().RequestURI, http.StatusText(statusCode)).Observe(float64(latencyTime))
// 		metric.RequestCount.WithLabelValues(reqMethod, ec.Request().URL.Path, http.StatusText(statusCode)).Inc()
// 		return err
// 	}
// }

func Monitoring(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		startTime := time.Now()
		err := next(ec)
		var errorCode metrics.ErrorCode
		if err != nil {
			errorCode = metrics.Failed
		} else {
			errorCode = metrics.Success
		}
		entity := fmt.Sprint(ec.Get("entity"))
		metric := dto.MetricDto{
			Entity:       entity,
			ServiceGroup: metrics.API_IN,
			HttpCode:     ec.Response().Status,
			ErrorCode:    errorCode,
			StartTime:    startTime,
		}
		util.SendMetricLatency(metric)
		return err
	}
}
