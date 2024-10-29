package authentication

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"slices"
	"strings"
)

type AuthInterceptor struct{}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

// UnaryServerInterceptor AuthInterceptor
func (i *AuthInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if slices.Contains(UnauthedRoutes, info.FullMethod) || strings.HasPrefix(info.FullMethod, "/grpc") {
			return handler(ctx, req)
		}

		newCtx, err := authorizeSession(ctx)
		if err != nil {
			return nil, err
		}
		fmt.Println("TESTING SHIT")
		return handler(newCtx, req)
	}
}

// StreamServerInterceptor AuthInterceptor
func (i *AuthInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if slices.Contains(UnauthedRoutes, info.FullMethod) || strings.HasPrefix(info.FullMethod, "/grpc") {
			return handler(srv, ss)
		}

		newCtx, err := authorizeSession(ss.Context())
		if err != nil {
			return err
		}

		wrappedStream := &WrappedStream{
			ServerStream: ss,
			wrappedCtx:   newCtx,
		}

		// Pass the wrapped stream to the handler
		return handler(srv, wrappedStream)
	}
}

type WrappedStream struct {
	grpc.ServerStream
	wrappedCtx context.Context
}

func (w *WrappedStream) Context() context.Context {
	return w.wrappedCtx
}
