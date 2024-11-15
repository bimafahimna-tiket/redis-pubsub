package service

import (
	"context"
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/mq"
	"poc-redis-pubsub/internal/pkg/pubsub"
)

type MessageService interface {
	SendMessage(ctx context.Context, message dto.MessageRequest) (string, error)
	SendMessagePubSub(ctx context.Context, message dto.MessagePubSubRequest) (string, error)
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
	if err := s.publisher.PublishMessage(ctx, "msg", message.Msg); err != nil {
		return "", err
	}
	return "Success Sending Message", nil
}
