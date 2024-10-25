package services

import (
	"context"
	"crspy2/licenses/app/grpc/utils"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"go.jetify.com/typeid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type AuthServer struct {
	pf.UnimplementedAuthServer
}

func (s *AuthServer) Login(ctx context.Context, in *pf.LoginRequest) (*pf.LoginResponse, error) {
	username := in.GetUsername()

	if len(username) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "username must be at least 3 characters in length")
	}

	password := in.GetPassword()

	if len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "password must be at least 8 characters in length")
	}

	staffMember, err := database.Client.Staff.Authenticate(username, password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
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
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}

	return &pf.LoginResponse{
		Message: "Successfully created database session",
		Data: &pf.LoginResponse_ResponseData{
			SessionId: sessionToken,
		},
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, in *pf.RegisterRequest) (*pf.StandardResponse, error) {
	username := in.GetUsername()

	if len(username) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "username must be at least 3 characters in length")
	}

	password := in.GetPassword()

	if len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "password must be at least 8 characters in length")
	}

	staffMember, _ := database.Client.Staff.GetByName(username)
	if staffMember != nil {
		return nil, status.Errorf(codes.AlreadyExists, "this username is already in use, please choose another one")
	}

	hashedPassword, _ := utils.HashPassword(password)

	staffMember, err := database.Client.Staff.Create(username, hashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}

	return &pf.StandardResponse{
		Message: "Registration complete. Wait while a staff member approves your account",
	}, nil
}

func (s *AuthServer) Logout(ctx context.Context, in *pf.LogoutRequest) (*pf.StandardResponse, error) {
	sessionId := in.GetSessionId()
	if sessionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "no session id found")
	}

	err := database.Client.Session.Delete(sessionId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pf.StandardResponse{
		Message: "The session has been deleted",
	}, nil
}

func (s *AuthServer) GetSessionInfo(ctx context.Context, in *pf.SingleSessionRequest) (*pf.SingleSessionResponse, error) {
	sessionId := in.GetSessionId()
	ip := in.GetIp()

	if sessionId == "" || ip == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid procedure call")
	}

	session, err := database.Client.Session.Get(sessionId, ip)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pf.SingleSessionResponse{
		Message: "Retrieved session information",
		Data: &pf.SessionObject{
			Id:        session.Id,
			IpAddress: session.IpAddress,
			UserAgent: session.UserAgent,
			Staff: &pf.StaffObject{
				Id:           session.Staff.Id,
				Name:         session.Staff.Name,
				PasswordHash: session.Staff.PasswordHash,
				Approved:     session.Staff.Approved,
				Perms:        session.Staff.GetPermissionNames(),
			},
		},
	}, nil
}

func (s *AuthServer) GetUserSessionsStream(in *pf.MultiSessionRequest, stream pf.Auth_GetUserSessionsStreamServer) error {
	staffId := in.GetStaffId()
	if staffId == "" {
		return status.Errorf(codes.InvalidArgument, "invalid procedure call")
	}

	sessions, err := database.Client.Session.GetUserSessions(staffId)
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, session := range sessions {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		sessionItem := &pf.SessionObject{
			Id:        session.Id,
			IpAddress: session.IpAddress,
			UserAgent: session.UserAgent,
			Staff: &pf.StaffObject{
				Id:           session.Staff.Id,
				Name:         session.Staff.Name,
				PasswordHash: session.Staff.PasswordHash,
				Approved:     session.Staff.Approved,
				Perms:        session.Staff.GetPermissionNames(),
			},
		}

		if err = stream.Send(sessionItem); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *AuthServer) RevokeSession(ctx context.Context, in *pf.SessionRevokeRequest) (*pf.StandardResponse, error) {
	sessionId := in.GetId()

	if sessionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid procedure call")
	}

	err := database.Client.Session.Delete(sessionId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pf.StandardResponse{
		Message: "Session revoked successfully",
	}, nil
}
