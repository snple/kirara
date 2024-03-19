// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: class_message.proto

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

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags     string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Config   string `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	Status   int32  `protobuf:"zigzag32,8,opt,name=status,proto3" json:"status,omitempty"`
	Save     int32  `protobuf:"zigzag32,9,opt,name=save,proto3" json:"save,omitempty"`
	Created  int64  `protobuf:"varint,10,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,11,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  int64  `protobuf:"varint,12,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_class_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_class_message_proto_rawDescGZIP(), []int{0}
}

func (x *Class) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Class) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Class) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Class) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Class) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Class) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Class) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Class) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Class) GetSave() int32 {
	if x != nil {
		return x.Save
	}
	return 0
}

func (x *Class) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Class) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Class) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type Attr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ClassId  string `protobuf:"bytes,3,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags     string `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	DataType string `protobuf:"bytes,8,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Value    string `protobuf:"bytes,9,opt,name=value,proto3" json:"value,omitempty"`
	HValue   string `protobuf:"bytes,10,opt,name=h_value,json=hValue,proto3" json:"h_value,omitempty"`
	LValue   string `protobuf:"bytes,11,opt,name=l_value,json=lValue,proto3" json:"l_value,omitempty"`
	TagId    string `protobuf:"bytes,12,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Config   string `protobuf:"bytes,13,opt,name=config,proto3" json:"config,omitempty"`
	Status   int32  `protobuf:"zigzag32,14,opt,name=status,proto3" json:"status,omitempty"`
	Access   int32  `protobuf:"zigzag32,15,opt,name=access,proto3" json:"access,omitempty"`
	Save     int32  `protobuf:"zigzag32,16,opt,name=save,proto3" json:"save,omitempty"`
	Created  int64  `protobuf:"varint,17,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,18,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  int64  `protobuf:"varint,19,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Attr) Reset() {
	*x = Attr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attr) ProtoMessage() {}

func (x *Attr) ProtoReflect() protoreflect.Message {
	mi := &file_class_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attr.ProtoReflect.Descriptor instead.
func (*Attr) Descriptor() ([]byte, []int) {
	return file_class_message_proto_rawDescGZIP(), []int{1}
}

func (x *Attr) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attr) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Attr) GetClassId() string {
	if x != nil {
		return x.ClassId
	}
	return ""
}

func (x *Attr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attr) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Attr) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Attr) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Attr) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Attr) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Attr) GetHValue() string {
	if x != nil {
		return x.HValue
	}
	return ""
}

func (x *Attr) GetLValue() string {
	if x != nil {
		return x.LValue
	}
	return ""
}

func (x *Attr) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *Attr) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Attr) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Attr) GetAccess() int32 {
	if x != nil {
		return x.Access
	}
	return 0
}

func (x *Attr) GetSave() int32 {
	if x != nil {
		return x.Save
	}
	return 0
}

func (x *Attr) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Attr) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Attr) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type AttrValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Updated int64  `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *AttrValue) Reset() {
	*x = AttrValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrValue) ProtoMessage() {}

func (x *AttrValue) ProtoReflect() protoreflect.Message {
	mi := &file_class_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrValue.ProtoReflect.Descriptor instead.
func (*AttrValue) Descriptor() ([]byte, []int) {
	return file_class_message_proto_rawDescGZIP(), []int{2}
}

func (x *AttrValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AttrValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AttrValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type AttrNameValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Updated int64  `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *AttrNameValue) Reset() {
	*x = AttrNameValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrNameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrNameValue) ProtoMessage() {}

func (x *AttrNameValue) ProtoReflect() protoreflect.Message {
	mi := &file_class_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrNameValue.ProtoReflect.Descriptor instead.
func (*AttrNameValue) Descriptor() ([]byte, []int) {
	return file_class_message_proto_rawDescGZIP(), []int{3}
}

func (x *AttrNameValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AttrNameValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttrNameValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AttrNameValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type AttrValueUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	CableId  string `protobuf:"bytes,3,opt,name=cable_id,json=cableId,proto3" json:"cable_id,omitempty"`
	Value    string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Updated  int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *AttrValueUpdated) Reset() {
	*x = AttrValueUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttrValueUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttrValueUpdated) ProtoMessage() {}

func (x *AttrValueUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_class_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttrValueUpdated.ProtoReflect.Descriptor instead.
func (*AttrValueUpdated) Descriptor() ([]byte, []int) {
	return file_class_message_proto_rawDescGZIP(), []int{4}
}

func (x *AttrValueUpdated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AttrValueUpdated) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *AttrValueUpdated) GetCableId() string {
	if x != nil {
		return x.CableId
	}
	return ""
}

func (x *AttrValueUpdated) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AttrValueUpdated) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_class_message_proto protoreflect.FileDescriptor

var file_class_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x96, 0x02, 0x0a, 0x05, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04,
	0x73, 0x61, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x22, 0xc4, 0x03, 0x0a, 0x04, 0x41, 0x74, 0x74, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x73, 0x61, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x4b, 0x0a, 0x09, 0x41, 0x74, 0x74,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x63, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x10,
	0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x69, 0x72,
	0x61, 0x72, 0x61, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_class_message_proto_rawDescOnce sync.Once
	file_class_message_proto_rawDescData = file_class_message_proto_rawDesc
)

func file_class_message_proto_rawDescGZIP() []byte {
	file_class_message_proto_rawDescOnce.Do(func() {
		file_class_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_class_message_proto_rawDescData)
	})
	return file_class_message_proto_rawDescData
}

var file_class_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_class_message_proto_goTypes = []interface{}{
	(*Class)(nil),            // 0: pb.Class
	(*Attr)(nil),             // 1: pb.Attr
	(*AttrValue)(nil),        // 2: pb.AttrValue
	(*AttrNameValue)(nil),    // 3: pb.AttrNameValue
	(*AttrValueUpdated)(nil), // 4: pb.AttrValueUpdated
}
var file_class_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_class_message_proto_init() }
func file_class_message_proto_init() {
	if File_class_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_class_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
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
		file_class_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attr); i {
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
		file_class_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrValue); i {
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
		file_class_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrNameValue); i {
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
		file_class_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttrValueUpdated); i {
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
			RawDescriptor: file_class_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_class_message_proto_goTypes,
		DependencyIndexes: file_class_message_proto_depIdxs,
		MessageInfos:      file_class_message_proto_msgTypes,
	}.Build()
	File_class_message_proto = out.File
	file_class_message_proto_rawDesc = nil
	file_class_message_proto_goTypes = nil
	file_class_message_proto_depIdxs = nil
}