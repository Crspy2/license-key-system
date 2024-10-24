package utils

import (
	"context"
	"go.uber.org/zap"
)

func GetLogger(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value("logger").(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return logger
}
