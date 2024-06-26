// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: slots/logic_service.proto

package slots

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

type FnListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// string device_id = 2;
	Tags string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Exec string `protobuf:"bytes,5,opt,name=exec,proto3" json:"exec,omitempty"`
}

func (x *FnListRequest) Reset() {
	*x = FnListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FnListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FnListRequest) ProtoMessage() {}

func (x *FnListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FnListRequest.ProtoReflect.Descriptor instead.
func (*FnListRequest) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{0}
}

func (x *FnListRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FnListRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *FnListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FnListRequest) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

type FnListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Fn    []*pb.Fn `protobuf:"bytes,3,rep,name=fn,proto3" json:"fn,omitempty"`
}

func (x *FnListResponse) Reset() {
	*x = FnListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FnListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FnListResponse) ProtoMessage() {}

func (x *FnListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FnListResponse.ProtoReflect.Descriptor instead.
func (*FnListResponse) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{1}
}

func (x *FnListResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FnListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FnListResponse) GetFn() []*pb.Fn {
	if x != nil {
		return x.Fn
	}
	return nil
}

type FnLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"zigzag32,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FnLinkRequest) Reset() {
	*x = FnLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FnLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FnLinkRequest) ProtoMessage() {}

func (x *FnLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FnLinkRequest.ProtoReflect.Descriptor instead.
func (*FnLinkRequest) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{2}
}

func (x *FnLinkRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FnLinkRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type FnPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// string device_id = 3;
	Exec string `protobuf:"bytes,4,opt,name=exec,proto3" json:"exec,omitempty"`
}

func (x *FnPullRequest) Reset() {
	*x = FnPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FnPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FnPullRequest) ProtoMessage() {}

func (x *FnPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FnPullRequest.ProtoReflect.Descriptor instead.
func (*FnPullRequest) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{3}
}

func (x *FnPullRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *FnPullRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FnPullRequest) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

type FnPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64    `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Fn    []*pb.Fn `protobuf:"bytes,3,rep,name=fn,proto3" json:"fn,omitempty"`
}

func (x *FnPullResponse) Reset() {
	*x = FnPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FnPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FnPullResponse) ProtoMessage() {}

func (x *FnPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FnPullResponse.ProtoReflect.Descriptor instead.
func (*FnPullResponse) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{4}
}

func (x *FnPullResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *FnPullResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FnPullResponse) GetFn() []*pb.Fn {
	if x != nil {
		return x.Fn
	}
	return nil
}

type LogicListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// string device_id = 2;
	Tags string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Exec string `protobuf:"bytes,5,opt,name=exec,proto3" json:"exec,omitempty"`
}

func (x *LogicListRequest) Reset() {
	*x = LogicListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicListRequest) ProtoMessage() {}

func (x *LogicListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicListRequest.ProtoReflect.Descriptor instead.
func (*LogicListRequest) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{5}
}

func (x *LogicListRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *LogicListRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *LogicListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LogicListRequest) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

type LogicListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Logic []*pb.Logic `protobuf:"bytes,3,rep,name=logic,proto3" json:"logic,omitempty"`
}

func (x *LogicListResponse) Reset() {
	*x = LogicListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicListResponse) ProtoMessage() {}

func (x *LogicListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicListResponse.ProtoReflect.Descriptor instead.
func (*LogicListResponse) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{6}
}

func (x *LogicListResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *LogicListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *LogicListResponse) GetLogic() []*pb.Logic {
	if x != nil {
		return x.Logic
	}
	return nil
}

type LogicPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// string device_id = 3;
	Exec string `protobuf:"bytes,4,opt,name=exec,proto3" json:"exec,omitempty"`
}

func (x *LogicPullRequest) Reset() {
	*x = LogicPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicPullRequest) ProtoMessage() {}

func (x *LogicPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicPullRequest.ProtoReflect.Descriptor instead.
func (*LogicPullRequest) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{7}
}

func (x *LogicPullRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *LogicPullRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LogicPullRequest) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

type LogicPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64       `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Logic []*pb.Logic `protobuf:"bytes,3,rep,name=logic,proto3" json:"logic,omitempty"`
}

func (x *LogicPullResponse) Reset() {
	*x = LogicPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slots_logic_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicPullResponse) ProtoMessage() {}

