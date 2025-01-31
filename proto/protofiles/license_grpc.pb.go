// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: license.proto

package protofiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	License_GetLicense_FullMethodName           = "/protofiles.License/GetLicense"
	License_UserLicenseKeyStream_FullMethodName = "/protofiles.License/UserLicenseKeyStream"
	License_CreateLicense_FullMethodName        = "/protofiles.License/CreateLicense"
	License_RedeemLicense_FullMethodName        = "/protofiles.License/RedeemLicense"
	License_RevokeUserKeys_FullMethodName       = "/protofiles.License/RevokeUserKeys"
)

// LicenseClient is the client API for License service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LicenseClient interface {
	GetLicense(ctx context.Context, in *LicenseKeyRequest, opts ...grpc.CallOption) (*LicenseObject, error)
	UserLicenseKeyStream(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LicenseObject], error)
	CreateLicense(ctx context.Context, in *CreateLicenseRequest, opts ...grpc.CallOption) (*LicenseObject, error)
	RedeemLicense(ctx context.Context, in *RedeemLicenseRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	RevokeUserKeys(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*StandardResponse, error)
}

type licenseClient struct {
	cc grpc.ClientConnInterface
}

func NewLicenseClient(cc grpc.ClientConnInterface) LicenseClient {
	return &licenseClient{cc}
}

func (c *licenseClient) GetLicense(ctx context.Context, in *LicenseKeyRequest, opts ...grpc.CallOption) (*LicenseObject, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LicenseObject)
	err := c.cc.Invoke(ctx, License_GetLicense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) UserLicenseKeyStream(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LicenseObject], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &License_ServiceDesc.Streams[0], License_UserLicenseKeyStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UserIdRequest, LicenseObject]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type License_UserLicenseKeyStreamClient = grpc.ServerStreamingClient[LicenseObject]

func (c *licenseClient) CreateLicense(ctx context.Context, in *CreateLicenseRequest, opts ...grpc.CallOption) (*LicenseObject, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LicenseObject)
	err := c.cc.Invoke(ctx, License_CreateLicense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) RedeemLicense(ctx context.Context, in *RedeemLicenseRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, License_RedeemLicense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseClient) RevokeUserKeys(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, License_RevokeUserKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicenseServer is the server API for License service.
// All implementations must embed UnimplementedLicenseServer
// for forward compatibility.
type LicenseServer interface {
	GetLicense(context.Context, *LicenseKeyRequest) (*LicenseObject, error)
	UserLicenseKeyStream(*UserIdRequest, grpc.ServerStreamingServer[LicenseObject]) error
	CreateLicense(context.Context, *CreateLicenseRequest) (*LicenseObject, error)
	RedeemLicense(context.Context, *RedeemLicenseRequest) (*StandardResponse, error)
	RevokeUserKeys(context.Context, *UserIdRequest) (*StandardResponse, error)
	mustEmbedUnimplementedLicenseServer()
}

// UnimplementedLicenseServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLicenseServer struct{}

func (UnimplementedLicenseServer) GetLicense(context.Context, *LicenseKeyRequest) (*LicenseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLicense not implemented")
}
func (UnimplementedLicenseServer) UserLicenseKeyStream(*UserIdRequest, grpc.ServerStreamingServer[LicenseObject]) error {
	return status.Errorf(codes.Unimplemented, "method UserLicenseKeyStream not implemented")
}
func (UnimplementedLicenseServer) CreateLicense(context.Context, *CreateLicenseRequest) (*LicenseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLicense not implemented")
}
func (UnimplementedLicenseServer) RedeemLicense(context.Context, *RedeemLicenseRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemLicense not implemented")
}
func (UnimplementedLicenseServer) RevokeUserKeys(context.Context, *UserIdRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUserKeys not implemented")
}
func (UnimplementedLicenseServer) mustEmbedUnimplementedLicenseServer() {}
func (UnimplementedLicenseServer) testEmbeddedByValue()                 {}

// UnsafeLicenseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LicenseServer will
// result in compilation errors.
type UnsafeLicenseServer interface {
	mustEmbedUnimplementedLicenseServer()
}

func RegisterLicenseServer(s grpc.ServiceRegistrar, srv LicenseServer) {
	// If the following call pancis, it indicates UnimplementedLicenseServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&License_ServiceDesc, srv)
}

func _License_GetLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LicenseKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).GetLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: License_GetLicense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).GetLicense(ctx, req.(*LicenseKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_UserLicenseKeyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LicenseServer).UserLicenseKeyStream(m, &grpc.GenericServerStream[UserIdRequest, LicenseObject]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type License_UserLicenseKeyStreamServer = grpc.ServerStreamingServer[LicenseObject]

func _License_CreateLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).CreateLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: License_CreateLicense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).CreateLicense(ctx, req.(*CreateLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_RedeemLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeemLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).RedeemLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: License_RedeemLicense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).RedeemLicense(ctx, req.(*RedeemLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _License_RevokeUserKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServer).RevokeUserKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: License_RevokeUserKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServer).RevokeUserKeys(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// License_ServiceDesc is the grpc.ServiceDesc for License service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var License_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.License",
	HandlerType: (*LicenseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLicense",
			Handler:    _License_GetLicense_Handler,
		},
		{
			MethodName: "CreateLicense",
			Handler:    _License_CreateLicense_Handler,
		},
		{
			MethodName: "RedeemLicense",
			Handler:    _License_RedeemLicense_Handler,
		},
		{
			MethodName: "RevokeUserKeys",
			Handler:    _License_RevokeUserKeys_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UserLicenseKeyStream",
			Handler:       _License_UserLicenseKeyStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "license.proto",
}
