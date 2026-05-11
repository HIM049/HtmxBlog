package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Cfg *Config

type Config struct {
	Database Database
	Service  Service
	Settings map[string]string
}

// Init loads configuration from environment variables.
func Init() {
	_ = godotenv.Load()

	Cfg = &Config{
		Database: ReadDatabase(),
		Service:  ReadService(),
	}
}

// getEnv reads an environment variable, panics if not set.
func getEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	panic("environment variable not found: " + key)
}
