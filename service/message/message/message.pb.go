// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: message.proto

package message

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int32   `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PageNumber *int32  `protobuf:"varint,2,opt,name=pageNumber,proto3,oneof" json:"pageNumber,omitempty"`
	LastId     *uint64 `protobuf:"varint,3,opt,name=lastId,proto3,oneof" json:"lastId,omitempty"`
	HasCount   bool    `protobuf:"varint,4,opt,name=hasCount,proto3" json:"hasCount,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *PageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageRequest) GetPageNumber() int32 {
	if x != nil && x.PageNumber != nil {
		return *x.PageNumber
	}
	return 0
}

func (x *PageRequest) GetLastId() uint64 {
	if x != nil && x.LastId != nil {
		return *x.LastId
	}
	return 0
}

func (x *PageRequest) GetHasCount() bool {
	if x != nil {
		return x.HasCount
	}
	return false
}

type HasNoReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receiver uint64 `protobuf:"varint,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (x *HasNoReadReq) Reset() {
	*x = HasNoReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasNoReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasNoReadReq) ProtoMessage() {}

func (x *HasNoReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasNoReadReq.ProtoReflect.Descriptor instead.
func (*HasNoReadReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *HasNoReadReq) GetReceiver() uint64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

type HasNoReadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *HasNoReadResp) Reset() {
	*x = HasNoReadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasNoReadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasNoReadResp) ProtoMessage() {}

func (x *HasNoReadResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasNoReadResp.ProtoReflect.Descriptor instead.
func (*HasNoReadResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *HasNoReadResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type MessageByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageByIdReq) Reset() {
	*x = MessageByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageByIdReq) ProtoMessage() {}

func (x *MessageByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageByIdReq.ProtoReflect.Descriptor instead.
func (*MessageByIdReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *MessageByIdReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MessageByUserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MessageByUserIdReq) Reset() {
	*x = MessageByUserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageByUserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageByUserIdReq) ProtoMessage() {}

func (x *MessageByUserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageByUserIdReq.ProtoReflect.Descriptor instead.
func (*MessageByUserIdReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *MessageByUserIdReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type MessageReqByReceiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receiver uint64       `protobuf:"varint,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Noread   *bool        `protobuf:"varint,2,opt,name=noread,proto3,oneof" json:"noread,omitempty"`
	Page     *PageRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *MessageReqByReceiver) Reset() {
	*x = MessageReqByReceiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReqByReceiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReqByReceiver) ProtoMessage() {}

func (x *MessageReqByReceiver) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReqByReceiver.ProtoReflect.Descriptor instead.
func (*MessageReqByReceiver) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *MessageReqByReceiver) GetReceiver() uint64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *MessageReqByReceiver) GetNoread() bool {
	if x != nil && x.Noread != nil {
		return *x.Noread
	}
	return false
}

func (x *MessageReqByReceiver) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type MessageDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Receiver  uint64 `protobuf:"varint,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Time      int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Read      bool   `protobuf:"varint,4,opt,name=read,proto3" json:"read,omitempty"`
	Type      int32  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	IsDeleted bool   `protobuf:"varint,6,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Title     string `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MessageDto) Reset() {
	*x = MessageDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageDto) ProtoMessage() {}

func (x *MessageDto) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageDto.ProtoReflect.Descriptor instead.
func (*MessageDto) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *MessageDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageDto) GetReceiver() uint64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *MessageDto) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *MessageDto) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *MessageDto) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessageDto) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *MessageDto) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MessageDto) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MessagePage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32         `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Messages   []*MessageDto `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessagePage) Reset() {
	*x = MessagePage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePage) ProtoMessage() {}

