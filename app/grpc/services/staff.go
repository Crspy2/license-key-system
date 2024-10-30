package services

import (
	"context"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type StaffServer struct {
	pf.UnimplementedStaffServer
}

func (s *StaffServer) SetStaffAccess(ctx context.Context, in *pf.StaffAccessRequest) (*pf.ApprovalResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageStaffPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to perform this action")
	}

	staffId := in.GetStaffId()
	approved := in.GetApproved()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	if session.Staff.ID == staffId {
		return nil, status.Errorf(codes.PermissionDenied, "You cannot modify your own permissions")
	}

	staff, err := database.Client.Staff.GetById(staffId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Staff member could not be found")
	}

	if !session.Staff.HasHigherPermissions(*staff) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to perform this action")
	}

	staff, err = database.Client.Staff.SetAccess(staffId, approved)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "Staff", "Staff Access Approved", fmt.Sprintf("%s has approved %s's access the panel", session.Staff.Name, staff.Name), time.Now())

	return &pf.ApprovalResponse{
		Message: fmt.Sprintf("%s's access to the panel has been updated", staff.Name),
		Staff: &pf.StaffObject{
			Id:       staff.ID,
			Name:     staff.Name,
			Role:     staff.Role,
			Perms:    staff.GetPermissionNames(),
			Approved: staff.Approved,
		},
	}, nil
}

func (s *StaffServer) GetStaff(ctx context.Context, in *pf.StaffIdRequest) (*pf.StaffObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageStaffPermission) && session.Staff.Role < database.DevRole {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to view staff")
	}

	staffId := in.GetStaffId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	staff, err := database.Client.Staff.GetById(staffId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pf.StaffObject{
		Id:       staff.ID,
		Name:     staff.Name,
		Role:     staff.Role,
		Perms:    staff.GetPermissionNames(),
		Approved: staff.Approved,
	}, nil
}

func (s *StaffServer) ListStaffStream(_ *empty.Empty, stream pf.Staff_ListStaffStreamServer) error {
	session := stream.Context().Value("session").(*database.SessionModel)
	if session == nil {
		return status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageStaffPermission) && session.Staff.Role < database.DevRole {
		return status.Errorf(codes.PermissionDenied, "You do not have permission to view staff")
	}

	staffMembers, err := database.Client.Staff.List()
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, member := range staffMembers {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		staffMember := &pf.StaffObject{
			Id:       member.ID,
			Name:     member.Name,
			Role:     member.Role,
			Perms:    member.GetPermissionNames(),
			Approved: member.Approved,
		}

		if err = stream.Send(staffMember); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *StaffServer) SetStaffPermissions(ctx context.Context, in *pf.MultiPermissionRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageStaffPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to update staff permissions")
	}

	staffId := in.GetStaffId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	if session.Staff.ID == staffId {
		return nil, status.Errorf(codes.PermissionDenied, "You cannot modify your own permissions")
	}

	staff, err := database.Client.Staff.GetById(staffId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Staff member could not be found")
	}

	if !session.Staff.HasHigherRole(*staff) || !session.Staff.HasHigherPermissions(*staff) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to perform this action")
	}

	perms := in.GetPermissions()

	var permissions []database.Permission

	for _, perm := range perms {
		permissions = append(permissions, database.Permission(perm))
	}

	var newPerms database.Permission
	for _, permission := range permissions {
		newPerms |= permission
	}
	staff.Perms = newPerms

	if !session.Staff.HasHigherPermissions(*staff) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to perform this action")
	}

	staff, err = database.Client.Staff.SetPermissions(staffId, permissions)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable to set user's permissions")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "Staff", "Staff Permissions Updated", fmt.Sprintf("%s has updated %s's permissions on the panel. They are now %s", session.Staff.Name, staff.Name, strings.Join(staff.GetPermissionNames(), ", ")), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("Successfully overwrote permissions for %s", staff.Name),
	}, nil
}

func (s *StaffServer) SetStaffRole(ctx context.Context, in *pf.StaffRoleRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageStaffPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to perform this action")
	}

	staffId := in.GetStaffId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A staff id is required")
	}
	if session.Staff.ID == staffId {
		return nil, status.Errorf(codes.PermissionDenied, "You cannot modify your own role")
	}

	staff, err := database.Client.Staff.GetById(staffId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Staff member could not be found")
	}

	if !session.Staff.HasHigherRole(*staff) || !session.Staff.HasHigherPermissions(*staff) {
		return nil, status.Errorf(codes.PermissionDenied, "#1 You do not have permission to perform this action")
	}

	role := in.GetRole()
	if role == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "The role parameter is required")
	}

	// Check if role would result in the user gaining more permissions than the current user
	staff.Role = role
	if !session.Staff.HasHigherRole(*staff) {
		return nil, status.Errorf(codes.PermissionDenied, "You cannot promote a user to a higher role than yourself")
	}

	staff, err = database.Client.Staff.SetRole(staffId, role)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable to set user's role")
	}

	_, _ = database.Client.Logs.LogEvent(session.StaffID, "Staff", "Staff Permissions Updated", fmt.Sprintf("%s has updated %s's role on the panel to %s", session.Staff.Name, staff.Name, staff.GetRoleText()), time.Now())

	return &pf.StandardResponse{
		Message: fmt.Sprintf("Successfully set %s's role", staff.Name),
	}, nil
}
