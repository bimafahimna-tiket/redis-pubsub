package mq

import (
	"context"
	"encoding/json"
	"poc-redis-pubsub/internal/pkg/logger"

	"github.com/hibiken/asynq"
)

type taskConsumer struct {
	server *asynq.Server
}

func NewTaskConsumer(opt asynq.RedisClientOpt) Consumer {
	server := asynq.NewServer(opt, asynq.Config{
		Queues: map[string]int{
			"default":  30,
			"critical": 15,
			"low":      5,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			logger.Log.Errorf("FAILED TO PROCESS TASK", "TYPE", task.Type(), "PAYLOAD", task.Payload())
		}),
		Logger: logger.Log,
	})

	return &taskConsumer{
		server: server,
	}
}

func (c *taskConsumer) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mux := asynq.NewServeMux()

			mux.HandleFunc(taskPrintMessage, c.PrintMessage)

			return c.server.Start(mux)
		}
	}
}

func (c *taskConsumer) PrintMessage(ctx context.Context, task *asynq.Task) error {
	var payload MessagePayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	logger.Log.Info("SUCCESS: ", payload.Msg)
	return nil
}

func (c *taskConsumer) Shutdown() {
	c.server.Shutdown()
}
