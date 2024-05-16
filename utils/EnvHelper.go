package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func GetFromEnv(key string) (*string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	value := os.Getenv(key)

	return &value, nil
}
