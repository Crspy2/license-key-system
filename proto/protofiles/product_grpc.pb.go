// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: product.proto

package protofiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Product_GetProduct_FullMethodName        = "/protofiles.Product/GetProduct"
	Product_ListProductStream_FullMethodName = "/protofiles.Product/ListProductStream"
	Product_CreateProduct_FullMethodName     = "/protofiles.Product/CreateProduct"
	Product_DeleteProduct_FullMethodName     = "/protofiles.Product/DeleteProduct"
	Product_CompensateProduct_FullMethodName = "/protofiles.Product/CompensateProduct"
	Product_SetProductStatus_FullMethodName  = "/protofiles.Product/SetProductStatus"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	GetProduct(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*ProductObject, error)
	ListProductStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProductObject], error)
	CreateProduct(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductObject, error)
	DeleteProduct(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	CompensateProduct(ctx context.Context, in *ProductCompRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	SetProductStatus(ctx context.Context, in *ProductStatusRequest, opts ...grpc.CallOption) (*StandardResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) GetProduct(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*ProductObject, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductObject)
	err := c.cc.Invoke(ctx, Product_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListProductStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ProductObject], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Product_ServiceDesc.Streams[0], Product_ListProductStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, ProductObject]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Product_ListProductStreamClient = grpc.ServerStreamingClient[ProductObject]

func (c *productClient) CreateProduct(ctx context.Context, in *ProductCreateRequest, opts ...grpc.CallOption) (*ProductObject, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductObject)
	err := c.cc.Invoke(ctx, Product_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *ProductIdRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, Product_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CompensateProduct(ctx context.Context, in *ProductCompRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, Product_CompensateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) SetProductStatus(ctx context.Context, in *ProductStatusRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, Product_SetProductStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility.
type ProductServer interface {
	GetProduct(context.Context, *ProductIdRequest) (*ProductObject, error)
	ListProductStream(*emptypb.Empty, grpc.ServerStreamingServer[ProductObject]) error
	CreateProduct(context.Context, *ProductCreateRequest) (*ProductObject, error)
	DeleteProduct(context.Context, *ProductIdRequest) (*StandardResponse, error)
	CompensateProduct(context.Context, *ProductCompRequest) (*StandardResponse, error)
	SetProductStatus(context.Context, *ProductStatusRequest) (*StandardResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServer struct{}

func (UnimplementedProductServer) GetProduct(context.Context, *ProductIdRequest) (*ProductObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServer) ListProductStream(*emptypb.Empty, grpc.ServerStreamingServer[ProductObject]) error {
	return status.Errorf(codes.Unimplemented, "method ListProductStream not implemented")
}
func (UnimplementedProductServer) CreateProduct(context.Context, *ProductCreateRequest) (*ProductObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *ProductIdRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) CompensateProduct(context.Context, *ProductCompRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompensateProduct not implemented")
}
func (UnimplementedProductServer) SetProductStatus(context.Context, *ProductStatusRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProductStatus not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}
func (UnimplementedProductServer) testEmbeddedByValue()                 {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	// If the following call pancis, it indicates UnimplementedProductServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct(ctx, req.(*ProductIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListProductStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServer).ListProductStream(m, &grpc.GenericServerStream[emptypb.Empty, ProductObject]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Product_ListProductStreamServer = grpc.ServerStreamingServer[ProductObject]

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*ProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*ProductIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CompensateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCompRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CompensateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CompensateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CompensateProduct(ctx, req.(*ProductCompRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_SetProductStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).SetProductStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_SetProductStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).SetProductStatus(ctx, req.(*ProductStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _Product_GetProduct_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "CompensateProduct",
			Handler:    _Product_CompensateProduct_Handler,
		},
		{
			MethodName: "SetProductStatus",
			Handler:    _Product_SetProductStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProductStream",
			Handler:       _Product_ListProductStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "product.proto",
}
