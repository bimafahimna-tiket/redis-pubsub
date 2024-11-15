package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"poc-redis-pubsub/internal/config"
	"poc-redis-pubsub/internal/pkg/logger"
	"poc-redis-pubsub/internal/pkg/mq"
	"syscall"
	"time"

	"github.com/hibiken/asynq"
)

type Server struct {
	http.Server
	GracePeriod time.Duration
}

func NewServer(config *config.Config, handler http.Handler) Server {
	return Server{
		Server: http.Server{
			Addr:    config.App.ServerAddress,
			Handler: handler,
		},
		GracePeriod: config.App.ServerGracePeriod,
	}
}

func (s *Server) Run() {
	go func() {
		logger.Log.Infof("listening on %s...", s.Addr)
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Errorf("failed to start run server: %v", err)
		}
		logger.Log.Infof("server is not receiving new requests...")
	}()

	// go s.StartConsumer()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	timeout := time.Duration(s.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	<-stop

	logger.Log.Infof("shutting down server...")
	if err := s.Shutdown(ctx); err != nil {
		logger.Log.Errorf("failed to shutdown server: %v", err)
	}

	logger.Log.Infof("server shutdown gracefully")
}

func (s *Server) StartConsumer() {
	logger.Log.Info("Start Worker")
	opt := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_SERVER_ADDRESS"),
	}

	consumer := mq.NewTaskConsumer(opt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := consumer.Start(ctx); err != nil {
			logger.Log.Errorf("failed to start worker: %s", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	logger.Log.Infof("shutting down worker...")
	consumer.Shutdown()

	logger.Log.Infof("worker shutdown gracefully")
}
