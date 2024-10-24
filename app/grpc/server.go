package grpc

import (
	"github.com/crspy2/license-panel/app/grpc/interceptors"
	"github.com/crspy2/license-panel/app/grpc/services"
	"github.com/crspy2/license-panel/pb/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func StartGRPCServer(l *zap.Logger) {
	l.Info("Starting GRPC server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		l.Panic("Failed to create listener: " + err.Error())
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryLoggingInterceptor(l),
			interceptors.UnaryRecoveryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			interceptors.StreamLoggingInterceptor(l),
			interceptors.StreamRecoveryInterceptor,
		),
	}

	s := grpc.NewServer(opts...)
	auth.RegisterAuthServer(s, &services.AuthServer{})

	reflection.Register(s)
	l.Info("GRPC server listening on tcp port :8080")
	if err = s.Serve(listener); err != nil {
		l.Panic("Failed to server: " + err.Error())
	}
}
