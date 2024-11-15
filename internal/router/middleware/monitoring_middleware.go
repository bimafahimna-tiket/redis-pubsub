package middleware

import (
	"net/http"
	"poc-redis-pubsub/internal/pkg/metric"
	"time"

	"github.com/labstack/echo/v4"
)

func Monitoring(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		metric.ActiveConnections.Inc()
		startTime := time.Now()
		err := next(ec)
		latencyTime := time.Since(startTime)
		metric.ActiveConnections.Dec()
		reqMethod := ec.Request().Method
		statusCode := ec.Response().Status
		metric.LatencyHistogram.WithLabelValues(reqMethod, ec.Request().RequestURI, http.StatusText(statusCode)).Observe(float64(latencyTime))
		metric.RequestCount.WithLabelValues(reqMethod, ec.Request().URL.Path, http.StatusText(statusCode)).Inc()
		return err
	}
}
