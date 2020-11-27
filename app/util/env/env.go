package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile(filePath string) error {
	return godotenv.Load(filePath)
}

func GetVar(key string) (string, error) {
	if value := os.Getenv(key); value != "" {
		return value, nil
	}

	return "", fmt.Errorf("empty [%s] env variable", key)
}

func GetVarWithFallback(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
