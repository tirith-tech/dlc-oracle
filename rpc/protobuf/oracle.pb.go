// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: rpc/protobuf/oracle.proto

package protobuf

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PubKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (x *PubKeyResponse) Reset() {
	*x = PubKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeyResponse) ProtoMessage() {}

func (x *PubKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeyResponse.ProtoReflect.Descriptor instead.
func (*PubKeyResponse) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{0}
}

func (x *PubKeyResponse) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

type DataSourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Id           uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	CurrentValue uint64 `protobuf:"varint,4,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	ValueError   string `protobuf:"bytes,5,opt,name=value_error,json=valueError,proto3" json:"value_error,omitempty"`
}

func (x *DataSourcesResponse) Reset() {
	*x = DataSourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourcesResponse) ProtoMessage() {}

func (x *DataSourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourcesResponse.ProtoReflect.Descriptor instead.
func (*DataSourcesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{1}
}

func (x *DataSourcesResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSourcesResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DataSourcesResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataSourcesResponse) GetCurrentValue() uint64 {
	if x != nil {
		return x.CurrentValue
	}
	return 0
}

func (x *DataSourcesResponse) GetValueError() string {
	if x != nil {
		return x.ValueError
	}
	return ""
}

type RPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *RPointRequest) Reset() {
	*x = RPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPointRequest) ProtoMessage() {}

func (x *RPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPointRequest.ProtoReflect.Descriptor instead.
func (*RPointRequest) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{2}
}

func (x *RPointRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RPointRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type RPointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RPoint string `protobuf:"bytes,1,opt,name=r_point,json=rPoint,proto3" json:"r_point,omitempty"`
}

func (x *RPointResponse) Reset() {
	*x = RPointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPointResponse) ProtoMessage() {}

func (x *RPointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPointResponse.ProtoReflect.Descriptor instead.
func (*RPointResponse) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{3}
}

func (x *RPointResponse) GetRPoint() string {
	if x != nil {
		return x.RPoint
	}
	return ""
}

type PublicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RPoint string `protobuf:"bytes,1,opt,name=r_point,json=rPoint,proto3" json:"r_point,omitempty"`
}

func (x *PublicationRequest) Reset() {
	*x = PublicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicationRequest) ProtoMessage() {}

func (x *PublicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicationRequest.ProtoReflect.Descriptor instead.
func (*PublicationRequest) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{4}
}

func (x *PublicationRequest) GetRPoint() string {
	if x != nil {
		return x.RPoint
	}
	return ""
}

type PublicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PublicationResponse) Reset() {
	*x = PublicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicationResponse) ProtoMessage() {}

func (x *PublicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicationResponse.ProtoReflect.Descriptor instead.
func (*PublicationResponse) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{5}
}

func (x *PublicationResponse) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PublicationResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PublicationResponse) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PublicationResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PublicationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  string `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote string `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *PublicationsRequest) Reset() {
	*x = PublicationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protobuf_oracle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicationsRequest) ProtoMessage() {}

func (x *PublicationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protobuf_oracle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicationsRequest.ProtoReflect.Descriptor instead.
func (*PublicationsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_protobuf_oracle_proto_rawDescGZIP(), []int{6}
}

func (x *PublicationsRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *PublicationsRequest) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

var File_rpc_protobuf_oracle_proto protoreflect.FileDescriptor

var file_rpc_protobuf_oracle_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x28, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3d,
	0x0a, 0x0d, 0x52, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x29, 0x0a,
	0x0e, 0x52, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x7b, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x32, 0xe6, 0x02, 0x0a, 0x0d, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x52,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x52,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x52, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0e,
	0x5a, 0x0c, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_protobuf_oracle_proto_rawDescOnce sync.Once
	file_rpc_protobuf_oracle_proto_rawDescData = file_rpc_protobuf_oracle_proto_rawDesc
)

func file_rpc_protobuf_oracle_proto_rawDescGZIP() []byte {
	file_rpc_protobuf_oracle_proto_rawDescOnce.Do(func() {
		file_rpc_protobuf_oracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_protobuf_oracle_proto_rawDescData)
	})
	return file_rpc_protobuf_oracle_proto_rawDescData
}

var file_rpc_protobuf_oracle_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_protobuf_oracle_proto_goTypes = []interface{}{
	(*PubKeyResponse)(nil),      // 0: oracle.PubKeyResponse
	(*DataSourcesResponse)(nil), // 1: oracle.DataSourcesResponse
	(*RPointRequest)(nil),       // 2: oracle.RPointRequest
	(*RPointResponse)(nil),      // 3: oracle.RPointResponse
	(*PublicationRequest)(nil),  // 4: oracle.PublicationRequest
	(*PublicationResponse)(nil), // 5: oracle.PublicationResponse
	(*PublicationsRequest)(nil), // 6: oracle.PublicationsRequest
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_rpc_protobuf_oracle_proto_depIdxs = []int32{
	7, // 0: oracle.OracleService.PubKey:input_type -> google.protobuf.Empty
	7, // 1: oracle.OracleService.DataSources:input_type -> google.protobuf.Empty
	2, // 2: oracle.OracleService.RPoint:input_type -> oracle.RPointRequest
	4, // 3: oracle.OracleService.Publication:input_type -> oracle.PublicationRequest
	6, // 4: oracle.OracleService.Publications:input_type -> oracle.PublicationsRequest
	0, // 5: oracle.OracleService.PubKey:output_type -> oracle.PubKeyResponse
	1, // 6: oracle.OracleService.DataSources:output_type -> oracle.DataSourcesResponse
	3, // 7: oracle.OracleService.RPoint:output_type -> oracle.RPointResponse
	5, // 8: oracle.OracleService.Publication:output_type -> oracle.PublicationResponse
	5, // 9: oracle.OracleService.Publications:output_type -> oracle.PublicationResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_protobuf_oracle_proto_init() }
func file_rpc_protobuf_oracle_proto_init() {
	if File_rpc_protobuf_oracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_protobuf_oracle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_protobuf_oracle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourcesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_protobuf_oracle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_protobuf_oracle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPointResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_protobuf_oracle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_protobuf_oracle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_protobuf_oracle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_protobuf_oracle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_protobuf_oracle_proto_goTypes,
		DependencyIndexes: file_rpc_protobuf_oracle_proto_depIdxs,
		MessageInfos:      file_rpc_protobuf_oracle_proto_msgTypes,
	}.Build()
	File_rpc_protobuf_oracle_proto = out.File
	file_rpc_protobuf_oracle_proto_rawDesc = nil
	file_rpc_protobuf_oracle_proto_goTypes = nil
	file_rpc_protobuf_oracle_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OracleServiceClient is the client API for OracleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OracleServiceClient interface {
	PubKey(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PubKeyResponse, error)
	DataSources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (OracleService_DataSourcesClient, error)
	RPoint(ctx context.Context, in *RPointRequest, opts ...grpc.CallOption) (*RPointResponse, error)
	Publication(ctx context.Context, in *PublicationRequest, opts ...grpc.CallOption) (*PublicationResponse, error)
	Publications(ctx context.Context, in *PublicationsRequest, opts ...grpc.CallOption) (OracleService_PublicationsClient, error)
}

type oracleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOracleServiceClient(cc grpc.ClientConnInterface) OracleServiceClient {
	return &oracleServiceClient{cc}
}

func (c *oracleServiceClient) PubKey(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PubKeyResponse, error) {
	out := new(PubKeyResponse)
	err := c.cc.Invoke(ctx, "/oracle.OracleService/PubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oracleServiceClient) DataSources(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (OracleService_DataSourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OracleService_serviceDesc.Streams[0], "/oracle.OracleService/DataSources", opts...)
	if err != nil {
		return nil, err
	}
	x := &oracleServiceDataSourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OracleService_DataSourcesClient interface {
	Recv() (*DataSourcesResponse, error)
	grpc.ClientStream
}

type oracleServiceDataSourcesClient struct {
	grpc.ClientStream
}

func (x *oracleServiceDataSourcesClient) Recv() (*DataSourcesResponse, error) {
	m := new(DataSourcesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oracleServiceClient) RPoint(ctx context.Context, in *RPointRequest, opts ...grpc.CallOption) (*RPointResponse, error) {
	out := new(RPointResponse)
	err := c.cc.Invoke(ctx, "/oracle.OracleService/RPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oracleServiceClient) Publication(ctx context.Context, in *PublicationRequest, opts ...grpc.CallOption) (*PublicationResponse, error) {
	out := new(PublicationResponse)
	err := c.cc.Invoke(ctx, "/oracle.OracleService/Publication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oracleServiceClient) Publications(ctx context.Context, in *PublicationsRequest, opts ...grpc.CallOption) (OracleService_PublicationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OracleService_serviceDesc.Streams[1], "/oracle.OracleService/Publications", opts...)
	if err != nil {
		return nil, err
	}
	x := &oracleServicePublicationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OracleService_PublicationsClient interface {
	Recv() (*PublicationResponse, error)
	grpc.ClientStream
}

type oracleServicePublicationsClient struct {
	grpc.ClientStream
}

func (x *oracleServicePublicationsClient) Recv() (*PublicationResponse, error) {
	m := new(PublicationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OracleServiceServer is the server API for OracleService service.
type OracleServiceServer interface {
	PubKey(context.Context, *empty.Empty) (*PubKeyResponse, error)
	DataSources(*empty.Empty, OracleService_DataSourcesServer) error
	RPoint(context.Context, *RPointRequest) (*RPointResponse, error)
	Publication(context.Context, *PublicationRequest) (*PublicationResponse, error)
	Publications(*PublicationsRequest, OracleService_PublicationsServer) error
}

// UnimplementedOracleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOracleServiceServer struct {
}

func (*UnimplementedOracleServiceServer) PubKey(context.Context, *empty.Empty) (*PubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PubKey not implemented")
}
func (*UnimplementedOracleServiceServer) DataSources(*empty.Empty, OracleService_DataSourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DataSources not implemented")
}
func (*UnimplementedOracleServiceServer) RPoint(context.Context, *RPointRequest) (*RPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPoint not implemented")
}
func (*UnimplementedOracleServiceServer) Publication(context.Context, *PublicationRequest) (*PublicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publication not implemented")
}
func (*UnimplementedOracleServiceServer) Publications(*PublicationsRequest, OracleService_PublicationsServer) error {
	return status.Errorf(codes.Unimplemented, "method Publications not implemented")
}

func RegisterOracleServiceServer(s *grpc.Server, srv OracleServiceServer) {
	s.RegisterService(&_OracleService_serviceDesc, srv)
}

func _OracleService_PubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServiceServer).PubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oracle.OracleService/PubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServiceServer).PubKey(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OracleService_DataSources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OracleServiceServer).DataSources(m, &oracleServiceDataSourcesServer{stream})
}

type OracleService_DataSourcesServer interface {
	Send(*DataSourcesResponse) error
	grpc.ServerStream
}

type oracleServiceDataSourcesServer struct {
	grpc.ServerStream
}

func (x *oracleServiceDataSourcesServer) Send(m *DataSourcesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OracleService_RPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServiceServer).RPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oracle.OracleService/RPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServiceServer).RPoint(ctx, req.(*RPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OracleService_Publication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServiceServer).Publication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oracle.OracleService/Publication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServiceServer).Publication(ctx, req.(*PublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OracleService_Publications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PublicationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OracleServiceServer).Publications(m, &oracleServicePublicationsServer{stream})
}

type OracleService_PublicationsServer interface {
	Send(*PublicationResponse) error
	grpc.ServerStream
}

type oracleServicePublicationsServer struct {
	grpc.ServerStream
}

func (x *oracleServicePublicationsServer) Send(m *PublicationResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _OracleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oracle.OracleService",
	HandlerType: (*OracleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PubKey",
			Handler:    _OracleService_PubKey_Handler,
		},
		{
			MethodName: "RPoint",
			Handler:    _OracleService_RPoint_Handler,
		},
		{
			MethodName: "Publication",
			Handler:    _OracleService_Publication_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DataSources",
			Handler:       _OracleService_DataSources_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Publications",
			Handler:       _OracleService_Publications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/protobuf/oracle.proto",
}
