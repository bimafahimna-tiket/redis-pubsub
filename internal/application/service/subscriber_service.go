package service

import (
	"context"
	"poc-redis-pubsub/internal/pkg/pubsub"
)

type SubscribeService interface {
	SubscribeTo(ctx context.Context, channel string) (string, error)
	UnsubscribeTo(ctx context.Context, channel string) (string, error)
}

type subscribeService struct {
	subscriber pubsub.Subscriber
}

func NewSubscribeService(subscriber pubsub.Subscriber) SubscribeService {
	return &subscribeService{
		subscriber: subscriber,
	}
}

func (s *subscribeService) SubscribeTo(ctx context.Context, channel string) (string, error) {
	if err := s.subscriber.Subscribe(channel); err != nil {
		return "", err
	}
	return "Success to subscribe", nil
}

func (s *subscribeService) UnsubscribeTo(ctx context.Context, channel string) (string, error) {
	if err := s.subscriber.Unsubscribe(channel); err != nil {
		return "", err
	}
	return "Success to unsubscribe", nil
}
