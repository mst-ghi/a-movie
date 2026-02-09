package config

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func InitializeAndSet() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
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

func GetAppKey() string {
	key := []byte(Get("APP_KEY"))

	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}

	return hex.EncodeToString(key)
}
