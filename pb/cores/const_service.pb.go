// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: cores/const_service.proto

package cores

import (
	pb "github.com/snple/kirara/pb"
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

type ConstListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	DeviceId string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Tags     string   `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ConstListRequest) Reset() {
	*x = ConstListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstListRequest) ProtoMessage() {}

func (x *ConstListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstListRequest.ProtoReflect.Descriptor instead.
func (*ConstListRequest) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{0}
}

func (x *ConstListRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ConstListRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConstListRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ConstListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ConstListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Const []*pb.Const `protobuf:"bytes,3,rep,name=const,proto3" json:"const,omitempty"`
}

func (x *ConstListResponse) Reset() {
	*x = ConstListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstListResponse) ProtoMessage() {}

func (x *ConstListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstListResponse.ProtoReflect.Descriptor instead.
func (*ConstListResponse) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConstListResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ConstListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ConstListResponse) GetConst() []*pb.Const {
	if x != nil {
		return x.Const
	}
	return nil
}

type ConstNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConstNameRequest) Reset() {
	*x = ConstNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstNameRequest) ProtoMessage() {}

func (x *ConstNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstNameRequest.ProtoReflect.Descriptor instead.
func (*ConstNameRequest) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConstNameRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConstNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConstCloneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *ConstCloneRequest) Reset() {
	*x = ConstCloneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstCloneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstCloneRequest) ProtoMessage() {}

func (x *ConstCloneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstCloneRequest.ProtoReflect.Descriptor instead.
func (*ConstCloneRequest) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConstCloneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConstCloneRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type ConstGetValueByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConstGetValueByNameRequest) Reset() {
	*x = ConstGetValueByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstGetValueByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstGetValueByNameRequest) ProtoMessage() {}

func (x *ConstGetValueByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstGetValueByNameRequest.ProtoReflect.Descriptor instead.
func (*ConstGetValueByNameRequest) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{4}
}

func (x *ConstGetValueByNameRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConstGetValueByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConstNameValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value    string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Updated  int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *ConstNameValue) Reset() {
	*x = ConstNameValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstNameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstNameValue) ProtoMessage() {}

func (x *ConstNameValue) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstNameValue.ProtoReflect.Descriptor instead.
func (*ConstNameValue) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{5}
}

func (x *ConstNameValue) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConstNameValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConstNameValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConstNameValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConstNameValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type ConstPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After    int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit    uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ConstPullRequest) Reset() {
	*x = ConstPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstPullRequest) ProtoMessage() {}

func (x *ConstPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstPullRequest.ProtoReflect.Descriptor instead.
func (*ConstPullRequest) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{6}
}

func (x *ConstPullRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *ConstPullRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ConstPullRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConstPullRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ConstPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64       `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Const []*pb.Const `protobuf:"bytes,3,rep,name=const,proto3" json:"const,omitempty"`
}

func (x *ConstPullResponse) Reset() {
	*x = ConstPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_const_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstPullResponse) ProtoMessage() {}

func (x *ConstPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_const_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstPullResponse.ProtoReflect.Descriptor instead.
func (*ConstPullResponse) Descriptor() ([]byte, []int) {
	return file_cores_const_service_proto_rawDescGZIP(), []int{7}
}

func (x *ConstPullResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *ConstPullResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ConstPullResponse) GetConst() []*pb.Const {
	if x != nil {
		return x.Const
	}
	return nil
}

var File_cores_const_service_proto protoreflect.FileDescriptor

var file_cores_const_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x1a, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x22,
	0x43, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x43, 0x6c, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x6f, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x32, 0x94, 0x06, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x20, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x1b, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2c,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x08,
	0x4e, 0x61, 0x6d, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x05,
	0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x17, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x08, 0x53, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x15,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x69, 0x72, 0x61, 0x72, 0x61, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cores_const_service_proto_rawDescOnce sync.Once
	file_cores_const_service_proto_rawDescData = file_cores_const_service_proto_rawDesc
)

func file_cores_const_service_proto_rawDescGZIP() []byte {
	file_cores_const_service_proto_rawDescOnce.Do(func() {
		file_cores_const_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cores_const_service_proto_rawDescData)
	})
	return file_cores_const_service_proto_rawDescData
}

