// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: ask.proto

package ask

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
		mi := &file_ask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[0]
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
	return file_ask_proto_rawDescGZIP(), []int{0}
}

type AskDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid      uint64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Time     int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title    string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AskDetail) Reset() {
	*x = AskDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskDetail) ProtoMessage() {}

func (x *AskDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskDetail.ProtoReflect.Descriptor instead.
func (*AskDetail) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{1}
}

func (x *AskDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AskDetail) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AskDetail) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AskDetail) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AskDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AskDetail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AskSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid      uint64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Time     int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title    string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *AskSummary) Reset() {
	*x = AskSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskSummary) ProtoMessage() {}

func (x *AskSummary) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskSummary.ProtoReflect.Descriptor instead.
func (*AskSummary) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{2}
}

func (x *AskSummary) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AskSummary) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AskSummary) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AskSummary) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AskSummary) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ReplyDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AskId    uint64 `protobuf:"varint,2,opt,name=askId,proto3" json:"askId,omitempty"`
	Uid      uint64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Time     int64  `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Nickname string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Content  string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ReplyDetail) Reset() {
	*x = ReplyDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyDetail) ProtoMessage() {}

func (x *ReplyDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyDetail.ProtoReflect.Descriptor instead.
func (*ReplyDetail) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReplyDetail) GetAskId() uint64 {
	if x != nil {
		return x.AskId
	}
	return 0
}

func (x *ReplyDetail) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReplyDetail) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ReplyDetail) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ReplyDetail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AskByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AskByIdReq) Reset() {
	*x = AskByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskByIdReq) ProtoMessage() {}

func (x *AskByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskByIdReq.ProtoReflect.Descriptor instead.
func (*AskByIdReq) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{4}
}

func (x *AskByIdReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReplyByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReplyByIdReq) Reset() {
	*x = ReplyByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyByIdReq) ProtoMessage() {}

func (x *ReplyByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyByIdReq.ProtoReflect.Descriptor instead.
func (*ReplyByIdReq) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyByIdReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asks []*AskSummary `protobuf:"bytes,1,rep,name=asks,proto3" json:"asks,omitempty"`
}

func (x *AskList) Reset() {
	*x = AskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskList) ProtoMessage() {}

func (x *AskList) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskList.ProtoReflect.Descriptor instead.
func (*AskList) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{6}
}

func (x *AskList) GetAsks() []*AskSummary {
	if x != nil {
		return x.Asks
	}
	return nil
}

type ReplyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replies []*ReplyDetail `protobuf:"bytes,1,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *ReplyList) Reset() {
	*x = ReplyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ask_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyList) ProtoMessage() {}

func (x *ReplyList) ProtoReflect() protoreflect.Message {
	mi := &file_ask_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyList.ProtoReflect.Descriptor instead.
func (*ReplyList) Descriptor() ([]byte, []int) {
	return file_ask_proto_rawDescGZIP(), []int{7}
}

func (x *ReplyList) GetReplies() []*ReplyDetail {
	if x != nil {
		return x.Replies
	}
	return nil
}

var File_ask_proto protoreflect.FileDescriptor

var file_ask_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x73, 0x6b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x8d, 0x01, 0x0a, 0x09, 0x41, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x74, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1c, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x07, 0x41, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x6b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x32, 0xb6, 0x03, 0x0a, 0x03, 0x41,
	0x73, 0x6b, 0x12, 0x3a, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x41, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x15, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x31,
	0x0a, 0x09, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x6b, 0x12, 0x10, 0x2e, 0x61, 0x73,
	0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e,
	0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x41, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x61, 0x73,
	0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x1a, 0x10, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x65, 0x64, 0x69, 0x74, 0x41, 0x73, 0x6b, 0x12, 0x14,
	0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x1a, 0x10, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x10, 0x2e, 0x61, 0x73, 0x6b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x09,
	0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x2e, 0x61, 0x73, 0x6b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x1a, 0x10, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x73, 0x6b,
	0x12, 0x15, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x61, 0x73, 0x6b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x05, 0x5a, 0x03, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ask_proto_rawDescOnce sync.Once
	file_ask_proto_rawDescData = file_ask_proto_rawDesc
)

func file_ask_proto_rawDescGZIP() []byte {
	file_ask_proto_rawDescOnce.Do(func() {
		file_ask_proto_rawDescData = protoimpl.X.CompressGZIP(file_ask_proto_rawDescData)
	})
	return file_ask_proto_rawDescData
}

