package main

import (
	"github.com/crspy2/license-panel/app/grpc"
	"github.com/crspy2/license-panel/config"
	"github.com/crspy2/license-panel/database"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
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

	sugar.Infoln("Loading environment variables...")
	err = godotenv.Load()
	if err != nil {
		sugar.Panicln("Error loading .env file")
	}
	config.LoadConfig()
	sugar.Infoln("Configuration files loaded")

	sugar.Infoln("Connecting to database...")
	database.ConnectToDatabase()
	sugar.Infoln("Database connection established")

	grpc.StartGRPCServer(logger)
}