var file_cores_const_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cores_const_service_proto_goTypes = []interface{}{
	(*ConstListRequest)(nil),           // 0: cores.ConstListRequest
	(*ConstListResponse)(nil),          // 1: cores.ConstListResponse
	(*ConstNameRequest)(nil),           // 2: cores.ConstNameRequest
	(*ConstCloneRequest)(nil),          // 3: cores.ConstCloneRequest
	(*ConstGetValueByNameRequest)(nil), // 4: cores.ConstGetValueByNameRequest
	(*ConstNameValue)(nil),             // 5: cores.ConstNameValue
	(*ConstPullRequest)(nil),           // 6: cores.ConstPullRequest
	(*ConstPullResponse)(nil),          // 7: cores.ConstPullResponse
	(*pb.Page)(nil),                    // 8: pb.Page
	(*pb.Const)(nil),                   // 9: pb.Const
	(*pb.Id)(nil),                      // 10: pb.Id
	(*pb.Name)(nil),                    // 11: pb.Name
	(*pb.ConstValue)(nil),              // 12: pb.ConstValue
	(*pb.MyBool)(nil),                  // 13: pb.MyBool
}
var file_cores_const_service_proto_depIdxs = []int32{
	8,  // 0: cores.ConstListRequest.page:type_name -> pb.Page
	8,  // 1: cores.ConstListResponse.page:type_name -> pb.Page
	9,  // 2: cores.ConstListResponse.const:type_name -> pb.Const
	9,  // 3: cores.ConstPullResponse.const:type_name -> pb.Const
	9,  // 4: cores.ConstService.Create:input_type -> pb.Const
	9,  // 5: cores.ConstService.Update:input_type -> pb.Const
	10, // 6: cores.ConstService.View:input_type -> pb.Id
	2,  // 7: cores.ConstService.Name:input_type -> cores.ConstNameRequest
	11, // 8: cores.ConstService.NameFull:input_type -> pb.Name
	10, // 9: cores.ConstService.Delete:input_type -> pb.Id
	0,  // 10: cores.ConstService.List:input_type -> cores.ConstListRequest
	3,  // 11: cores.ConstService.Clone:input_type -> cores.ConstCloneRequest
	10, // 12: cores.ConstService.ViewWithDeleted:input_type -> pb.Id
	6,  // 13: cores.ConstService.Pull:input_type -> cores.ConstPullRequest
	9,  // 14: cores.ConstService.Sync:input_type -> pb.Const
	10, // 15: cores.ConstService.GetValue:input_type -> pb.Id
	12, // 16: cores.ConstService.SetValue:input_type -> pb.ConstValue
	12, // 17: cores.ConstService.SetValueForce:input_type -> pb.ConstValue
	4,  // 18: cores.ConstService.GetValueByName:input_type -> cores.ConstGetValueByNameRequest
	5,  // 19: cores.ConstService.SetValueByName:input_type -> cores.ConstNameValue
	5,  // 20: cores.ConstService.SetValueByNameForce:input_type -> cores.ConstNameValue
	9,  // 21: cores.ConstService.Create:output_type -> pb.Const
	9,  // 22: cores.ConstService.Update:output_type -> pb.Const
	9,  // 23: cores.ConstService.View:output_type -> pb.Const
	9,  // 24: cores.ConstService.Name:output_type -> pb.Const
	9,  // 25: cores.ConstService.NameFull:output_type -> pb.Const
	13, // 26: cores.ConstService.Delete:output_type -> pb.MyBool
	1,  // 27: cores.ConstService.List:output_type -> cores.ConstListResponse
	13, // 28: cores.ConstService.Clone:output_type -> pb.MyBool
	9,  // 29: cores.ConstService.ViewWithDeleted:output_type -> pb.Const
	7,  // 30: cores.ConstService.Pull:output_type -> cores.ConstPullResponse
	13, // 31: cores.ConstService.Sync:output_type -> pb.MyBool
	12, // 32: cores.ConstService.GetValue:output_type -> pb.ConstValue
	13, // 33: cores.ConstService.SetValue:output_type -> pb.MyBool
	13, // 34: cores.ConstService.SetValueForce:output_type -> pb.MyBool
	5,  // 35: cores.ConstService.GetValueByName:output_type -> cores.ConstNameValue
	13, // 36: cores.ConstService.SetValueByName:output_type -> pb.MyBool
	13, // 37: cores.ConstService.SetValueByNameForce:output_type -> pb.MyBool
	21, // [21:38] is the sub-list for method output_type
	4,  // [4:21] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cores_const_service_proto_init() }
func file_cores_const_service_proto_init() {
	if File_cores_const_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cores_const_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstListRequest); i {
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
		file_cores_const_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstListResponse); i {
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
		file_cores_const_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstNameRequest); i {
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
		file_cores_const_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstCloneRequest); i {
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
		file_cores_const_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstGetValueByNameRequest); i {
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
		file_cores_const_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstNameValue); i {
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
		file_cores_const_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstPullRequest); i {
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
		file_cores_const_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstPullResponse); i {
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
			RawDescriptor: file_cores_const_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cores_const_service_proto_goTypes,
		DependencyIndexes: file_cores_const_service_proto_depIdxs,
		MessageInfos:      file_cores_const_service_proto_msgTypes,
	}.Build()
	File_cores_const_service_proto = out.File
	file_cores_const_service_proto_rawDesc = nil
	file_cores_const_service_proto_goTypes = nil
	file_cores_const_service_proto_depIdxs = nil
}