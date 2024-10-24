package interceptors

import (
	"context"
	"github.com/crspy2/license-panel/app/grpc/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"

	"google.golang.org/grpc"
)

func UnaryRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger := utils.GetLogger(ctx)
			logger.Error("Panic occurred",
				zap.String("procedure", info.FullMethod),
				zap.Any("panic", r),
				zap.ByteString("stack", debug.Stack()),
			)
			err = status.Errorf(codes.Internal, "Internal server error")
		}
	}()
	return handler(ctx, req)
}

// StreamRecoveryInterceptor for streaming RPCs
func StreamRecoveryInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			ctx := ss.Context()
			logger := utils.GetLogger(ctx)
			logger.Error("Panic occurred",
				zap.Any("panic", r),
				zap.ByteString("stack", debug.Stack()),
			)
			err = status.Errorf(codes.Internal, r.(error).Error())
		}
	}()
	return handler(srv, ss)
}