var file_ask_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ask_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: askclient.Empty
	(*AskDetail)(nil),    // 1: askclient.AskDetail
	(*AskSummary)(nil),   // 2: askclient.AskSummary
	(*ReplyDetail)(nil),  // 3: askclient.ReplyDetail
	(*AskByIdReq)(nil),   // 4: askclient.AskByIdReq
	(*ReplyByIdReq)(nil), // 5: askclient.ReplyByIdReq
	(*AskList)(nil),      // 6: askclient.AskList
	(*ReplyList)(nil),    // 7: askclient.ReplyList
}
var file_ask_proto_depIdxs = []int32{
	2,  // 0: askclient.AskList.asks:type_name -> askclient.AskSummary
	3,  // 1: askclient.ReplyList.replies:type_name -> askclient.ReplyDetail
	4,  // 2: askclient.Ask.getAskById:input_type -> askclient.AskByIdReq
	0,  // 3: askclient.Ask.getAllAsk:input_type -> askclient.Empty
	1,  // 4: askclient.Ask.addAsk:input_type -> askclient.AskDetail
	1,  // 5: askclient.Ask.editAsk:input_type -> askclient.AskDetail
	3,  // 6: askclient.Ask.addReply:input_type -> askclient.ReplyDetail
	3,  // 7: askclient.Ask.editReply:input_type -> askclient.ReplyDetail
	4,  // 8: askclient.Ask.removeAsk:input_type -> askclient.AskByIdReq
	5,  // 9: askclient.Ask.removeReply:input_type -> askclient.ReplyByIdReq
	2,  // 10: askclient.Ask.getAskById:output_type -> askclient.AskSummary
	6,  // 11: askclient.Ask.getAllAsk:output_type -> askclient.AskList
	0,  // 12: askclient.Ask.addAsk:output_type -> askclient.Empty
	0,  // 13: askclient.Ask.editAsk:output_type -> askclient.Empty
	0,  // 14: askclient.Ask.addReply:output_type -> askclient.Empty
	0,  // 15: askclient.Ask.editReply:output_type -> askclient.Empty
	0,  // 16: askclient.Ask.removeAsk:output_type -> askclient.Empty
	0,  // 17: askclient.Ask.removeReply:output_type -> askclient.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_ask_proto_init() }
func file_ask_proto_init() {
	if File_ask_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ask_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ask_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskDetail); i {
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
		file_ask_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskSummary); i {
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
		file_ask_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyDetail); i {
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
		file_ask_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskByIdReq); i {
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
		file_ask_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyByIdReq); i {
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
		file_ask_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskList); i {
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
		file_ask_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyList); i {
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
			RawDescriptor: file_ask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ask_proto_goTypes,
		DependencyIndexes: file_ask_proto_depIdxs,
		MessageInfos:      file_ask_proto_msgTypes,
	}.Build()
	File_ask_proto = out.File
	file_ask_proto_rawDesc = nil
	file_ask_proto_goTypes = nil
	file_ask_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AskClient is the client API for Ask service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AskClient interface {
	GetAskById(ctx context.Context, in *AskByIdReq, opts ...grpc.CallOption) (*AskSummary, error)
	GetAllAsk(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AskList, error)
	AddAsk(ctx context.Context, in *AskDetail, opts ...grpc.CallOption) (*Empty, error)
	EditAsk(ctx context.Context, in *AskDetail, opts ...grpc.CallOption) (*Empty, error)
	AddReply(ctx context.Context, in *ReplyDetail, opts ...grpc.CallOption) (*Empty, error)
	EditReply(ctx context.Context, in *ReplyDetail, opts ...grpc.CallOption) (*Empty, error)
	RemoveAsk(ctx context.Context, in *AskByIdReq, opts ...grpc.CallOption) (*Empty, error)
	RemoveReply(ctx context.Context, in *ReplyByIdReq, opts ...grpc.CallOption) (*Empty, error)
}

type askClient struct {
	cc grpc.ClientConnInterface
}

func NewAskClient(cc grpc.ClientConnInterface) AskClient {
	return &askClient{cc}
}

func (c *askClient) GetAskById(ctx context.Context, in *AskByIdReq, opts ...grpc.CallOption) (*AskSummary, error) {
	out := new(AskSummary)
	err := c.cc.Invoke(ctx, "/askclient.Ask/getAskById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) GetAllAsk(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AskList, error) {
	out := new(AskList)
	err := c.cc.Invoke(ctx, "/askclient.Ask/getAllAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) AddAsk(ctx context.Context, in *AskDetail, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/askclient.Ask/addAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) EditAsk(ctx context.Context, in *AskDetail, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/askclient.Ask/editAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) AddReply(ctx context.Context, in *ReplyDetail, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/askclient.Ask/addReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) EditReply(ctx context.Context, in *ReplyDetail, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/askclient.Ask/editReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) RemoveAsk(ctx context.Context, in *AskByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/askclient.Ask/removeAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *askClient) RemoveReply(ctx context.Context, in *ReplyByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/askclient.Ask/removeReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AskServer is the server API for Ask service.
type AskServer interface {
	GetAskById(context.Context, *AskByIdReq) (*AskSummary, error)
	GetAllAsk(context.Context, *Empty) (*AskList, error)
	AddAsk(context.Context, *AskDetail) (*Empty, error)
	EditAsk(context.Context, *AskDetail) (*Empty, error)
	AddReply(context.Context, *ReplyDetail) (*Empty, error)
	EditReply(context.Context, *ReplyDetail) (*Empty, error)
	RemoveAsk(context.Context, *AskByIdReq) (*Empty, error)
	RemoveReply(context.Context, *ReplyByIdReq) (*Empty, error)
}

// UnimplementedAskServer can be embedded to have forward compatible implementations.
type UnimplementedAskServer struct {
}

func (*UnimplementedAskServer) GetAskById(context.Context, *AskByIdReq) (*AskSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAskById not implemented")
}
func (*UnimplementedAskServer) GetAllAsk(context.Context, *Empty) (*AskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAsk not implemented")
}
func (*UnimplementedAskServer) AddAsk(context.Context, *AskDetail) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAsk not implemented")
}
func (*UnimplementedAskServer) EditAsk(context.Context, *AskDetail) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAsk not implemented")
}
func (*UnimplementedAskServer) AddReply(context.Context, *ReplyDetail) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReply not implemented")
}
func (*UnimplementedAskServer) EditReply(context.Context, *ReplyDetail) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditReply not implemented")
}
func (*UnimplementedAskServer) RemoveAsk(context.Context, *AskByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAsk not implemented")
}
func (*UnimplementedAskServer) RemoveReply(context.Context, *ReplyByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReply not implemented")
}