func (x *MessagePage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePage.ProtoReflect.Descriptor instead.
func (*MessagePage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *MessagePage) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *MessagePage) GetMessages() []*MessageDto {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0c, 0x48,
	0x61, 0x73, 0x4e, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x4e, 0x6f,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x20, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x06, 0x6e, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x06, 0x6e, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6e,
	0x6f, 0x72, 0x65, 0x61, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x0b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x32, 0xfa, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x42, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x3e, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x45, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x14, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x40, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x12,
	0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c,
	0x6c, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_message_proto_goTypes = []interface{}{
	(*Empty)(nil),                // 0: messageclient.Empty
	(*PageRequest)(nil),          // 1: messageclient.PageRequest
	(*HasNoReadReq)(nil),         // 2: messageclient.HasNoReadReq
	(*HasNoReadResp)(nil),        // 3: messageclient.HasNoReadResp
	(*MessageByIdReq)(nil),       // 4: messageclient.MessageByIdReq
	(*MessageByUserIdReq)(nil),   // 5: messageclient.MessageByUserIdReq
	(*MessageReqByReceiver)(nil), // 6: messageclient.MessageReqByReceiver
	(*MessageDto)(nil),           // 7: messageclient.MessageDto
	(*MessagePage)(nil),          // 8: messageclient.MessagePage
}
var file_message_proto_depIdxs = []int32{
	1, // 0: messageclient.MessageReqByReceiver.page:type_name -> messageclient.PageRequest
	7, // 1: messageclient.MessagePage.messages:type_name -> messageclient.MessageDto
	6, // 2: messageclient.Message.GetMessagePage:input_type -> messageclient.MessageReqByReceiver
	4, // 3: messageclient.Message.ReadMsg:input_type -> messageclient.MessageByIdReq
	5, // 4: messageclient.Message.ReadAllMsg:input_type -> messageclient.MessageByUserIdReq
	4, // 5: messageclient.Message.GetMessageById:input_type -> messageclient.MessageByIdReq
	7, // 6: messageclient.Message.SendMessage:input_type -> messageclient.MessageDto
	4, // 7: messageclient.Message.DeleteMsg:input_type -> messageclient.MessageByIdReq
	5, // 8: messageclient.Message.DeleteAllMsg:input_type -> messageclient.MessageByUserIdReq
	8, // 9: messageclient.Message.GetMessagePage:output_type -> messageclient.MessagePage
	0, // 10: messageclient.Message.ReadMsg:output_type -> messageclient.Empty
	0, // 11: messageclient.Message.ReadAllMsg:output_type -> messageclient.Empty
	7, // 12: messageclient.Message.GetMessageById:output_type -> messageclient.MessageDto
	0, // 13: messageclient.Message.SendMessage:output_type -> messageclient.Empty
	0, // 14: messageclient.Message.DeleteMsg:output_type -> messageclient.Empty
	0, // 15: messageclient.Message.DeleteAllMsg:output_type -> messageclient.Empty
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasNoReadReq); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasNoReadResp); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageByIdReq); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageByUserIdReq); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReqByReceiver); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageDto); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePage); i {
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
	file_message_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_message_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageClient interface {
	GetMessagePage(ctx context.Context, in *MessageReqByReceiver, opts ...grpc.CallOption) (*MessagePage, error)
	ReadMsg(ctx context.Context, in *MessageByIdReq, opts ...grpc.CallOption) (*Empty, error)
	ReadAllMsg(ctx context.Context, in *MessageByUserIdReq, opts ...grpc.CallOption) (*Empty, error)
	GetMessageById(ctx context.Context, in *MessageByIdReq, opts ...grpc.CallOption) (*MessageDto, error)
	SendMessage(ctx context.Context, in *MessageDto, opts ...grpc.CallOption) (*Empty, error)
	DeleteMsg(ctx context.Context, in *MessageByIdReq, opts ...grpc.CallOption) (*Empty, error)
	DeleteAllMsg(ctx context.Context, in *MessageByUserIdReq, opts ...grpc.CallOption) (*Empty, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) GetMessagePage(ctx context.Context, in *MessageReqByReceiver, opts ...grpc.CallOption) (*MessagePage, error) {
	out := new(MessagePage)
	err := c.cc.Invoke(ctx, "/messageclient.Message/GetMessagePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) ReadMsg(ctx context.Context, in *MessageByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/messageclient.Message/ReadMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) ReadAllMsg(ctx context.Context, in *MessageByUserIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/messageclient.Message/ReadAllMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetMessageById(ctx context.Context, in *MessageByIdReq, opts ...grpc.CallOption) (*MessageDto, error) {
	out := new(MessageDto)
	err := c.cc.Invoke(ctx, "/messageclient.Message/GetMessageById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendMessage(ctx context.Context, in *MessageDto, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/messageclient.Message/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) DeleteMsg(ctx context.Context, in *MessageByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/messageclient.Message/DeleteMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) DeleteAllMsg(ctx context.Context, in *MessageByUserIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/messageclient.Message/DeleteAllMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
type MessageServer interface {
	GetMessagePage(context.Context, *MessageReqByReceiver) (*MessagePage, error)
	ReadMsg(context.Context, *MessageByIdReq) (*Empty, error)
	ReadAllMsg(context.Context, *MessageByUserIdReq) (*Empty, error)
	GetMessageById(context.Context, *MessageByIdReq) (*MessageDto, error)
	SendMessage(context.Context, *MessageDto) (*Empty, error)
	DeleteMsg(context.Context, *MessageByIdReq) (*Empty, error)
	DeleteAllMsg(context.Context, *MessageByUserIdReq) (*Empty, error)
}

// UnimplementedMessageServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (*UnimplementedMessageServer) GetMessagePage(context.Context, *MessageReqByReceiver) (*MessagePage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagePage not implemented")
}
func (*UnimplementedMessageServer) ReadMsg(context.Context, *MessageByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMsg not implemented")
}
func (*UnimplementedMessageServer) ReadAllMsg(context.Context, *MessageByUserIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllMsg not implemented")
}
func (*UnimplementedMessageServer) GetMessageById(context.Context, *MessageByIdReq) (*MessageDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageById not implemented")
}
func (*UnimplementedMessageServer) SendMessage(context.Context, *MessageDto) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedMessageServer) DeleteMsg(context.Context, *MessageByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMsg not implemented")
}
func (*UnimplementedMessageServer) DeleteAllMsg(context.Context, *MessageByUserIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllMsg not implemented")
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_GetMessagePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReqByReceiver)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetMessagePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/GetMessagePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetMessagePage(ctx, req.(*MessageReqByReceiver))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_ReadMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).ReadMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/ReadMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).ReadMsg(ctx, req.(*MessageByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_ReadAllMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).ReadAllMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/ReadAllMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).ReadAllMsg(ctx, req.(*MessageByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetMessageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetMessageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/GetMessageById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetMessageById(ctx, req.(*MessageByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMessage(ctx, req.(*MessageDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_DeleteMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).DeleteMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/DeleteMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).DeleteMsg(ctx, req.(*MessageByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_DeleteAllMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).DeleteAllMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messageclient.Message/DeleteAllMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).DeleteAllMsg(ctx, req.(*MessageByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messageclient.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessagePage",
			Handler:    _Message_GetMessagePage_Handler,
		},
		{
			MethodName: "ReadMsg",
			Handler:    _Message_ReadMsg_Handler,
		},
		{
			MethodName: "ReadAllMsg",
			Handler:    _Message_ReadAllMsg_Handler,
		},
		{
			MethodName: "GetMessageById",
			Handler:    _Message_GetMessageById_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Message_SendMessage_Handler,
		},
		{
			MethodName: "DeleteMsg",
			Handler:    _Message_DeleteMsg_Handler,
		},
		{
			MethodName: "DeleteAllMsg",
			Handler:    _Message_DeleteAllMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
