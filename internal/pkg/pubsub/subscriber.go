package pubsub

type Subscriber interface {
	Subscribe(channel string) error
	Unsubscribe(channel string) error
}
