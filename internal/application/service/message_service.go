package service

import (
	"context"
	"poc-redis-pubsub/internal/cache"
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/logger"
	"poc-redis-pubsub/internal/pkg/mq"
	"poc-redis-pubsub/internal/pkg/pubsub"
)

type MessageService interface {
	SendMessage(ctx context.Context, message dto.MessageRequest) (string, error)
	SendMessagePubSub(ctx context.Context, message dto.MessagePubSubRequest) (string, error)
	GetAllCache(ctx context.Context) (dto.GetAllCacheResponse, error)
	UpdateCache(ctx context.Context, cache dto.UpdateCacheRequest) (string, error)
}

type messageService struct {
	producer  mq.Producer
	publisher pubsub.Publisher
}

func NewMessageService(producer mq.Producer, publisher pubsub.Publisher) MessageService {
	return &messageService{
		producer:  producer,
		publisher: publisher,
	}
}

func (s *messageService) SendMessage(ctx context.Context, message dto.MessageRequest) (string, error) {
	msg := mq.MessagePayload{
		Msg: message.Msg,
	}
	err := s.producer.ProducePrintMessageTask(ctx, &msg)
	if err != nil {
		return "", err
	}
	return "Success Sending Message", nil
}

func (s *messageService) SendMessagePubSub(ctx context.Context, message dto.MessagePubSubRequest) (string, error) {
	data := pubsub.NewJsonPayload(pubsub.TypeMessage, "", message.Msg)
	if err := s.publisher.PublishMessage(ctx, "msg", data); err != nil {
		return "", err
	}
	return "Success Sending Message", nil
}

func (s *messageService) GetAllCache(ctx context.Context) (dto.GetAllCacheResponse, error) {
	var res []string
	for key := range cache.Cache {
		res = append(res, key)
	}
	return dto.GetAllCacheResponse{Cache: res}, nil
}

func (s *messageService) UpdateCache(ctx context.Context, cache dto.UpdateCacheRequest) (string, error) {
	data := pubsub.NewJsonPayload(pubsub.TypeCache, cache.Operation, cache.Cache)
	err := s.publisher.PublishCache(ctx, "cache", data)
	if err != nil {
		logger.Log.Error("failed to update cache")
		return "Failed to update cache", nil
	}
	return "Successfully update cache", nil
}
