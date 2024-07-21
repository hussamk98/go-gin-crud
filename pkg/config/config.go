package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load reads the .env file and loads it into the environment
func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	return nil
}

// GetEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func GetEnv(key string) string {
	return os.Getenv(key)
}
