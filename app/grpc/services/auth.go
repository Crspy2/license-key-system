package services

import (
	"context"
	"errors"
	"github.com/crspy2/license-panel/app/grpc/utils"
	"github.com/crspy2/license-panel/database"
	"github.com/crspy2/license-panel/pb/auth"
	"go.jetify.com/typeid"
	"time"
)

type AuthServer struct {
	auth.UnimplementedAuthServer
}

func (s *AuthServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	username := in.GetUsername()

	if len(username) < 3 {
		return nil, errors.New("username must be at least 3 characters in length")
	}

	password := in.GetPassword()

	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters in length")
	}

	staffMember, err := database.Client.Staff.Authenticate(username, password)
	if err != nil {
		return nil, err
	}

	ip := in.GetIp()
	userAgent := in.GetUserAgent()

	sessionToken := typeid.Must(typeid.WithPrefix("sess")).String()
	sessionInfo := database.SessionModal{
		Id:        sessionToken,
		StaffId:   staffMember.Id,
		IpAddress: ip,
		UserAgent: userAgent,
		ExpiresAt: time.Now().Add(5 * time.Hour),
	}

	_ = database.Client.Session.DeleteByIP(ip)
	err = database.Client.Session.Create(&sessionInfo)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Message: "Successfully created database session",
		Data: &auth.LoginResponse_ResponseData{
			SessionId: sessionToken,
		},
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, in *auth.RegisterRequest) (*auth.StandardResponse, error) {
	username := in.GetUsername()

	if len(username) < 3 {
		return nil, errors.New("username must be at least 3 characters in length")
	}

	password := in.GetPassword()

	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters in length")
	}

	staffMember, _ := database.Client.Staff.GetByName(username)
	if staffMember != nil {
		return nil, errors.New("this username is already in use, please choose another one")
	}

	hashedPassword, _ := utils.HashPassword(password)

	staffMember, err := database.Client.Staff.Create(username, hashedPassword)
	if err != nil {
		return nil, err
	}

	return &auth.StandardResponse{
		Message: "Registration complete. Wait while a staff member approves your account",
	}, nil
}

func (s *AuthServer) Logout(ctx context.Context, in *auth.LogoutRequest) (*auth.StandardResponse, error) {
	sessionId := in.GetSessionId()
	if sessionId == "" {
		return nil, errors.New("no session id found")
	}

	err := database.Client.Session.Delete(sessionId)
	if err != nil {
		return nil, err
	}

	return &auth.StandardResponse{
		Message: "The session has been deleted",
	}, nil
}

func (s *AuthServer) GetSessionInfo(ctx context.Context, in *auth.SingleSessionRequest) (*auth.SingleSessionResponse, error) {
	sessionId := in.GetSessionId()
	ip := in.GetIp()

	if sessionId == "" || ip == "" {
		return nil, errors.New("invalid procedure call")
	}

	session, err := database.Client.Session.Get(sessionId, ip)
	if err != nil {
		return nil, err
	}
	return &auth.SingleSessionResponse{
		Message: "Retrieved session information",
		Data: &auth.SessionObject{
			Id:        session.Id,
			IpAddress: session.IpAddress,
			UserAgent: session.UserAgent,
			Staff: &auth.StaffObject{
				Id:           session.Staff.Id,
				Name:         session.Staff.Name,
				PasswordHash: session.Staff.PasswordHash,
				Approved:     session.Staff.Approved,
				Perms:        session.Staff.PermsToString(),
			},
		},
	}, nil
}

func (s *AuthServer) GetUserSessionsStream(in *auth.MultiSessionRequest, stream auth.Auth_GetUserSessionsStreamServer) error {
	//logger := utils.GetLogger(stream.Context())

	userId := in.GetUserId()
	if userId == "" {
		return errors.New("invalid procedure call")
	}

	sessions, err := database.Client.Session.GetUserSessions(userId)
	if err != nil {
		//logger.Fatal(err.Error())
		return err
	}

	for _, session := range sessions {
		// Check if the context has been canceled
		select {
		case <-stream.Context().Done():
			//logger.Info("Context canceled, stopping stream")
			return stream.Context().Err()
		default:
			// Continue processing
		}

		// Map the GORM model to the Protobuf message
		sessionItem := &auth.SessionObject{
			Id:        session.Id,
			IpAddress: session.IpAddress,
			UserAgent: session.UserAgent,
			Staff: &auth.StaffObject{
				Id:           session.Staff.Id,
				Name:         session.Staff.Name,
				PasswordHash: session.Staff.PasswordHash,
				Approved:     session.Staff.Approved,
				Perms:        session.Staff.PermsToString(),
			},
		}

		// Send the item to the client
		if err = stream.Send(sessionItem); err != nil {
			//logger.Error("Failed to send item", zap.Error(err))
			return err
		}
	}

	//logger.Info("Successfully streamed items")
	return nil
}

func (s *AuthServer) RevokeSession(ctx context.Context, in *auth.SessionRevokeRequest) (*auth.StandardResponse, error) {
	sessionId := in.GetId()

	if sessionId == "" {
		return nil, errors.New("invalid procedure call")
	}

	err := database.Client.Session.Delete(sessionId)
	if err != nil {
		return nil, err
	}

	return &auth.StandardResponse{
		Message: "Session Revoked",
	}, nil
}