func (x *LogicPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slots_logic_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicPullResponse.ProtoReflect.Descriptor instead.
func (*LogicPullResponse) Descriptor() ([]byte, []int) {
	return file_slots_logic_service_proto_rawDescGZIP(), []int{8}
}

func (x *LogicPullResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *LogicPullResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LogicPullResponse) GetLogic() []*pb.Logic {
	if x != nil {
		return x.Logic
	}
	return nil
}

var File_slots_logic_service_proto protoreflect.FileDescriptor

var file_slots_logic_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6c, 0x6f,
	0x74, 0x73, 0x1a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69,
	0x0a, 0x0d, 0x46, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x22, 0x5c, 0x0a, 0x0e, 0x46, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x02, 0x66, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x6e, 0x52, 0x02, 0x66, 0x6e, 0x22, 0x37, 0x0a, 0x0d, 0x46, 0x6e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x4f, 0x0a, 0x0d, 0x46, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x65,
	0x63, 0x22, 0x54, 0x0a, 0x0e, 0x46, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x02, 0x66, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x6e, 0x52, 0x02, 0x66, 0x6e, 0x22, 0x6c, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x63,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x22, 0x68, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x22,
	0x52, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x78, 0x65, 0x63, 0x22, 0x60, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x50, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x32, 0xf6, 0x02, 0x0a, 0x09, 0x46, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x06, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x6e, 0x1a, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6e, 0x22, 0x00, 0x12,
	0x1a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x6e, 0x1a, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6e, 0x22, 0x00, 0x12, 0x18, 0x0a, 0x04, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x6e, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6e, 0x22,
	0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x6c, 0x6f, 0x74,
	0x73, 0x2e, 0x46, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x46, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x14, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x46, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a,
	0x06, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6e, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x50, 0x75, 0x6c,
	0x6c, 0x12, 0x14, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x46, 0x6e, 0x50, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e,
	0x46, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x1c, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6e,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x32, 0xf1,
	0x02, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x22,
	0x00, 0x12, 0x20, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x63, 0x22, 0x00, 0x12, 0x1b, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x22, 0x00,
	0x12, 0x1d, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x22, 0x00, 0x12,
	0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0f,
	0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x63, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x73,
	0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c,
	0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x69, 0x72, 0x61, 0x72, 0x61, 0x2f, 0x70, 0x62,
	0x2f, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x3b, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slots_logic_service_proto_rawDescOnce sync.Once
	file_slots_logic_service_proto_rawDescData = file_slots_logic_service_proto_rawDesc
)

func file_slots_logic_service_proto_rawDescGZIP() []byte {
	file_slots_logic_service_proto_rawDescOnce.Do(func() {
		file_slots_logic_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_slots_logic_service_proto_rawDescData)
	})
	return file_slots_logic_service_proto_rawDescData
}

