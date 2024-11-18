package pubsub

import (
	"context"
	"fmt"
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/util"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
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
	metric := dto.MetricDto{
		Entity:       "Publish-Message",
		ServiceGroup: metrics.API_OUT,
		ErrorCode:    0,
		CustomTag:    nil,
		StartTime:    time.Now(),
	}
	if err := p.client.Publish(ctx, channel, msg).Err(); err != nil {
		metric.ErrorCode = metrics.Failed
		return fmt.Errorf("failed to publish message to channel %s: %v", channel, err)
	}
	customTag := map[string]interface{}{"Publish": ""}
	metric.CustomTag = customTag
	util.SendMetricLatency(metric)
	return nil
}

func (p *redisPub) PublishCache(ctx context.Context, channel, cache string) error {
	metric := dto.MetricDto{
		Entity:       "Cache",
		ServiceGroup: metrics.API_OUT,
		ErrorCode:    metrics.Success,
		CustomTag:    map[string]interface{}{"Published-Cache": cache},
		StartTime:    time.Now(),
	}
	if err := p.client.Publish(ctx, channel, cache).Err(); err != nil {
		metric.ErrorCode = metrics.Failed
		return fmt.Errorf("failed to publish message to channel %s: %v", channel, err)
	}
	// customTag := map[string]interface{}{"Publish-": ""}
	// metric.CustomTag = customTag
	util.SendMetricLatency(metric)
	return nil
}