func RegisterAskServer(s *grpc.Server, srv AskServer) {
	s.RegisterService(&_Ask_serviceDesc, srv)
}

func _Ask_GetAskById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).GetAskById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/GetAskById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).GetAskById(ctx, req.(*AskByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_GetAllAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).GetAllAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/GetAllAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).GetAllAsk(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_AddAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).AddAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/AddAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).AddAsk(ctx, req.(*AskDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_EditAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).EditAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/EditAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).EditAsk(ctx, req.(*AskDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_AddReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).AddReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/AddReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).AddReply(ctx, req.(*ReplyDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_EditReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).EditReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/EditReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).EditReply(ctx, req.(*ReplyDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_RemoveAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).RemoveAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/RemoveAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).RemoveAsk(ctx, req.(*AskByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ask_RemoveReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskServer).RemoveReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/askclient.Ask/RemoveReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskServer).RemoveReply(ctx, req.(*ReplyByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "askclient.Ask",
	HandlerType: (*AskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAskById",
			Handler:    _Ask_GetAskById_Handler,
		},
		{
			MethodName: "getAllAsk",
			Handler:    _Ask_GetAllAsk_Handler,
		},
		{
			MethodName: "addAsk",
			Handler:    _Ask_AddAsk_Handler,
		},
		{
			MethodName: "editAsk",
			Handler:    _Ask_EditAsk_Handler,
		},
		{
			MethodName: "addReply",
			Handler:    _Ask_AddReply_Handler,
		},
		{
			MethodName: "editReply",
			Handler:    _Ask_EditReply_Handler,
		},
		{
			MethodName: "removeAsk",
			Handler:    _Ask_RemoveAsk_Handler,
		},
		{
			MethodName: "removeReply",
			Handler:    _Ask_RemoveReply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ask.proto",
}
