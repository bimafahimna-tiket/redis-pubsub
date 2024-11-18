package config

import (
	"log"
	"os"
	"strconv"
)

type ContainerConfig struct {
	TotalContainer int
}

func initContainerConfig() ContainerConfig {
	total, err := strconv.Atoi(os.Getenv("TOTAL_CONTAINER"))
	if err != nil {
		log.Fatal("failed to parse TOTAL_CONTAINER")
	}

	return ContainerConfig{
		TotalContainer: total,
	}
}
