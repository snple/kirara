// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: logic_message.proto

package pb

import (
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

type Fn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags     string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Exec     string `protobuf:"bytes,7,opt,name=exec,proto3" json:"exec,omitempty"`
	Main     string `protobuf:"bytes,8,opt,name=main,proto3" json:"main,omitempty"`
	Config   string `protobuf:"bytes,9,opt,name=config,proto3" json:"config,omitempty"`
	Link     int32  `protobuf:"zigzag32,10,opt,name=link,proto3" json:"link,omitempty"`
	Status   int32  `protobuf:"zigzag32,11,opt,name=status,proto3" json:"status,omitempty"`
	Debug    int32  `protobuf:"zigzag32,12,opt,name=debug,proto3" json:"debug,omitempty"`
	Created  int64  `protobuf:"varint,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,14,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  int64  `protobuf:"varint,15,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Fn) Reset() {
	*x = Fn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fn) ProtoMessage() {}

func (x *Fn) ProtoReflect() protoreflect.Message {
	mi := &file_logic_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fn.ProtoReflect.Descriptor instead.
func (*Fn) Descriptor() ([]byte, []int) {
	return file_logic_message_proto_rawDescGZIP(), []int{0}
}

func (x *Fn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Fn) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Fn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fn) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Fn) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Fn) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Fn) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

func (x *Fn) GetMain() string {
	if x != nil {
		return x.Main
	}
	return ""
}

func (x *Fn) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Fn) GetLink() int32 {
	if x != nil {
		return x.Link
	}
	return 0
}

func (x *Fn) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Fn) GetDebug() int32 {
	if x != nil {
		return x.Debug
	}
	return 0
}

func (x *Fn) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Fn) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Fn) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type Logic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags     string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Exec     string `protobuf:"bytes,7,opt,name=exec,proto3" json:"exec,omitempty"`
	Main     string `protobuf:"bytes,8,opt,name=main,proto3" json:"main,omitempty"`
	Config   string `protobuf:"bytes,9,opt,name=config,proto3" json:"config,omitempty"`
	Status   int32  `protobuf:"zigzag32,10,opt,name=status,proto3" json:"status,omitempty"`
	Created  int64  `protobuf:"varint,11,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,12,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  int64  `protobuf:"varint,13,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Logic) Reset() {
	*x = Logic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logic) ProtoMessage() {}

func (x *Logic) ProtoReflect() protoreflect.Message {
	mi := &file_logic_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logic.ProtoReflect.Descriptor instead.
func (*Logic) Descriptor() ([]byte, []int) {
	return file_logic_message_proto_rawDescGZIP(), []int{1}
}

func (x *Logic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Logic) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Logic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Logic) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Logic) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Logic) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Logic) GetExec() string {
	if x != nil {
		return x.Exec
	}
	return ""
}

func (x *Logic) GetMain() string {
	if x != nil {
		return x.Main
	}
	return ""
}

func (x *Logic) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Logic) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Logic) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Logic) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Logic) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_logic_message_proto protoreflect.FileDescriptor

var file_logic_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xd1, 0x02, 0x0a, 0x02, 0x46, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x65,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xaa, 0x02,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b,
	0x69, 0x72, 0x61, 0x72, 0x61, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_logic_message_proto_rawDescOnce sync.Once
	file_logic_message_proto_rawDescData = file_logic_message_proto_rawDesc
)

func file_logic_message_proto_rawDescGZIP() []byte {
	file_logic_message_proto_rawDescOnce.Do(func() {
		file_logic_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_logic_message_proto_rawDescData)
	})
	return file_logic_message_proto_rawDescData
}

var file_logic_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logic_message_proto_goTypes = []interface{}{
	(*Fn)(nil),    // 0: pb.Fn
	(*Logic)(nil), // 1: pb.Logic
}
var file_logic_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logic_message_proto_init() }
func file_logic_message_proto_init() {
	if File_logic_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logic_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fn); i {
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
		file_logic_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logic); i {
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
			RawDescriptor: file_logic_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logic_message_proto_goTypes,
		DependencyIndexes: file_logic_message_proto_depIdxs,
		MessageInfos:      file_logic_message_proto_msgTypes,
	}.Build()
	File_logic_message_proto = out.File
	file_logic_message_proto_rawDesc = nil
	file_logic_message_proto_goTypes = nil
	file_logic_message_proto_depIdxs = nil
}