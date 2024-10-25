package main

import (
	"crspy2/licenses/app/grpc"
	"crspy2/licenses/config"
	"crspy2/licenses/database"
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

	config.LoadConfig(sugar)

	sugar.Infoln("Connecting to database...")
	database.ConnectToDatabase()
	sugar.Infoln("Database connection established")

	grpc.StartGRPCServer(sugar)
}
