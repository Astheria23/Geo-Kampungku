package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

// LoadEnv loads environment variables from .env file once.
func LoadEnv() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Printf("config: no .env file found or already loaded: %v", err)
		}
	})
}

// GetEnv fetches an environment variable or returns fallback when not set.
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
