package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type (
	Config struct {
		GrpcPort            string
		Database            Database
		SSL                 SSL
		CookieEncryptionKey string
	}

	SSL struct {
		Cert string
		Key  string
	}

	Database struct {
		URI string
	}
)

var Conf Config

func LoadConfig(s *zap.SugaredLogger) {
	s.Infoln("Loading environment variables...")
	_ = godotenv.Load()

	Conf = Config{
		GrpcPort: os.Getenv("PORT"),
		Database: Database{
			URI: os.Getenv("DATABASE_URL"),
		},
		SSL: SSL{
			Cert: os.Getenv("SSL_ENCRYPTION_CERT"),
			Key:  os.Getenv("SSL_ENCRYPTION_KEY"),
		},
		CookieEncryptionKey: os.Getenv("COOKIE_ENCRYPTION_KEY"),
	}

	s.Infoln("Configuration files loaded")
}
