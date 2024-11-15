package config

import "os"

type RedisConfig struct {
	ServerAddress string
}

func initRedisConfig() RedisConfig {
	return RedisConfig{
		ServerAddress: os.Getenv("REDIS_SERVER_ADDRESS"),
	}
}
