package authentication

import (
	"context"
	"google.golang.org/grpc"
	"slices"
	"strings"
)

type CsrfInterceptor struct{}

func NewCsrfInterceptor() *CsrfInterceptor {
	return &CsrfInterceptor{}
}

// UnaryServerInterceptor CsrfInterceptor
func (i *CsrfInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if slices.Contains(UnauthedRoutes, info.FullMethod) || strings.HasPrefix(info.FullMethod, "/grpc") {
			return handler(ctx, req)
		}

		if err := checkCSRFToken(ctx); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor CsrfInterceptor
func (i *CsrfInterceptor) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if slices.Contains(UnauthedRoutes, info.FullMethod) || strings.HasPrefix(info.FullMethod, "/grpc") {
			return handler(srv, ss)
		}

		if err := checkCSRFToken(ss.Context()); err != nil {
			return err
		}
		return handler(srv, ss)
	}
}
