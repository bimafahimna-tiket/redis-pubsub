package controller

import (
	"poc-redis-pubsub/internal/application/service"
	"poc-redis-pubsub/internal/config"
	"poc-redis-pubsub/internal/pkg/mq"
	"poc-redis-pubsub/internal/pkg/pubsub"
	"poc-redis-pubsub/internal/pkg/util"

	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
)

type ControllerOpts struct {
	*MessageController
	*SubscribeController
}

func Init(config *config.Config) *ControllerOpts {
	producer := mq.NewTaskProducer(asynq.RedisClientOpt{
		Addr: config.Redis.ServerAddress,
	})
	redisOpt := &redis.Options{
		Addr: config.Redis.ServerAddress,
	}

	publisher := pubsub.NewRedisPub(redisOpt)
	subscriber := pubsub.NewRedisSub(redisOpt)
	subscriber.Subscribe("msg")

	util.Init(config)

	messageService := service.NewMessageService(producer, publisher)
	subscribeService := service.NewSubscribeService(subscriber)

	messageController := NewMessageController(messageService)
	subscribeController := NewSubscribeController(subscribeService)

	return &ControllerOpts{
		MessageController:   messageController,
		SubscribeController: subscribeController,
	}
}
