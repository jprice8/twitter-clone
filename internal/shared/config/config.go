package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load reads .env file config into env
func Load(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if err := godotenv.Load(path); err != nil {
		return err
	}

	return nil
}

func lookup(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetString(key string, fallback string) string {
	return lookup(key, fallback)
}

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
