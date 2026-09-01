package config

import (
	"crypto/sha256"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var requiredKeys = []string{
	"APP_KEY",
	"APP_TZ",
	"SERVER_HOST",
	"SERVER_PORT",
	"MOVIE_API_URL",
	"MOVIE_API_KEY",
}

func InitializeAndSet() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: copy .env.example to .env and fill in the values")
	}

	for _, key := range requiredKeys {
		if Get(key) == "" {
			log.Fatalf("Missing required environment variable: %s", key)
		}
	}

	location, err := time.LoadLocation(Get("APP_TZ"))

	if err != nil {
		log.Fatalf("Err loading location: %v", err)
	}

	time.Local = location
}

func Get(key string) string {
	return os.Getenv(key)
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func GetRunAddress() string {
	return Get("SERVER_HOST") + ":" + Get("SERVER_PORT")
}

// GetAppKey derives a stable 32-byte AES-256 key from the APP_KEY env value.
func GetAppKey() []byte {
	sum := sha256.Sum256([]byte(Get("APP_KEY")))
	return sum[:]
}
