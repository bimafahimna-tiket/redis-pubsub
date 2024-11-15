package mq

import (
	"context"

	"github.com/hibiken/asynq"
)

type Consumer interface {
	Start(ctx context.Context) error
	PrintMessage(ctx context.Context, task *asynq.Task) error
	Shutdown()
}
