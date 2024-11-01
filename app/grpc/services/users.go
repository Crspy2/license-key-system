package services

import (
	"context"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserServer struct {
	pf.UnimplementedUserServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pf.UserCreateRequest) (*pf.UserObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageUsersPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to create a user")
	}

	name := in.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A username is required for user creation")
	}

	password := in.GetPassword()
	if password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A password is required for user creation")
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user, err := database.Client.Users.Create(name, string(passwordBytes))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "A user with that name already exists")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "User", "Manual User Creation", fmt.Sprintf("%s has manually created a user account with username %s", session.Staff.Name, user.Name), time.Now())

	return &pf.UserObject{
		Id:     user.ID,
		Name:   user.Name,
		Banned: user.Banned,
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, in *pf.UserIdRequest) (*pf.UserObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.DefaultPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to view user information")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "User id is required")
	}

	user, err := database.Client.Users.Get(userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User could not be found")
	}

	return &pf.UserObject{
		Id:     user.ID,
		Name:   user.Name,
		Banned: user.Banned,
	}, nil
}

func (s *UserServer) ListUsersStream(_ *empty.Empty, stream pf.User_ListUsersStreamServer) error {
	session := stream.Context().Value("session").(*database.SessionModel)
	if session == nil {
		return status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.DefaultPermission) {
		return status.Errorf(codes.PermissionDenied, "You do not have permission to view user information")
	}

	users, err := database.Client.Users.List()
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, user := range users {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		u := &pf.UserObject{
			Id:     user.ID,
			Name:   user.Name,
			Banned: user.Banned,
		}

		if err = stream.Send(u); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *UserServer) ResetHardwareId(ctx context.Context, in *pf.UserIdRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.HWIDResetPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to reset Hardware IDs")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "User id is required")
	}

	user, err := database.Client.Users.ResetHWID(userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "User", "HWID Reset", fmt.Sprintf("%s has issued a hardware id reset for %s's account", session.Staff.Name, user.Name), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("%s's Hardware ID has been reset. Their account will lock to the next device they sign in on", user.Name),
	}, nil
}

func (s *UserServer) ResetPassword(ctx context.Context, in *pf.UserIdRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.PassResetPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to reset account passwords")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "User id is required")
	}

	user, err := database.Client.Users.ResetPassword(userId)
	if err != nil {
		return nil, err
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "User", "Password Reset", fmt.Sprintf("%s has issued a password reset for %s's account", session.Staff.Name, user.Name), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("%s's password has been reset. The next password they enter at login will become their new password", user.Name),
	}, nil
}

func (s *UserServer) BanUser(ctx context.Context, in *pf.UserIdRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageUsersPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to ban users")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "User id is required")
	}

	user, err := database.Client.Users.Ban(userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User could not be found")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "User", "User Ban Issued", fmt.Sprintf("%s has issued a ban to %s's account", session.Staff.Name, user.Name), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("A ban has been placed on %s's account. They will no longer be able to sign in", user.Name),
	}, nil
}

func (s *UserServer) RevokeBan(ctx context.Context, in *pf.UserIdRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageUsersPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to revoke user user bans")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "User id is required")
	}

	user, err := database.Client.Users.Unban(userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User could not be found")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "User", "User Ban Revoked", fmt.Sprintf("%s has revoked a ban to %s's account", session.Staff.Name, user.Name), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("The ban on %s's account has been removed", user.Name),
	}, nil
}
