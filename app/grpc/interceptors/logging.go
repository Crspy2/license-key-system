package interceptors

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type LoggingInterceptor struct{}

// NewLoggingInterceptor returns a new AuthInterceptor instance
func NewLoggingInterceptor() *LoggingInterceptor {
	return &LoggingInterceptor{}
}

// UnaryServerInterceptor LoggingInterceptor logs unary requests and responses using zap.Logger
func (i *LoggingInterceptor) UnaryServerInterceptor(l *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestLogger := l.WithOptions(zap.AddCallerSkip(3)).With(zap.String("procedure", info.FullMethod))

		ctx = context.WithValue(ctx, "logger", requestLogger)

		start := time.Now()

		// Log the incoming request with -->
		requestLogger.Info("\u001B[1m\u001B[32m---> \u001B[0m gRPC unary request received ",
			zap.Any("request", req),
		)

		// Call the handler to proceed with the request
		resp, err := handler(ctx, req)

		// Log the response with <-- and time taken
		duration := time.Since(start)
		if err != nil {
			requestLogger.Error("\u001B[1m\u001B[31m<--- \u001B[0m gRPC unary request failed ",
				zap.Any("request", req),
				zap.Error(err),
				zap.Duration("duration", duration),
			)
		} else {
			requestLogger.Info("\u001B[1m\u001B[32m<--- \u001B[0m gRPC unary request completed ",
				zap.Any("response", resp),
				zap.Duration("duration", duration),
			)
		}
		return resp, err
	}
}

// StreamServerInterceptor LoggingInterceptor returns a new streaming server interceptor that adds logging.
func (i *LoggingInterceptor) StreamServerInterceptor(l *zap.SugaredLogger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// Create a request-specific logger
		requestLogger := l.WithOptions(zap.AddCallerSkip(3)).With(zap.String("procedure", info.FullMethod))

		// Inject the logger into the context
		ctx := ss.Context()
		ctx = context.WithValue(ctx, "logger", requestLogger)

		// Wrap the ServerStream to override the Context method
		wrappedStream := &wrappedServerStream{
			ServerStream: ss,
			ctx:          ctx,
		}

		start := time.Now()

		// Log the incoming stream request
		requestLogger.Info("\u001B[1m\u001B[32m---> \u001B[0m gRPC stream request received ",
			zap.String("method", info.FullMethod),
		)

		// Call the handler to proceed with the request
		err := handler(srv, wrappedStream)

		// Log the end of the stream with duration
		duration := time.Since(start)
		if err != nil {
			requestLogger.Error("\u001B[1m\u001B[31m<--- \u001B[0m gRPC stream request failed ",
				zap.Error(err),
				zap.Duration("duration", duration),
			)
		} else {
			requestLogger.Info("\u001B[1m\u001B[32m<--- \u001B[0m gRPC stream request completed ",
				zap.Duration("duration", duration),
			)
		}
		return err
	}
}

// wrappedServerStream is a thin wrapper around grpc.ServerStream that allows us to override the Context method.
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

// Context returns the wrapper's overridden context.
func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

// RecvMsg forwards the call to the underlying ServerStream.
func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	return w.ServerStream.RecvMsg(m)
}

// SendMsg forwards the call to the underlying ServerStream.
func (w *wrappedServerStream) SendMsg(m interface{}) error {
	return w.ServerStream.SendMsg(m)
}
