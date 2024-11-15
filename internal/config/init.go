package config

import (
	"poc-redis-pubsub/internal/pkg/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	App   AppConfig
	Redis RedisConfig
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("failed to load .env")
	}
	return &Config{
		App:   initAppConfig(),
		Redis: initRedisConfig(),
	}
}
