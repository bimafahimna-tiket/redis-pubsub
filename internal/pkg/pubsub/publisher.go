package pubsub

import "context"

type Publisher interface {
	PublishMessage(ctx context.Context, channel, msg string) error
}
