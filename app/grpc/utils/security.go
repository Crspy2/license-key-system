package utils

import (
	"crspy2/licenses/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/credentials"
	"io"
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

// Retrieve encryption key from the environment variable
func getEncryptionKey() ([]byte, error) {
	keyBase64 := config.Conf.CookieEncryptionKey
	if keyBase64 == "" {
		return nil, errors.New("encryption Key not set")
	}

	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the encryption key: %v", err)
	}

	if len(key) != 32 {
		return nil, errors.New("encryption key must be 32 bytes for AES-256")
	}

	return key, nil
}

// EncryptToken encrypts text using AES-256-GCM and return the Base64-encoded result
func EncryptToken(text string) (string, error) {
	key, err := getEncryptionKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %v", err)
	}

	// Encrypt and add authentication tag
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(text), nil)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptToken decrypts a Base64-encoded ciphertext using AES-256-GCM
func DecryptToken(encryptedText string) (string, error) {
	key, err := getEncryptionKey()
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %v", err)
	}

	return string(plaintext), nil
}
