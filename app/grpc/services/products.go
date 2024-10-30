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

type ProductServer struct {
	pf.UnimplementedProductServer
}

func (s *ProductServer) GetProduct(ctx context.Context, in *pf.ProductIdRequest) (*pf.ProductObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.DefaultPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to view products")
	}

	productId := in.GetProductId()
	if productId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Product Id is required")
	}

	product, err := database.Client.Products.GetById(productId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product could not be found")
	}

	var pausedAt *timestamppb.Timestamp
	if product.PausedAt.Valid {
		pausedAt = timestamppb.New(product.PausedAt.Time)
	} else {
		pausedAt = nil
	}
	return &pf.ProductObject{
		Id:       product.ID,
		Name:     product.Name,
		Status:   product.Status,
		PausedAt: pausedAt,
	}, nil
}

func (s *ProductServer) ListProductStream(_ *empty.Empty, stream pf.Product_ListProductStreamServer) error {
	session := stream.Context().Value("session").(*database.SessionModel)
	if session == nil {
		return status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.DefaultPermission) {
		return status.Errorf(codes.PermissionDenied, "You do not have permission to view products")
	}

	products, err := database.Client.Products.GetAll()
	if err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	for _, product := range products {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		var pausedAt *timestamppb.Timestamp
		if product.PausedAt.Valid {
			pausedAt = timestamppb.New(product.PausedAt.Time)
		} else {
			pausedAt = nil
		}

		p := &pf.ProductObject{
			Id:       product.ID,
			Name:     product.Name,
			Status:   product.Status,
			PausedAt: pausedAt,
		}

		if err = stream.Send(p); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}
	return nil
}

func (s *ProductServer) CreateProduct(ctx context.Context, in *pf.ProductCreateRequest) (*pf.ProductObject, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageProductsPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to create products")
	}

	productName := in.GetName()
	if productName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A product name is required")
	}

	product, err := database.Client.Products.Create(productName)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product could not be found")
	}

	var pausedAt *timestamppb.Timestamp
	if product.PausedAt.Valid {
		pausedAt = timestamppb.New(product.PausedAt.Time)
	} else {
		pausedAt = nil
	}

	return &pf.ProductObject{
		Id:       product.ID,
		Name:     product.Name,
		Status:   product.Status,
		PausedAt: pausedAt,
	}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, in *pf.ProductIdRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ManageProductsPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to delete products")
	}

	productId := in.GetProductId()
	if productId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A product id is required")
	}

	_, err := database.Client.Products.Delete(productId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product could not be found")
	}

	return &pf.StandardResponse{
		Message: "Product deleted",
	}, nil
}

func (s *ProductServer) CompensateProduct(ctx context.Context, in *pf.ProductCompRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.CompensationPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to compensate products")
	}

	productId := in.GetProductId()
	if productId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A product id is required")
	}

	compTime := in.GetCompTime()
	if compTime == nil {
		return nil, status.Errorf(codes.InvalidArgument, "A product status is required")
	}

	_, err := database.Client.Products.CompensateKeys(productId, compTime.AsDuration())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product could not be found")
	}

	return &pf.StandardResponse{
		Message: "Product status updated",
	}, nil
}

func (s *ProductServer) SetProductStatus(ctx context.Context, in *pf.ProductStatusRequest) (*pf.StandardResponse, error) {
	session := ctx.Value("session").(*database.SessionModel)
	if session == nil {
		return nil, status.Errorf(codes.Unauthenticated, "No session information found")
	}

	if !session.Staff.HasPermission(database.ChangeStatusPermission) {
		return nil, status.Errorf(codes.PermissionDenied, "You do not have permission to change product status")
	}

	productId := in.GetProductId()
	if productId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A product id is required")
	}

	productStatus := in.GetStatus()
	if productStatus == "" {
		return nil, status.Errorf(codes.InvalidArgument, "A product status is required")
	}

	_, err := database.Client.Products.SetProductStatus(productId, productStatus)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product could not be found")
	}

	return &pf.StandardResponse{
		Message: "Product status updated",
	}, nil
}
