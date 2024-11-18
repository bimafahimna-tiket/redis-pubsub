package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"poc-redis-pubsub/internal/cache"
	"poc-redis-pubsub/internal/domain/dto"
	"poc-redis-pubsub/internal/pkg/logger"
	"poc-redis-pubsub/internal/pkg/util"
	"strconv"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/metrics"
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
		wrapErr := fmt.Errorf("failed to subscribe channel %s: %v", channel, err)
		logger.Log.Errorf(wrapErr.Error())
		return wrapErr
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

	unsubscribe := make(chan os.Signal, 1)

	go func() {
		logger.Log.Infof("Subscribed to channel %s", channel)
		for {
			select {
			case <-ctx.Done():
				unsubscribe <- syscall.SIGINT
				return
			case msg := <-conn.Channel():
				var p Payload
				err := json.Unmarshal([]byte(msg.Payload), &p)
				if err != nil {
					logger.Log.Errorf("error unmarshalling")
					return
				}
				if p.Type == TypeCache {
					metric := dto.MetricDto{
						Entity:       "Cache",
						ServiceGroup: metrics.API_IN,
						ErrorCode:    metrics.Success,
						HttpCode:     0,
						CustomTag:    map[string]interface{}{"Msg-Id": strconv.FormatInt(p.UniqueID, 10)},
						StartTime:    time.Now(),
					}
					util.SendMetricLatency(metric)
					switch {
					case p.Operation == OperationAdd:
						cache.Cache[p.Msg] = nil
					case p.Operation == OperationRemove:
						delete(cache.Cache, p.Msg)
					}
					logger.Log.Info("Cache updated: ", p.Msg)
				} else {
					logger.Log.Infof("Processing message from %s channel: %s", channel, p.Msg)
				}
			}
		}
	}()
	stopListen := make(chan os.Signal, 1)
	signal.Notify(stopListen, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stopListen:
		defer cancel()
		conn.Close()
	case <-unsubscribe:
		conn.Close()
	}
	logger.Log.Infof("Unsubscribed from channel %s", channel)
}
