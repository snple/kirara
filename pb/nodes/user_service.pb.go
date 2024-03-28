// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: nodes/user_service.proto

package nodes

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

// user
type UserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Tags string   `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UserListRequest) Reset() {
	*x = UserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRequest) ProtoMessage() {}

func (x *UserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRequest.ProtoReflect.Descriptor instead.
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return file_nodes_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserListRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *UserListRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *UserListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	User  []*pb.User `protobuf:"bytes,3,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_nodes_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserListResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *UserListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *UserListResponse) GetUser() []*pb.User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Type  string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UserPullRequest) Reset() {
	*x = UserPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPullRequest) ProtoMessage() {}

func (x *UserPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPullRequest.ProtoReflect.Descriptor instead.
func (*UserPullRequest) Descriptor() ([]byte, []int) {
	return file_nodes_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserPullRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *UserPullRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *UserPullRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UserPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64      `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	User  []*pb.User `protobuf:"bytes,3,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *UserPullResponse) Reset() {
	*x = UserPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPullResponse) ProtoMessage() {}

func (x *UserPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPullResponse.ProtoReflect.Descriptor instead.
func (*UserPullResponse) Descriptor() ([]byte, []int) {
	return file_nodes_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserPullResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *UserPullResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *UserPullResponse) GetUser() []*pb.User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_nodes_user_service_proto protoreflect.FileDescriptor

var file_nodes_user_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x51, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5c,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xe4, 0x01, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x1c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x25, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c,
	0x12, 0x16, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x69, 0x72, 0x61, 0x72, 0x61, 0x2f, 0x70,
	0x62, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodes_user_service_proto_rawDescOnce sync.Once
	file_nodes_user_service_proto_rawDescData = file_nodes_user_service_proto_rawDesc
)

func file_nodes_user_service_proto_rawDescGZIP() []byte {
	file_nodes_user_service_proto_rawDescOnce.Do(func() {
		file_nodes_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodes_user_service_proto_rawDescData)
	})
	return file_nodes_user_service_proto_rawDescData
}

var file_nodes_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nodes_user_service_proto_goTypes = []interface{}{
	(*UserListRequest)(nil),  // 0: nodes.UserListRequest
	(*UserListResponse)(nil), // 1: nodes.UserListResponse
	(*UserPullRequest)(nil),  // 2: nodes.UserPullRequest
	(*UserPullResponse)(nil), // 3: nodes.UserPullResponse
	(*pb.Page)(nil),          // 4: pb.Page
	(*pb.User)(nil),          // 5: pb.User
	(*pb.Id)(nil),            // 6: pb.Id
	(*pb.Name)(nil),          // 7: pb.Name
}
var file_nodes_user_service_proto_depIdxs = []int32{
	4, // 0: nodes.UserListRequest.page:type_name -> pb.Page
	4, // 1: nodes.UserListResponse.page:type_name -> pb.Page
	5, // 2: nodes.UserListResponse.user:type_name -> pb.User
	5, // 3: nodes.UserPullResponse.user:type_name -> pb.User
	6, // 4: nodes.UserService.View:input_type -> pb.Id
	7, // 5: nodes.UserService.Name:input_type -> pb.Name
	0, // 6: nodes.UserService.List:input_type -> nodes.UserListRequest
	6, // 7: nodes.UserService.ViewWithDeleted:input_type -> pb.Id
	2, // 8: nodes.UserService.Pull:input_type -> nodes.UserPullRequest
	5, // 9: nodes.UserService.View:output_type -> pb.User
	5, // 10: nodes.UserService.Name:output_type -> pb.User
	1, // 11: nodes.UserService.List:output_type -> nodes.UserListResponse
	5, // 12: nodes.UserService.ViewWithDeleted:output_type -> pb.User
	3, // 13: nodes.UserService.Pull:output_type -> nodes.UserPullResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nodes_user_service_proto_init() }
func file_nodes_user_service_proto_init() {
	if File_nodes_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodes_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListRequest); i {
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
		file_nodes_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
		file_nodes_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPullRequest); i {
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
		file_nodes_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPullResponse); i {
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
			RawDescriptor: file_nodes_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodes_user_service_proto_goTypes,
		DependencyIndexes: file_nodes_user_service_proto_depIdxs,
		MessageInfos:      file_nodes_user_service_proto_msgTypes,
	}.Build()
	File_nodes_user_service_proto = out.File
	file_nodes_user_service_proto_rawDesc = nil
	file_nodes_user_service_proto_goTypes = nil
	file_nodes_user_service_proto_depIdxs = nil
}
