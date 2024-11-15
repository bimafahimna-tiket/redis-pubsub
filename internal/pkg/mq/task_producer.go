package mq

import (
	"context"
	"encoding/json"
	"poc-redis-pubsub/internal/pkg/logger"

	"github.com/hibiken/asynq"
)

type taskProducer struct {
	client *asynq.Client
}

func NewTaskProducer(opt asynq.RedisClientOpt) Producer {
	client := asynq.NewClient(opt)
	return &taskProducer{
		client: client,
	}
}

func (p *taskProducer) ProducePrintMessageTask(ctx context.Context, payload *MessagePayload) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	task := asynq.NewTask(taskPrintMessage, jsonPayload, payload.Opts...)

	info, err := p.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}
	logger.Log.Info(
		"TYPE: ", task.Type(),
		"PAYLOAD: ", string(task.Payload()),
		"QUEUE: ", info.Queue,
		"MAX_RETRY: ", info.MaxRetry,
	)
	return nil
}
