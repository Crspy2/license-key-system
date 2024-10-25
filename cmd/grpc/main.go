package main

import (
	"crspy2/licenses/app/grpc"
	"crspy2/licenses/config"
	"crspy2/licenses/database"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func main() {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalln("Failed to create Zap logger instance")
	}

	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Fatalln("Failed to terminate Zap logger instance")
		}
	}(logger)

	sugar := logger.Sugar()

	sugar.Infoln("Initialized Zap Logging")

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") != "production" {
		sugar.Infoln("Loading environment variables...")
		err = godotenv.Load()
		if err != nil {
			sugar.Fatalln("Error loading .env file")
		}
		config.LoadConfig()
		sugar.Infoln("Configuration files loaded")
	}

	sugar.Infoln("Connecting to database...")
	database.ConnectToDatabase()
	sugar.Infoln("Database connection established")

	grpc.StartGRPCServer(sugar)
}
