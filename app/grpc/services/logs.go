package services

import (
	"context"
	"crspy2/licenses/database"
	pf "crspy2/licenses/proto/protofiles"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LogServer struct {
	pf.UnimplementedLogServer
}

func (s *LogServer) GetLog(ctx context.Context, in *pf.LogIdRequest) (*pf.LogObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ViewLogsPermission) && session.Staff.Role < database.DevRole {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to view logs")
	}

	logId := in.GetLogId()

	if logId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A log id is required")
	}

	log, err := database.Client.Logs.Get(logId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pf.LogObject{
		Id:          log.ID,
		Object:      log.Object,
		Title:       log.Title,
		Description: log.Description,
		OccurredAt:  timestamppb.New(log.OccurredAt),
		Staff: &pf.StaffObject{
			Id:       log.Staff.ID,
			Name:     log.Staff.Name,
			Role:     log.Staff.Role,
			Perms:    log.Staff.GetPermissionNames(),
			Approved: log.Staff.Approved,
		},
	}, nil

}

func (s *LogServer) ListLogsStream(_ *empty.Empty, stream pf.Log_ListLogsStreamServer) error {
	session := stream.Context().Value("session").(*database.SessionModel)
	if session == nil {
		return status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ViewLogsPermission) && session.Staff.Role < database.DevRole {
		return status.Errorf(codes.PermissionDenied, "You do not have permission to view logs")
	}

	logs, err := database.Client.Logs.List()
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, log := range logs {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		l := &pf.LogObject{
			Id:          log.ID,
			Object:      log.Object,
			Title:       log.Title,
			Description: log.Description,
			OccurredAt:  timestamppb.New(log.OccurredAt),
			Staff: &pf.StaffObject{
				Id:       log.Staff.ID,
				Name:     log.Staff.Name,
				Role:     log.Staff.Role,
				Perms:    log.Staff.GetPermissionNames(),
				Approved: log.Staff.Approved,
			},
		}

		if err = stream.Send(l); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}
