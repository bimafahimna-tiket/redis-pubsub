package config

import (
	"log"
	"os"
	"time"
)

type AppConfig struct {
	ServerAddress     string
	ServerGracePeriod time.Duration
}

func initAppConfig() AppConfig {
	gracePeriod, err := time.ParseDuration(os.Getenv("SERVER_GRACE_PERIOD"))
	if err != nil {
		log.Fatal("failed to parse SERVER_GRACE_PERIOD")
	}

	return AppConfig{
		ServerAddress:     os.Getenv("SERVER_ADDRESS"),
		ServerGracePeriod: gracePeriod,
	}
}
