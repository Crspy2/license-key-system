package services

import (
	"context"
	"crspy2/licenses/app/grpc/utils"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

type AuthServer struct {
	pf.UnimplementedAuthServer
}

func (s *AuthServer) Login(ctx context.Context, in *pf.LoginRequest) (*pf.LoginResponse, error) {
	username := in.GetUsername()

	if len(username) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "Username must be at least 3 characters in length")
	}

	password := in.GetPassword()

	if len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "Password must be at least 8 characters in length")
	}

	staffMember, err := database.Client.Staff.Authenticate(username, password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Missing metadata")
	}

	ip, ok := md["x-forwarded-for"]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Missing IP address")
	}

	userAgent, ok := md["x-client-user-agent"]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Missing IP address")
	}

	sessionInfo := database.SessionModel{
		StaffID:   staffMember.ID,
		IpAddress: ip[0],
		UserAgent: userAgent[0],
		ExpiresAt: time.Now().Add(5 * time.Hour),
	}

	_ = database.Client.Sessions.DeleteByIP(ip[0])
	err = database.Client.Sessions.Create(&sessionInfo)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}

	encryptedSessionToken, err := utils.EncryptToken(sessionInfo.ID)
	if err != nil {
		return nil, err
	}

	encryptedCsrfToken, err := utils.EncryptToken(sessionInfo.CsrfToken)
	if err != nil {
		return nil, err
	}

	if !sessionInfo.Staff.Approved {
		return nil, status.Errorf(codes.PermissionDenied, "Your account has not yet been approved")
	}

	return &pf.LoginResponse{
		Message: "Successfully created database session",
		Data: &pf.LoginResponse_ResponseData{
			SessionId: encryptedSessionToken,
			CsrfToken: encryptedCsrfToken,
		},
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, in *pf.RegisterRequest) (*pf.StandardResponse, error) {
	username := in.GetUsername()

	if len(username) < 3 {
		return nil, status.Errorf(codes.InvalidArgument, "Username must be at least 3 characters in length")
	}

	password := in.GetPassword()

	if len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "Password must be at least 8 characters in length")
	}

	staffMember, _ := database.Client.Staff.GetByName(username)
	if staffMember != nil {
		return nil, status.Errorf(codes.AlreadyExists, "This username is already in use, please choose another one")
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

func (s *AuthServer) Logout(ctx context.Context, _ *empty.Empty) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	err := database.Client.Sessions.Delete(session)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pf.StandardResponse{
		Message: "The session has been deleted",
	}, nil
}

func (s *AuthServer) GetSessionInfo(ctx context.Context, _ *empty.Empty) (*pf.SingleSessionResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	return &pf.SingleSessionResponse{
		Message: "Retrieved session information",
		Data: &pf.SessionObject{
			Id:        session.ID,
			IpAddress: session.IpAddress,
			UserAgent: session.UserAgent,
			Staff: &pf.StaffObject{
				Id:       session.Staff.ID,
				Name:     session.Staff.Name,
				Role:     session.Staff.Role,
				Image:    &session.Staff.Image,
				Perms:    session.Staff.GetPermissionNames(),
				Approved: session.Staff.Approved,
			},
		},
	}, nil
}

func (s *AuthServer) ListSessionStream(in *pf.MultiSessionRequest, stream pf.Auth_ListSessionStreamServer) error {
	staffId := in.GetStaffId()
	if staffId == "" {
		return status.Errorf(codes.InvalidArgument, "Invalid procedure call")
	}

	sessions, err := database.Client.Sessions.ListUserSessions(staffId)
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	if len(sessions) == 0 {
		return status.Errorf(codes.NotFound, "No sessions found for this user")
	}

	for _, session := range sessions {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		sessionItem := &pf.SessionObject{
			Id:        session.ID,
			IpAddress: session.IpAddress,
			UserAgent: session.UserAgent,
			Staff: &pf.StaffObject{
				Id:       session.Staff.ID,
				Name:     session.Staff.Name,
				Role:     session.Staff.Role,
				Image:    &session.Staff.Image,
				Perms:    session.Staff.GetPermissionNames(),
				Approved: session.Staff.Approved,
			},
		}

		if err = stream.Send(sessionItem); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *AuthServer) RevokeSession(ctx context.Context, in *pf.SessionRevokeRequest) (*pf.StandardResponse, error) {
	sessionId := in.GetSessionId()

	if sessionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid procedure call")
	}

	session, err := database.Client.Sessions.Get(sessionId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	err = database.Client.Sessions.Delete(session)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pf.StandardResponse{
		Message: "Session revoked successfully",
	}, nil
}
