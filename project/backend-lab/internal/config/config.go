package config

import (
	"os"
	"time"
)

type Config struct {
	Address         string
	ShutdownTimeout time.Duration
}

func Load() Config {
	address := os.Getenv("TASK_API_ADDRESS")
	if address == "" {
		address = "127.0.0.1:8080"
	}
	return Config{Address: address, ShutdownTimeout: 5 * time.Second}
}
