package pubsub

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type redisPub struct {
	client *redis.Client
}

func NewRedisPub(opt *redis.Options) Publisher {
	client := redis.NewClient(opt)
	return &redisPub{
		client: client,
	}
}

func (p *redisPub) PublishMessage(ctx context.Context, channel, msg string) error {
	if err := p.client.Publish(ctx, channel, msg).Err(); err != nil {
		return fmt.Errorf("failed to publish message to channel %s: %v", channel, err)
	}
	return nil
}
