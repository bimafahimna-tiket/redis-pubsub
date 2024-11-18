package router

import (
	"net/http"
	"poc-redis-pubsub/internal/application/controller"
	"poc-redis-pubsub/internal/router/middleware"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(opts *controller.ControllerOpts) http.Handler {
	e := echo.New()
	e.Use(echo_middleware.Recover(), middleware.Logger())

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG")
	})
	e.POST("/message", opts.SendMessage)
	e.POST("/pubsub/subscribe", opts.SubscribeToChannel)
	e.POST("/pubsub/unsubscribe", opts.UnsubscribeToChannel)
	e.GET("/cache", opts.GetAllCache)

	m := e.Group("", middleware.Monitoring)
	m.POST("/pubsub/message", opts.SendMessagePubSub)
	e.PATCH("/pubsub/cache", opts.UpdateCache)

	return e
}
