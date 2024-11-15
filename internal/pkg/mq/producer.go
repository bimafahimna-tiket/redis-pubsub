package mq

import (
	"context"
)

type Producer interface {
	ProducePrintMessageTask(ctx context.Context, payload *MessagePayload) error
}