var file_slots_logic_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_slots_logic_service_proto_goTypes = []interface{}{
	(*FnListRequest)(nil),     // 0: slots.FnListRequest
	(*FnListResponse)(nil),    // 1: slots.FnListResponse
	(*FnLinkRequest)(nil),     // 2: slots.FnLinkRequest
	(*FnPullRequest)(nil),     // 3: slots.FnPullRequest
	(*FnPullResponse)(nil),    // 4: slots.FnPullResponse
	(*LogicListRequest)(nil),  // 5: slots.LogicListRequest
	(*LogicListResponse)(nil), // 6: slots.LogicListResponse
	(*LogicPullRequest)(nil),  // 7: slots.LogicPullRequest
	(*LogicPullResponse)(nil), // 8: slots.LogicPullResponse
	(*pb.Page)(nil),           // 9: pb.Page
	(*pb.Fn)(nil),             // 10: pb.Fn
	(*pb.Logic)(nil),          // 11: pb.Logic
	(*pb.Id)(nil),             // 12: pb.Id
	(*pb.Name)(nil),           // 13: pb.Name
	(*pb.MyBool)(nil),         // 14: pb.MyBool
}
var file_slots_logic_service_proto_depIdxs = []int32{
	9,  // 0: slots.FnListRequest.page:type_name -> pb.Page
	9,  // 1: slots.FnListResponse.page:type_name -> pb.Page
	10, // 2: slots.FnListResponse.fn:type_name -> pb.Fn
	10, // 3: slots.FnPullResponse.fn:type_name -> pb.Fn
	9,  // 4: slots.LogicListRequest.page:type_name -> pb.Page
	9,  // 5: slots.LogicListResponse.page:type_name -> pb.Page
	11, // 6: slots.LogicListResponse.logic:type_name -> pb.Logic
	11, // 7: slots.LogicPullResponse.logic:type_name -> pb.Logic
	10, // 8: slots.FnService.Create:input_type -> pb.Fn
	10, // 9: slots.FnService.Update:input_type -> pb.Fn
	12, // 10: slots.FnService.View:input_type -> pb.Id
	13, // 11: slots.FnService.Name:input_type -> pb.Name
	12, // 12: slots.FnService.Delete:input_type -> pb.Id
	0,  // 13: slots.FnService.List:input_type -> slots.FnListRequest
	2,  // 14: slots.FnService.Link:input_type -> slots.FnLinkRequest
	12, // 15: slots.FnService.ViewWithDeleted:input_type -> pb.Id
	3,  // 16: slots.FnService.Pull:input_type -> slots.FnPullRequest
	10, // 17: slots.FnService.Sync:input_type -> pb.Fn
	11, // 18: slots.LogicService.Create:input_type -> pb.Logic
	11, // 19: slots.LogicService.Update:input_type -> pb.Logic
	12, // 20: slots.LogicService.View:input_type -> pb.Id
	13, // 21: slots.LogicService.Name:input_type -> pb.Name
	12, // 22: slots.LogicService.Delete:input_type -> pb.Id
	5,  // 23: slots.LogicService.List:input_type -> slots.LogicListRequest
	12, // 24: slots.LogicService.ViewWithDeleted:input_type -> pb.Id
	7,  // 25: slots.LogicService.Pull:input_type -> slots.LogicPullRequest
	11, // 26: slots.LogicService.Sync:input_type -> pb.Logic
	10, // 27: slots.FnService.Create:output_type -> pb.Fn
	10, // 28: slots.FnService.Update:output_type -> pb.Fn
	10, // 29: slots.FnService.View:output_type -> pb.Fn
	10, // 30: slots.FnService.Name:output_type -> pb.Fn
	14, // 31: slots.FnService.Delete:output_type -> pb.MyBool
	1,  // 32: slots.FnService.List:output_type -> slots.FnListResponse
	14, // 33: slots.FnService.Link:output_type -> pb.MyBool
	10, // 34: slots.FnService.ViewWithDeleted:output_type -> pb.Fn
	4,  // 35: slots.FnService.Pull:output_type -> slots.FnPullResponse
	14, // 36: slots.FnService.Sync:output_type -> pb.MyBool
	11, // 37: slots.LogicService.Create:output_type -> pb.Logic
	11, // 38: slots.LogicService.Update:output_type -> pb.Logic
	11, // 39: slots.LogicService.View:output_type -> pb.Logic
	11, // 40: slots.LogicService.Name:output_type -> pb.Logic
	14, // 41: slots.LogicService.Delete:output_type -> pb.MyBool
	6,  // 42: slots.LogicService.List:output_type -> slots.LogicListResponse
	11, // 43: slots.LogicService.ViewWithDeleted:output_type -> pb.Logic
	8,  // 44: slots.LogicService.Pull:output_type -> slots.LogicPullResponse
	14, // 45: slots.LogicService.Sync:output_type -> pb.MyBool
	27, // [27:46] is the sub-list for method output_type
	8,  // [8:27] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_slots_logic_service_proto_init() }
func file_slots_logic_service_proto_init() {
	if File_slots_logic_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slots_logic_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FnListRequest); i {
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
		file_slots_logic_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FnListResponse); i {
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
		file_slots_logic_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FnLinkRequest); i {
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
		file_slots_logic_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FnPullRequest); i {
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
		file_slots_logic_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FnPullResponse); i {
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
		file_slots_logic_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicListRequest); i {
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
		file_slots_logic_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicListResponse); i {
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
		file_slots_logic_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicPullRequest); i {
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
		file_slots_logic_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicPullResponse); i {
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
			RawDescriptor: file_slots_logic_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_slots_logic_service_proto_goTypes,
		DependencyIndexes: file_slots_logic_service_proto_depIdxs,
		MessageInfos:      file_slots_logic_service_proto_msgTypes,
	}.Build()
	File_slots_logic_service_proto = out.File
	file_slots_logic_service_proto_rawDesc = nil
	file_slots_logic_service_proto_goTypes = nil
	file_slots_logic_service_proto_depIdxs = nil
}
