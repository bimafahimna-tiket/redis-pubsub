package config

import (
	"log"
	"os"
	"time"
)

type AppConfig struct {
	ServerHost        string
	ServerPort        string
	ServerAddress     string
	ServerGracePeriod time.Duration
}

func initAppConfig() AppConfig {
	gracePeriod, err := time.ParseDuration(os.Getenv("SERVER_GRACE_PERIOD"))
	if err != nil {
		log.Fatal("failed to parse SERVER_GRACE_PERIOD")
	}

	return AppConfig{
		ServerHost:        os.Getenv("SERVER_DOMAIN"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		ServerAddress:     os.Getenv("SERVER_ADDRESS"),
		ServerGracePeriod: gracePeriod,
	}
}
