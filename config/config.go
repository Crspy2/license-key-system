package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type (
	Config struct {
		SentryDSN            string
		SessionEncryptionKey string
		Database             Database
		Ratelimit            Ratelimit
		Redis                Redis
	}

	Database struct {
		URI string
	}

	Ratelimit struct {
		Window int
		Max    int
	}

	Redis struct {
		Host     string
		Port     int
		Password string
		Threads  int
	}
)

var Conf Config

func LoadConfig() {
	rateLimitWindow, _ := strconv.Atoi(os.Getenv("RATELIMIT_WINDOW"))
	rateLimitMax, _ := strconv.Atoi(os.Getenv("RATELIMIT_MAX"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	redisThreads, _ := strconv.Atoi(os.Getenv("REDIS_THREADS"))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Conf = Config{
		SentryDSN:            os.Getenv("SENTRY_DSN"),
		SessionEncryptionKey: os.Getenv("SESSION_SECRET"),
		Database: Database{
			URI: os.Getenv("DATABASE_URL"),
		},
		Ratelimit: Ratelimit{
			Window: rateLimitWindow,
			Max:    rateLimitMax,
		},
		Redis: Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     redisPort,
			Password: os.Getenv("REDIS_PASSWORD"),
			Threads:  redisThreads,
		},
	}
}
