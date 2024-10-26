package grpc

import (
	"crspy2/licenses/app/grpc/interceptors"
	"crspy2/licenses/app/grpc/interceptors/authentication"
	"crspy2/licenses/app/grpc/services"
	"crspy2/licenses/app/grpc/utils"
	"crspy2/licenses/config"
	pf "crspy2/licenses/proto/protofiles"
	"encoding/base64"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"os"
)

func StartGRPCServer(l *zap.SugaredLogger) {
	l.Info("Starting GRPC server")
	listener, err := net.Listen("tcp", ":"+config.Conf.GrpcPort)
	if err != nil {
		l.Fatalf("Failed to create listener: %v", err.Error())
	}

	l.Info("Loading SSL encryption data")
	cert, err := base64.StdEncoding.DecodeString(config.Conf.SSL.Cert)
	if err != nil {
		l.Fatalf("Failed to decode SSL encryption certificate: %v", err)
	}

	key, err := base64.StdEncoding.DecodeString(config.Conf.SSL.Key)
	if err != nil {
		l.Fatalf("Failed to decode SSL encryption key: %v", err)
	}

	certFile := "/tmp/server.cert"
	keyFile := "/tmp/server.key"
	err = os.WriteFile(certFile, cert, 0644)
	if err != nil {
		l.Fatalf("Failed to save decoded SSL certificate: %v", err)
	}
	err = os.WriteFile(keyFile, key, 0600)
	if err != nil {
		l.Fatalf("Failed to save decoded SSL key: %v", err)
	}

	creds, err := utils.GenerateSSLCert(l)
	if err != nil {
		l.Fatalln(err.Error())
	}
	l.Info("SSL encryption key loaded")

	recoverOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			return status.Errorf(codes.Unknown, p.(string))
		}),
	}

	loggingInterceptor := interceptors.NewLoggingInterceptor()
	authInterceptor := authentication.NewAuthInterceptor()
	csrfInterceptor := authentication.NewCsrfInterceptor()

	grpcOpts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			loggingInterceptor.UnaryServerInterceptor(l),
			recovery.UnaryServerInterceptor(recoverOpts...),
			authInterceptor.UnaryServerInterceptor(),
			csrfInterceptor.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			loggingInterceptor.StreamServerInterceptor(l),
			recovery.StreamServerInterceptor(recoverOpts...),
			authInterceptor.StreamServerInterceptor(),
			csrfInterceptor.StreamServerInterceptor(),
		),
	}

	s := grpc.NewServer(grpcOpts...)
	pf.RegisterAuthServer(s, &services.AuthServer{})
	pf.RegisterStaffServer(s, &services.StaffServer{})

	// TODO: Remove this line in production
	reflection.Register(s)

	l.Infof("GRPC server listening on tcp port %s", config.Conf.GrpcPort)
	if err = s.Serve(listener); err != nil {
		l.Panic("Failed to server: " + err.Error())
	}
}
