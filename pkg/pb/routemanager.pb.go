// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.0
// source: routemanager.proto

package routemanager

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum for business logic errors.
type BusinessError int32

const (
	BusinessError_NO_ERROR             BusinessError = 0
	BusinessError_INVALID_DESTINATION  BusinessError = 1
	BusinessError_ROUTE_NOT_FOUND      BusinessError = 2
	BusinessError_ROUTE_ALREADY_EXISTS BusinessError = 3 // Extend with more business errors as needed.
)

// Enum value maps for BusinessError.
var (
	BusinessError_name = map[int32]string{
		0: "NO_ERROR",
		1: "INVALID_DESTINATION",
		2: "ROUTE_NOT_FOUND",
		3: "ROUTE_ALREADY_EXISTS",
	}
	BusinessError_value = map[string]int32{
		"NO_ERROR":             0,
		"INVALID_DESTINATION":  1,
		"ROUTE_NOT_FOUND":      2,
		"ROUTE_ALREADY_EXISTS": 3,
	}
)

func (x BusinessError) Enum() *BusinessError {
	p := new(BusinessError)
	*p = x
	return p
}

func (x BusinessError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BusinessError) Descriptor() protoreflect.EnumDescriptor {
	return file_routemanager_proto_enumTypes[0].Descriptor()
}

func (BusinessError) Type() protoreflect.EnumType {
	return &file_routemanager_proto_enumTypes[0]
}

func (x BusinessError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BusinessError.Descriptor instead.
func (BusinessError) EnumDescriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{0}
}

// Request and response messages.
type AddRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *AddRouteRequest) Reset() {
	*x = AddRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routemanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRouteRequest) ProtoMessage() {}

func (x *AddRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRouteRequest.ProtoReflect.Descriptor instead.
func (*AddRouteRequest) Descriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{0}
}

func (x *AddRouteRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type AddRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   BusinessError `protobuf:"varint,3,opt,name=error,proto3,enum=routemanager.BusinessError" json:"error,omitempty"`
}

func (x *AddRouteResponse) Reset() {
	*x = AddRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routemanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRouteResponse) ProtoMessage() {}

func (x *AddRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routemanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRouteResponse.ProtoReflect.Descriptor instead.
func (*AddRouteResponse) Descriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{1}
}

func (x *AddRouteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddRouteResponse) GetError() BusinessError {
	if x != nil {
		return x.Error
	}
	return BusinessError_NO_ERROR
}

type RemoveRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *RemoveRouteRequest) Reset() {
	*x = RemoveRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routemanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRouteRequest) ProtoMessage() {}

func (x *RemoveRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRouteRequest.ProtoReflect.Descriptor instead.
func (*RemoveRouteRequest) Descriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveRouteRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type RemoveRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   BusinessError `protobuf:"varint,3,opt,name=error,proto3,enum=routemanager.BusinessError" json:"error,omitempty"`
}

func (x *RemoveRouteResponse) Reset() {
	*x = RemoveRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routemanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRouteResponse) ProtoMessage() {}

func (x *RemoveRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routemanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRouteResponse.ProtoReflect.Descriptor instead.
func (*RemoveRouteResponse) Descriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveRouteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveRouteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RemoveRouteResponse) GetError() BusinessError {
	if x != nil {
		return x.Error
	}
	return BusinessError_NO_ERROR
}

type ListRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoutesRequest) Reset() {
	*x = ListRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routemanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesRequest) ProtoMessage() {}

func (x *ListRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{4}
}

type ListRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Routes  []string      `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
	Error   BusinessError `protobuf:"varint,4,opt,name=error,proto3,enum=routemanager.BusinessError" json:"error,omitempty"`
}

func (x *ListRoutesResponse) Reset() {
	*x = ListRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routemanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesResponse) ProtoMessage() {}

func (x *ListRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routemanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return file_routemanager_proto_rawDescGZIP(), []int{5}
}

func (x *ListRoutesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ListRoutesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListRoutesResponse) GetRoutes() []string {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ListRoutesResponse) GetError() BusinessError {
	if x != nil {
		return x.Error
	}
	return BusinessError_NO_ERROR
}

var File_routemanager_proto protoreflect.FileDescriptor

var file_routemanager_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x22, 0x33, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x36, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x13, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x93, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2a, 0x65, 0x0a, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x45,
	0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52,
	0x4f, 0x55, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x03, 0x32, 0x84, 0x02, 0x0a, 0x0c, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x69, 0x6c, 0x61, 0x6c, 0x63, 0x61, 0x6c, 0x69, 0x73, 0x6b, 0x61, 0x6e, 0x2f, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x2d, 0x74, 0x68, 0x65, 0x2d, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routemanager_proto_rawDescOnce sync.Once
	file_routemanager_proto_rawDescData = file_routemanager_proto_rawDesc
)

func file_routemanager_proto_rawDescGZIP() []byte {
	file_routemanager_proto_rawDescOnce.Do(func() {
		file_routemanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_routemanager_proto_rawDescData)
	})
	return file_routemanager_proto_rawDescData
}

var file_routemanager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_routemanager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_routemanager_proto_goTypes = []interface{}{
	(BusinessError)(0),          // 0: routemanager.BusinessError
	(*AddRouteRequest)(nil),     // 1: routemanager.AddRouteRequest
	(*AddRouteResponse)(nil),    // 2: routemanager.AddRouteResponse
	(*RemoveRouteRequest)(nil),  // 3: routemanager.RemoveRouteRequest
	(*RemoveRouteResponse)(nil), // 4: routemanager.RemoveRouteResponse
	(*ListRoutesRequest)(nil),   // 5: routemanager.ListRoutesRequest
	(*ListRoutesResponse)(nil),  // 6: routemanager.ListRoutesResponse
}
var file_routemanager_proto_depIdxs = []int32{
	0, // 0: routemanager.AddRouteResponse.error:type_name -> routemanager.BusinessError
	0, // 1: routemanager.RemoveRouteResponse.error:type_name -> routemanager.BusinessError
	0, // 2: routemanager.ListRoutesResponse.error:type_name -> routemanager.BusinessError
	1, // 3: routemanager.RouteManager.AddRoute:input_type -> routemanager.AddRouteRequest
	3, // 4: routemanager.RouteManager.RemoveRoute:input_type -> routemanager.RemoveRouteRequest
	5, // 5: routemanager.RouteManager.ListRoutes:input_type -> routemanager.ListRoutesRequest
	2, // 6: routemanager.RouteManager.AddRoute:output_type -> routemanager.AddRouteResponse
	4, // 7: routemanager.RouteManager.RemoveRoute:output_type -> routemanager.RemoveRouteResponse
	6, // 8: routemanager.RouteManager.ListRoutes:output_type -> routemanager.ListRoutesResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_routemanager_proto_init() }
func file_routemanager_proto_init() {
	if File_routemanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routemanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRouteRequest); i {
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
		file_routemanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRouteResponse); i {
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
		file_routemanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRouteRequest); i {
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
		file_routemanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRouteResponse); i {
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
		file_routemanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutesRequest); i {
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
		file_routemanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutesResponse); i {
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
			RawDescriptor: file_routemanager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routemanager_proto_goTypes,
		DependencyIndexes: file_routemanager_proto_depIdxs,
		EnumInfos:         file_routemanager_proto_enumTypes,
		MessageInfos:      file_routemanager_proto_msgTypes,
	}.Build()
	File_routemanager_proto = out.File
	file_routemanager_proto_rawDesc = nil
	file_routemanager_proto_goTypes = nil
	file_routemanager_proto_depIdxs = nil
}