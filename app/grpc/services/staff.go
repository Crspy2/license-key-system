package services

import (
	"context"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StaffServer struct {
	pf.UnimplementedStaffServer
}

func (s *StaffServer) ApproveStaff(_ context.Context, in *pf.StaffId) (*pf.ApprovalResponse, error) {
	staffId := in.GetId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	staff, err := database.Client.Staff.ApproveStaff(staffId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pf.ApprovalResponse{
		Message: "Successfully created database session",
		Staff: &pf.StaffObject{
			Id:           staff.Id,
			Name:         staff.Name,
			PasswordHash: staff.PasswordHash,
			Perms:        staff.GetPermissionNames(),
			Approved:     staff.Approved,
		},
	}, nil
}

func (s *StaffServer) GetStaff(_ context.Context, in *pf.StaffId) (*pf.StaffObject, error) {
	staffId := in.GetId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	staff, err := database.Client.Staff.GetById(staffId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pf.StaffObject{
		Id:           staff.Id,
		Name:         staff.Name,
		PasswordHash: staff.PasswordHash,
		Perms:        staff.GetPermissionNames(),
		Approved:     staff.Approved,
	}, nil
}

func (s *StaffServer) GetAllStaffStream(_ *empty.Empty, stream pf.Staff_GetAllStaffStreamServer) error {
	staffMembers, err := database.Client.Staff.GetAll()
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
			Id:           member.Id,
			Name:         member.Name,
			PasswordHash: member.PasswordHash,
			Perms:        member.GetPermissionNames(),
			Approved:     member.Approved,
		}

		if err = stream.Send(staffMember); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *StaffServer) SetStaffPermissions(_ context.Context, in *pf.MultiPermissionRequest) (*pf.StandardResponse, error) {
	staffId := in.GetStaffId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	perms := in.GetPermissions()

	if len(perms) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Permissions are required")
	}

	var permissions []database.Permission

	for _, perm := range perms {
		permissions = append(permissions, database.Permission(perm))
	}

	err := database.Client.Staff.SetPermissions(staffId, permissions)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pf.StandardResponse{
		Message: "Successfully overwrote staff permissions",
	}, nil
}

func (s *StaffServer) AddStaffPermission(_ context.Context, in *pf.SinglePermissionRequest) (*pf.StandardResponse, error) {
	staffId := in.GetStaffId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	perm := in.GetPermission()
	if perm == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "You must specify a permission to add")
	}

	err := database.Client.Staff.AddPermission(staffId, database.Permission(perm))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pf.StandardResponse{
		Message: "Successfully added a permission to the staff member",
	}, nil
}

func (s *StaffServer) RemoveStaffPermission(_ context.Context, in *pf.SinglePermissionRequest) (*pf.StandardResponse, error) {
	staffId := in.GetStaffId()

	if staffId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Staff Id is required")
	}

	perm := in.GetPermission()

	if perm == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "You must specify a permission to add")
	}

	err := database.Client.Staff.RemovePermission(staffId, database.Permission(perm))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pf.StandardResponse{
		Message: "Successfully removed a permission to the staff member",
	}, nil
}
