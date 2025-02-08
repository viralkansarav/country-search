package config

import (
	"os"
)

var (
	Port = getEnv("PORT", "8000")
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
