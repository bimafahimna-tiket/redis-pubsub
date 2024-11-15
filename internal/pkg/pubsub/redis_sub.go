package pubsub

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"poc-redis-pubsub/internal/pkg/logger"
	"syscall"

	"github.com/go-redis/redis/v8"
)

type redisSub struct {
	client     *redis.Client
	connection map[string]*redis.PubSub
	cancelFunc map[string]context.CancelFunc
}

func NewRedisSub(opt *redis.Options) Subscriber {
	client := redis.NewClient(opt)
	return &redisSub{
		client:     client,
		connection: make(map[string]*redis.PubSub),
		cancelFunc: make(map[string]context.CancelFunc),
	}
}

func (s *redisSub) Subscribe(channel string) error {
	if _, ok := s.connection[channel]; ok {
		return fmt.Errorf("already subscribe to channel")
	}
	ctx := context.Background()
	conn := s.client.Subscribe(ctx, channel)
	s.connection[channel] = conn
	if err := conn.Ping(ctx); err != nil {
		return fmt.Errorf("failed to subscribe channel %s: %v", channel, err)
	}
	go s.listen(ctx, channel)
	return nil
}

func (s *redisSub) Unsubscribe(channel string) error {
	conn, err := s.getConnection(channel)
	if err != nil {
		return fmt.Errorf("failed to unsubscribe channel %s: %v", channel, err)
	}
	conn.Close()
	s.cancelFunc[channel]()
	s.removeConnection(channel)
	return nil
}

func (s *redisSub) getConnection(channel string) (*redis.PubSub, error) {
	conn, ok := s.connection[channel]
	if !ok {
		return nil, fmt.Errorf("not subscribing to channel %s", channel)
	}
	return conn, nil
}

func (s *redisSub) removeConnection(channel string) {
	delete(s.connection, channel)
	delete(s.cancelFunc, channel)
}

func (s *redisSub) listen(ctx context.Context, channel string) {
	conn, _ := s.getConnection(channel)
	ctx, cancel := context.WithCancel(ctx)
	s.cancelFunc[channel] = cancel
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Log.Infof("Subscribed to channel %s", channel)
		for {
			select {
			case <-ctx.Done():
				stop <- syscall.SIGINT
				return
			case msg := <-conn.Channel():
				logger.Log.Infof("Processing message from %s channel: %s", channel, msg.Payload)
			}
		}
	}()

	<-stop
	conn.Close()
	logger.Log.Infof("Unsubscribed from channel %s", channel)

}
