// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: license.proto

package protofiles

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LicenseObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Duration         *durationpb.Duration   `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	TimesCompensated uint64                 `protobuf:"varint,3,opt,name=timesCompensated,proto3" json:"timesCompensated,omitempty"`
	HoursCompensated uint64                 `protobuf:"varint,4,opt,name=hoursCompensated,proto3" json:"hoursCompensated,omitempty"`
	Activation       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=activation,proto3,oneof" json:"activation,omitempty"`
	Expiration       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expiration,proto3,oneof" json:"expiration,omitempty"`
	User             *UserObject            `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	Product          *ProductObject         `protobuf:"bytes,8,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *LicenseObject) Reset() {
	*x = LicenseObject{}
	mi := &file_license_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LicenseObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseObject) ProtoMessage() {}

func (x *LicenseObject) ProtoReflect() protoreflect.Message {
	mi := &file_license_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseObject.ProtoReflect.Descriptor instead.
func (*LicenseObject) Descriptor() ([]byte, []int) {
	return file_license_proto_rawDescGZIP(), []int{0}
}

func (x *LicenseObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LicenseObject) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *LicenseObject) GetTimesCompensated() uint64 {
	if x != nil {
		return x.TimesCompensated
	}
	return 0
}

func (x *LicenseObject) GetHoursCompensated() uint64 {
	if x != nil {
		return x.HoursCompensated
	}
	return 0
}

func (x *LicenseObject) GetActivation() *timestamppb.Timestamp {
	if x != nil {
		return x.Activation
	}
	return nil
}

func (x *LicenseObject) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *LicenseObject) GetUser() *UserObject {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *LicenseObject) GetProduct() *ProductObject {
	if x != nil {
		return x.Product
	}
	return nil
}

type LicenseKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseKey string `protobuf:"bytes,1,opt,name=licenseKey,proto3" json:"licenseKey,omitempty"`
}

func (x *LicenseKeyRequest) Reset() {
	*x = LicenseKeyRequest{}
	mi := &file_license_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LicenseKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseKeyRequest) ProtoMessage() {}

func (x *LicenseKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_license_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseKeyRequest.ProtoReflect.Descriptor instead.
func (*LicenseKeyRequest) Descriptor() ([]byte, []int) {
	return file_license_proto_rawDescGZIP(), []int{1}
}

func (x *LicenseKeyRequest) GetLicenseKey() string {
	if x != nil {
		return x.LicenseKey
	}
	return ""
}

type CreateLicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string               `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Duration  *durationpb.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *CreateLicenseRequest) Reset() {
	*x = CreateLicenseRequest{}
	mi := &file_license_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLicenseRequest) ProtoMessage() {}

func (x *CreateLicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_license_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLicenseRequest.ProtoReflect.Descriptor instead.
func (*CreateLicenseRequest) Descriptor() ([]byte, []int) {
	return file_license_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLicenseRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateLicenseRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type RedeemLicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseKey string `protobuf:"bytes,1,opt,name=licenseKey,proto3" json:"licenseKey,omitempty"`
	UserId     uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RedeemLicenseRequest) Reset() {
	*x = RedeemLicenseRequest{}
	mi := &file_license_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedeemLicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedeemLicenseRequest) ProtoMessage() {}

func (x *RedeemLicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_license_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedeemLicenseRequest.ProtoReflect.Descriptor instead.
func (*RedeemLicenseRequest) Descriptor() ([]byte, []int) {
	return file_license_proto_rawDescGZIP(), []int{3}
}

func (x *RedeemLicenseRequest) GetLicenseKey() string {
	if x != nil {
		return x.LicenseKey
	}
	return ""
}

func (x *RedeemLicenseRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_license_proto protoreflect.FileDescriptor

var file_license_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x0d, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x68,
	0x6f, 0x75, 0x72, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x11, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x6b, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x14, 0x52, 0x65,
	0x64, 0x65, 0x65, 0x6d, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x8b, 0x03, 0x0a, 0x07, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4e,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4b, 0x65, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x30, 0x01, 0x12, 0x4c,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4f, 0x0a, 0x0d,
	0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x65,
	0x6d, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x72, 0x73, 0x70,
	0x79, 0x32, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_license_proto_rawDescOnce sync.Once
	file_license_proto_rawDescData = file_license_proto_rawDesc
)

func file_license_proto_rawDescGZIP() []byte {
	file_license_proto_rawDescOnce.Do(func() {
		file_license_proto_rawDescData = protoimpl.X.CompressGZIP(file_license_proto_rawDescData)
	})
	return file_license_proto_rawDescData
}

var file_license_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_license_proto_goTypes = []any{
	(*LicenseObject)(nil),         // 0: protofiles.LicenseObject
	(*LicenseKeyRequest)(nil),     // 1: protofiles.LicenseKeyRequest
	(*CreateLicenseRequest)(nil),  // 2: protofiles.CreateLicenseRequest
	(*RedeemLicenseRequest)(nil),  // 3: protofiles.RedeemLicenseRequest
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*UserObject)(nil),            // 6: protofiles.UserObject
	(*ProductObject)(nil),         // 7: protofiles.ProductObject
	(*UserIdRequest)(nil),         // 8: protofiles.UserIdRequest
	(*StandardResponse)(nil),      // 9: protofiles.StandardResponse
}
var file_license_proto_depIdxs = []int32{
	4,  // 0: protofiles.LicenseObject.duration:type_name -> google.protobuf.Duration
	5,  // 1: protofiles.LicenseObject.activation:type_name -> google.protobuf.Timestamp
	5,  // 2: protofiles.LicenseObject.expiration:type_name -> google.protobuf.Timestamp
	6,  // 3: protofiles.LicenseObject.user:type_name -> protofiles.UserObject
	7,  // 4: protofiles.LicenseObject.product:type_name -> protofiles.ProductObject
	4,  // 5: protofiles.CreateLicenseRequest.duration:type_name -> google.protobuf.Duration
	1,  // 6: protofiles.License.GetLicense:input_type -> protofiles.LicenseKeyRequest
	8,  // 7: protofiles.License.UserLicenseKeyStream:input_type -> protofiles.UserIdRequest
	2,  // 8: protofiles.License.CreateLicense:input_type -> protofiles.CreateLicenseRequest
	3,  // 9: protofiles.License.RedeemLicense:input_type -> protofiles.RedeemLicenseRequest
	8,  // 10: protofiles.License.RevokeUserKeys:input_type -> protofiles.UserIdRequest
	0,  // 11: protofiles.License.GetLicense:output_type -> protofiles.LicenseObject
	0,  // 12: protofiles.License.UserLicenseKeyStream:output_type -> protofiles.LicenseObject
	0,  // 13: protofiles.License.CreateLicense:output_type -> protofiles.LicenseObject
	9,  // 14: protofiles.License.RedeemLicense:output_type -> protofiles.StandardResponse
	9,  // 15: protofiles.License.RevokeUserKeys:output_type -> protofiles.StandardResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_license_proto_init() }
func file_license_proto_init() {
	if File_license_proto != nil {
		return
	}
	file_product_proto_init()
	file_user_proto_init()
	file_globals_proto_init()
	file_license_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_license_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_license_proto_goTypes,
		DependencyIndexes: file_license_proto_depIdxs,
		MessageInfos:      file_license_proto_msgTypes,
	}.Build()
	File_license_proto = out.File
	file_license_proto_rawDesc = nil
	file_license_proto_goTypes = nil
	file_license_proto_depIdxs = nil
}
