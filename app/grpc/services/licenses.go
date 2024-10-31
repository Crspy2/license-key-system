package services

import (
	"context"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type LicenseServer struct {
	pf.UnimplementedLicenseServer
}

func (s *LicenseServer) GetLicense(ctx context.Context, in *pf.LicenseKeyRequest) (*pf.LicenseObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.DefaultPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to view license information")
	}

	licenseKey := in.GetLicenseKey()
	if licenseKey == "" {
		return nil, status.Errorf(codes.InvalidArgument, "License key is required")
	}

	license, err := database.Client.Licenses.Get(licenseKey)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "License could not be found")
	}

	var activationTimestamp *timestamppb.Timestamp
	if license.Activation.Valid {
		activationTimestamp = timestamppb.New(license.Activation.Time)
	} else {
		activationTimestamp = nil
	}

	var expirationTimestamp *timestamppb.Timestamp
	if license.Expiration.Valid {
		expirationTimestamp = timestamppb.New(license.Expiration.Time)
	} else {
		expirationTimestamp = nil
	}

	var pausedAt *timestamppb.Timestamp
	if license.Product.PausedAt.Valid {
		pausedAt = timestamppb.New(license.Product.PausedAt.Time)
	} else {
		pausedAt = nil
	}

	return &pf.LicenseObject{
		Id:               license.ID,
		Duration:         durationpb.New(license.Duration),
		TimesCompensated: uint64(license.TimesCompensated),
		HoursCompensated: uint64(license.HoursCompensated),
		Activation:       activationTimestamp,
		Expiration:       expirationTimestamp,
		User: &pf.UserObject{
			Id:     license.User.ID,
			Name:   license.User.Name,
			Banned: license.User.Banned,
		},
		Product: &pf.ProductObject{
			Id:       license.Product.ID,
			Name:     license.Product.Name,
			Status:   license.Product.Status,
			PausedAt: pausedAt,
		},
	}, nil
}

func (s *LicenseServer) UserLicenseKeyStream(in *pf.UserIdRequest, stream pf.License_UserLicenseKeyStreamServer) error {
	session := stream.Context().Value("session").(*database.SessionModel)
	if session == nil {
		return status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.DefaultPermission) {
		return status.Errorf(codes.PermissionDenied, "You do not have permission to view license keys")
	}

	userId := in.GetUserId()
	if userId >= 0 {
		return status.Errorf(codes.InvalidArgument, "A user id must be provided")
	}

	users, err := database.Client.Users.Get(userId)
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, license := range users.Licenses {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		var activationTimestamp *timestamppb.Timestamp
		if license.Activation.Valid {
			activationTimestamp = timestamppb.New(license.Activation.Time)
		} else {
			activationTimestamp = nil
		}

		var expirationTimestamp *timestamppb.Timestamp
		if license.Expiration.Valid {
			expirationTimestamp = timestamppb.New(license.Expiration.Time)
		} else {
			expirationTimestamp = nil
		}

		var pausedAt *timestamppb.Timestamp
		if license.Product.PausedAt.Valid {
			pausedAt = timestamppb.New(license.Product.PausedAt.Time)
		} else {
			pausedAt = nil
		}

		p := &pf.LicenseObject{
			Id:               license.ID,
			Duration:         durationpb.New(license.Duration),
			TimesCompensated: uint64(license.TimesCompensated),
			HoursCompensated: uint64(license.HoursCompensated),
			Activation:       activationTimestamp,
			Expiration:       expirationTimestamp,
			User: &pf.UserObject{
				Id:     license.User.ID,
				Name:   license.User.Name,
				Banned: license.User.Banned,
			},
			Product: &pf.ProductObject{
				Id:       license.Product.ID,
				Name:     license.Product.Name,
				Status:   license.Product.Status,
				PausedAt: pausedAt,
			},
		}

		if err = stream.Send(p); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *LicenseServer) CreateLicense(ctx context.Context, in *pf.CreateLicenseRequest) (*pf.LicenseObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.KeyGenPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to generate keys")
	}

	productId := in.GetProductId()
	if productId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "License key is required")
	}
	duration := in.GetDuration()
	if duration == nil {
		return nil, status.Errorf(codes.InvalidArgument, "License key is required")
	}

	license, err := database.Client.Licenses.Create(productId, duration.AsDuration())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product could not be found")
	}

	var pausedAt *timestamppb.Timestamp
	if license.Product.PausedAt.Valid {
		pausedAt = timestamppb.New(license.Product.PausedAt.Time)
	} else {
		pausedAt = nil
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "License Key", "License Key Generated", fmt.Sprintf("%s has generated a %fh key for %s", session.Staff.Name, license.Duration.Hours(), license.Product.Name), time.Now())

	return &pf.LicenseObject{
		Id:               license.ID,
		Duration:         durationpb.New(license.Duration),
		TimesCompensated: uint64(license.TimesCompensated),
		HoursCompensated: uint64(license.HoursCompensated),
		Activation:       nil,
		Expiration:       nil,
		User: &pf.UserObject{
			Id:     license.User.ID,
			Name:   license.User.Name,
			Banned: license.User.Banned,
		},
		Product: &pf.ProductObject{
			Id:       license.Product.ID,
			Name:     license.Product.Name,
			Status:   license.Product.Status,
			PausedAt: pausedAt,
		},
	}, nil
}

func (s *LicenseServer) RevokeUserKeys(ctx context.Context, in *pf.UserIdRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageUsersPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to manage users")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "A user id must be provided")
	}

	user, err := database.Client.Licenses.Revoke(userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable to revoke license keys")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "License Key", "License Keys Revoked", fmt.Sprintf("%s has revoked license keys on %s's account", session.Staff.Name, user.Name), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("License keys on %s's account have been revoked", user.Name),
	}, nil
}

func (s *LicenseServer) RedeemLicense(ctx context.Context, in *pf.RedeemLicenseRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.KeyGenPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to generate keys")
	}

	licenseKey := in.GetLicenseKey()
	if licenseKey == "" {
		return nil, status.Errorf(codes.InvalidArgument, "License key is required")
	}

	userId := in.GetUserId()
	if userId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "License key is required")
	}

	license, err := database.Client.Licenses.Redeem(licenseKey, userId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable to redeem license key")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "License Key", "License Key Manually Redeemed", fmt.Sprintf("%s has manually redeemed a %fh key for %s on the account with UID=%d", session.Staff.Name, license.Duration.Hours(), license.Product.Name, userId), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("Successfully redeemed license key on %s's account", license.User.Name),
	}, nil
}
