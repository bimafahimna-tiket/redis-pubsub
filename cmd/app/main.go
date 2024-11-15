package main

import (
	"poc-redis-pubsub/internal/application/controller"
	"poc-redis-pubsub/internal/config"
	"poc-redis-pubsub/internal/pkg/logger"
	"poc-redis-pubsub/internal/router"
	"poc-redis-pubsub/internal/server"
)

func main() {
	logger.SetSlogLogger()
	config := config.InitConfig()
	opts := controller.Init(config)
	route := router.Init(opts)

	s := server.NewServer(config, route)
	s.Run()
}
