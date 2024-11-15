package main

import (
	"context"
	"os"
	"os/signal"
	"poc-redis-pubsub/internal/config"
	"poc-redis-pubsub/internal/pkg/logger"
	"poc-redis-pubsub/internal/pkg/mq"
	"syscall"

	"github.com/hibiken/asynq"
)

func main() {
	config := config.InitConfig()
	logger.SetSlogLogger()
	opt := asynq.RedisClientOpt{
		Addr: config.Redis.ServerAddress,
	}

	consumer := mq.NewTaskConsumer(opt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := consumer.Start(ctx); err != nil {
			logger.Log.Errorf("failed to start worker: %s", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	logger.Log.Infof("shutting down worker...")
	consumer.Shutdown()

	logger.Log.Infof("worker shutdown gracefully")
}
