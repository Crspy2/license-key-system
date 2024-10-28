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

func (s *StaffServer) ApproveStaff(_ context.Context, in *pf.StaffIdRequest) (*pf.ApprovalResponse, error) {
	staffId := in.GetStaffId()

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
			Image:        &staff.Image,
			PasswordHash: staff.PasswordHash,
			Role:         staff.Role,
			Perms:        staff.GetPermissionNames(),
			Approved:     staff.Approved,
		},
	}, nil
}

func (s *StaffServer) GetStaff(_ context.Context, in *pf.StaffIdRequest) (*pf.StaffObject, error) {
	staffId := in.GetStaffId()

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
		Image:        &staff.Image,
		PasswordHash: staff.PasswordHash,
		Role:         staff.Role,
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
			Image:        &member.Image,
			PasswordHash: member.PasswordHash,
			Role:         member.Role,
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
