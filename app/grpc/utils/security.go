package utils

import (
	"crspy2/licenses/config"
	"encoding/base64"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/credentials"
	"os"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func GenerateSSLCert(l *zap.SugaredLogger) (credentials.TransportCredentials, error) {
	l.Info("Loading SSL encryption data")
	cert, err := base64.StdEncoding.DecodeString(config.Conf.SSL.Cert)
	if err != nil {
		return nil, errors.New("Failed to decode SSL encryption certificate: " + err.Error())

	}

	key, err := base64.StdEncoding.DecodeString(config.Conf.SSL.Key)
	if err != nil {
		return nil, errors.New("Failed to decode SSL encryption key: " + err.Error())
	}

	certFile := "/tmp/server.cert"
	keyFile := "/tmp/server.key"
	err = os.WriteFile(certFile, cert, 0644)
	if err != nil {
		return nil, errors.New("Failed to save decoded SSL certificate: " + err.Error())
	}
	err = os.WriteFile(keyFile, key, 0600)
	if err != nil {
		return nil, errors.New("Failed to save decoded SSL key: " + err.Error())
	}

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return nil, errors.New("Failed to set up TLS: " + err.Error())
	}
	l.Info("SSL encryption key loaded")

	return creds, nil
}
