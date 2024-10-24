package grpc

import (
	"github.com/crspy2/license-panel/app/grpc/interceptors"
	"github.com/crspy2/license-panel/app/grpc/services"
	"github.com/crspy2/license-panel/pb/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
)

func StartGRPCServer(l *zap.SugaredLogger) {
	l.Info("Starting GRPC server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		l.Panic("Failed to create listener: " + err.Error())
	}

	recoverOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			return status.Errorf(codes.Unknown, p.(string))
		}),
	}

	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryLoggingInterceptor(l),
			recovery.UnaryServerInterceptor(recoverOpts...),
		),
		grpc.ChainStreamInterceptor(
			interceptors.StreamLoggingInterceptor(l),
			recovery.StreamServerInterceptor(recoverOpts...),
		),
	}

	s := grpc.NewServer(grpcOpts...)
	auth.RegisterAuthServer(s, &services.AuthServer{})

	reflection.Register(s)
	l.Info("GRPC server listening on tcp port :8080")
	if err = s.Serve(listener); err != nil {
		l.Panic("Failed to server: " + err.Error())
	}
}
