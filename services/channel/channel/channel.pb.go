// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: channel.proto

package channel

import (
	context "context"
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

type ChannelModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ChannelId          string `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	ChannelSecretId    string `protobuf:"bytes,3,opt,name=channelSecretId,proto3" json:"channelSecretId,omitempty"`
	ChannelAccessToken string `protobuf:"bytes,4,opt,name=channelAccessToken,proto3" json:"channelAccessToken,omitempty"`
	CreatedAt          string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt          string `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ChannelModel) Reset() {
	*x = ChannelModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelModel) ProtoMessage() {}

func (x *ChannelModel) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelModel.ProtoReflect.Descriptor instead.
func (*ChannelModel) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelModel) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ChannelModel) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelModel) GetChannelSecretId() string {
	if x != nil {
		return x.ChannelSecretId
	}
	return ""
}

func (x *ChannelModel) GetChannelAccessToken() string {
	if x != nil {
		return x.ChannelAccessToken
	}
	return ""
}

func (x *ChannelModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ChannelModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{1}
}

func (x *TestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TestReply) Reset() {
	*x = TestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReply) ProtoMessage() {}

func (x *TestReply) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReply.ProtoReflect.Descriptor instead.
func (*TestReply) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{2}
}

func (x *TestReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel *ChannelModel `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{4}
}

func (x *GetReply) GetChannel() *ChannelModel {
	if x != nil {
		return x.Channel
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId          string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
	ChannelSecretId    string `protobuf:"bytes,2,opt,name=channelSecretId,proto3" json:"channelSecretId,omitempty"`
	ChannelAccessToken string `protobuf:"bytes,3,opt,name=channelAccessToken,proto3" json:"channelAccessToken,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *CreateRequest) GetChannelSecretId() string {
	if x != nil {
		return x.ChannelSecretId
	}
	return ""
}

func (x *CreateRequest) GetChannelAccessToken() string {
	if x != nil {
		return x.ChannelAccessToken
	}
	return ""
}

type CreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel *ChannelModel `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *CreateReply) Reset() {
	*x = CreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_channel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReply) ProtoMessage() {}

func (x *CreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_channel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReply.ProtoReflect.Descriptor instead.
func (*CreateReply) Descriptor() ([]byte, []int) {
	return file_channel_proto_rawDescGZIP(), []int{6}
}

func (x *CreateReply) GetChannel() *ChannelModel {
	if x != nil {
		return x.Channel
	}
	return nil
}

var File_channel_proto protoreflect.FileDescriptor

var file_channel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xd2, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1d, 0x0a,
	0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x09,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x32, 0xa8, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x32,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_channel_proto_rawDescOnce sync.Once
	file_channel_proto_rawDescData = file_channel_proto_rawDesc
)

func file_channel_proto_rawDescGZIP() []byte {
	file_channel_proto_rawDescOnce.Do(func() {
		file_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_channel_proto_rawDescData)
	})
	return file_channel_proto_rawDescData
}

var file_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_channel_proto_goTypes = []interface{}{
	(*ChannelModel)(nil),  // 0: channel.ChannelModel
	(*TestRequest)(nil),   // 1: channel.TestRequest
	(*TestReply)(nil),     // 2: channel.TestReply
	(*GetRequest)(nil),    // 3: channel.GetRequest
	(*GetReply)(nil),      // 4: channel.GetReply
	(*CreateRequest)(nil), // 5: channel.CreateRequest
	(*CreateReply)(nil),   // 6: channel.CreateReply
}
var file_channel_proto_depIdxs = []int32{
	0, // 0: channel.GetReply.channel:type_name -> channel.ChannelModel
	0, // 1: channel.CreateReply.channel:type_name -> channel.ChannelModel
	1, // 2: channel.Channel.Test:input_type -> channel.TestRequest
	3, // 3: channel.Channel.Get:input_type -> channel.GetRequest
	5, // 4: channel.Channel.Create:input_type -> channel.CreateRequest
	2, // 5: channel.Channel.Test:output_type -> channel.TestReply
	4, // 6: channel.Channel.Get:output_type -> channel.GetReply
	6, // 7: channel.Channel.Create:output_type -> channel.CreateReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_channel_proto_init() }
func file_channel_proto_init() {
	if File_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelModel); i {
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
		file_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReply); i {
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
		file_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
		file_channel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_channel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReply); i {
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
			RawDescriptor: file_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_channel_proto_goTypes,
		DependencyIndexes: file_channel_proto_depIdxs,
		MessageInfos:      file_channel_proto_msgTypes,
	}.Build()
	File_channel_proto = out.File
	file_channel_proto_rawDesc = nil
	file_channel_proto_goTypes = nil
	file_channel_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChannelClient is the client API for Channel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChannelClient interface {
	// 疎通テスト用
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
}

type channelClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelClient(cc grpc.ClientConnInterface) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error) {
	out := new(TestReply)
	err := c.cc.Invoke(ctx, "/channel.Channel/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/channel.Channel/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/channel.Channel/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
type ChannelServer interface {
	// 疎通テスト用
	Test(context.Context, *TestRequest) (*TestReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Create(context.Context, *CreateRequest) (*CreateReply, error)
}

// UnimplementedChannelServer can be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (*UnimplementedChannelServer) Test(context.Context, *TestRequest) (*TestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (*UnimplementedChannelServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedChannelServer) Create(context.Context, *CreateRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterChannelServer(s *grpc.Server, srv ChannelServer) {
	s.RegisterService(&_Channel_serviceDesc, srv)
}

func _Channel_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.Channel/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Channel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "channel.Channel",
	HandlerType: (*ChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Channel_Test_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Channel_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Channel_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel.proto",
}
