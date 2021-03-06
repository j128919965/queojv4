// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: solution.proto

package solution

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
		mi := &file_solution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[0]
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
	return file_solution_proto_rawDescGZIP(), []int{0}
}

type SolutionDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid       uint64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Pid       int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Time      uint64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Nickname  string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title     string `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Summary   string `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	Content   string `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	IsTeacher bool   `protobuf:"varint,9,opt,name=isTeacher,proto3" json:"isTeacher,omitempty"`
}

func (x *SolutionDetail) Reset() {
	*x = SolutionDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionDetail) ProtoMessage() {}

func (x *SolutionDetail) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionDetail.ProtoReflect.Descriptor instead.
func (*SolutionDetail) Descriptor() ([]byte, []int) {
	return file_solution_proto_rawDescGZIP(), []int{1}
}

func (x *SolutionDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SolutionDetail) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SolutionDetail) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SolutionDetail) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *SolutionDetail) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SolutionDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SolutionDetail) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *SolutionDetail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SolutionDetail) GetIsTeacher() bool {
	if x != nil {
		return x.IsTeacher
	}
	return false
}

type SolutionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time      uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title     string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Summary   string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Pid       int32  `protobuf:"varint,6,opt,name=pid,proto3" json:"pid,omitempty"`
	IsTeacher bool   `protobuf:"varint,7,opt,name=isTeacher,proto3" json:"isTeacher,omitempty"`
}

func (x *SolutionSummary) Reset() {
	*x = SolutionSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionSummary) ProtoMessage() {}

func (x *SolutionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionSummary.ProtoReflect.Descriptor instead.
func (*SolutionSummary) Descriptor() ([]byte, []int) {
	return file_solution_proto_rawDescGZIP(), []int{2}
}

func (x *SolutionSummary) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SolutionSummary) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *SolutionSummary) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SolutionSummary) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SolutionSummary) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *SolutionSummary) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SolutionSummary) GetIsTeacher() bool {
	if x != nil {
		return x.IsTeacher
	}
	return false
}

type SolutionAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid       int32  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Uid       uint64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title     string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	IsTeacher bool   `protobuf:"varint,6,opt,name=isTeacher,proto3" json:"isTeacher,omitempty"`
}

func (x *SolutionAddReq) Reset() {
	*x = SolutionAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionAddReq) ProtoMessage() {}

func (x *SolutionAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionAddReq.ProtoReflect.Descriptor instead.
func (*SolutionAddReq) Descriptor() ([]byte, []int) {
	return file_solution_proto_rawDescGZIP(), []int{3}
}

func (x *SolutionAddReq) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SolutionAddReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SolutionAddReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SolutionAddReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SolutionAddReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SolutionAddReq) GetIsTeacher() bool {
	if x != nil {
		return x.IsTeacher
	}
	return false
}

type AllSolutionByPidReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *AllSolutionByPidReq) Reset() {
	*x = AllSolutionByPidReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSolutionByPidReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSolutionByPidReq) ProtoMessage() {}

func (x *AllSolutionByPidReq) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSolutionByPidReq.ProtoReflect.Descriptor instead.
func (*AllSolutionByPidReq) Descriptor() ([]byte, []int) {
	return file_solution_proto_rawDescGZIP(), []int{4}
}

func (x *AllSolutionByPidReq) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type SolutionByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SolutionByIdReq) Reset() {
	*x = SolutionByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionByIdReq) ProtoMessage() {}

func (x *SolutionByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionByIdReq.ProtoReflect.Descriptor instead.
func (*SolutionByIdReq) Descriptor() ([]byte, []int) {
	return file_solution_proto_rawDescGZIP(), []int{5}
}

func (x *SolutionByIdReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SolutionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solutions []*SolutionSummary `protobuf:"bytes,1,rep,name=solutions,proto3" json:"solutions,omitempty"`
}

func (x *SolutionList) Reset() {
	*x = SolutionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solution_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolutionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionList) ProtoMessage() {}

func (x *SolutionList) ProtoReflect() protoreflect.Message {
	mi := &file_solution_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionList.ProtoReflect.Descriptor instead.
func (*SolutionList) Descriptor() ([]byte, []int) {
	return file_solution_proto_rawDescGZIP(), []int{6}
}

func (x *SolutionList) GetSolutions() []*SolutionSummary {
	if x != nil {
		return x.Solutions
	}
	return nil
}

var File_solution_proto protoreflect.FileDescriptor

var file_solution_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xdc, 0x01, 0x0a, 0x0e, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x0f, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x9e, 0x01, 0x0a,
	0x0e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x27, 0x0a,
	0x13, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x0c, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x09, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xd5, 0x03, 0x0a, 0x08, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x58, 0x0a, 0x13, 0x67,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50,
	0x69, 0x64, 0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0e, 0x67,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0c, 0x65, 0x64, 0x69,
	0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x15, 0x2e, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0a, 0x5a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_solution_proto_rawDescOnce sync.Once
	file_solution_proto_rawDescData = file_solution_proto_rawDesc
)

func file_solution_proto_rawDescGZIP() []byte {
	file_solution_proto_rawDescOnce.Do(func() {
		file_solution_proto_rawDescData = protoimpl.X.CompressGZIP(file_solution_proto_rawDescData)
	})
	return file_solution_proto_rawDescData
}

var file_solution_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_solution_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: solutionclient.Empty
	(*SolutionDetail)(nil),      // 1: solutionclient.SolutionDetail
	(*SolutionSummary)(nil),     // 2: solutionclient.SolutionSummary
	(*SolutionAddReq)(nil),      // 3: solutionclient.SolutionAddReq
	(*AllSolutionByPidReq)(nil), // 4: solutionclient.AllSolutionByPidReq
	(*SolutionByIdReq)(nil),     // 5: solutionclient.SolutionByIdReq
	(*SolutionList)(nil),        // 6: solutionclient.SolutionList
}
var file_solution_proto_depIdxs = []int32{
	2, // 0: solutionclient.SolutionList.solutions:type_name -> solutionclient.SolutionSummary
	5, // 1: solutionclient.Solution.getSolutionDetail:input_type -> solutionclient.SolutionByIdReq
	4, // 2: solutionclient.Solution.getAllSolutionByPid:input_type -> solutionclient.AllSolutionByPidReq
	3, // 3: solutionclient.Solution.addSolution:input_type -> solutionclient.SolutionAddReq
	0, // 4: solutionclient.Solution.getAllSolution:input_type -> solutionclient.Empty
	5, // 5: solutionclient.Solution.delSolution:input_type -> solutionclient.SolutionByIdReq
	1, // 6: solutionclient.Solution.editSolution:input_type -> solutionclient.SolutionDetail
	1, // 7: solutionclient.Solution.getSolutionDetail:output_type -> solutionclient.SolutionDetail
	6, // 8: solutionclient.Solution.getAllSolutionByPid:output_type -> solutionclient.SolutionList
	0, // 9: solutionclient.Solution.addSolution:output_type -> solutionclient.Empty
	6, // 10: solutionclient.Solution.getAllSolution:output_type -> solutionclient.SolutionList
	0, // 11: solutionclient.Solution.delSolution:output_type -> solutionclient.Empty
	0, // 12: solutionclient.Solution.editSolution:output_type -> solutionclient.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_solution_proto_init() }
func file_solution_proto_init() {
	if File_solution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_solution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_solution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionDetail); i {
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
		file_solution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionSummary); i {
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
		file_solution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionAddReq); i {
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
		file_solution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSolutionByPidReq); i {
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
		file_solution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionByIdReq); i {
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
		file_solution_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolutionList); i {
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
			RawDescriptor: file_solution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_solution_proto_goTypes,
		DependencyIndexes: file_solution_proto_depIdxs,
		MessageInfos:      file_solution_proto_msgTypes,
	}.Build()
	File_solution_proto = out.File
	file_solution_proto_rawDesc = nil
	file_solution_proto_goTypes = nil
	file_solution_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SolutionClient is the client API for Solution service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SolutionClient interface {
	GetSolutionDetail(ctx context.Context, in *SolutionByIdReq, opts ...grpc.CallOption) (*SolutionDetail, error)
	GetAllSolutionByPid(ctx context.Context, in *AllSolutionByPidReq, opts ...grpc.CallOption) (*SolutionList, error)
	AddSolution(ctx context.Context, in *SolutionAddReq, opts ...grpc.CallOption) (*Empty, error)
	GetAllSolution(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SolutionList, error)
	DelSolution(ctx context.Context, in *SolutionByIdReq, opts ...grpc.CallOption) (*Empty, error)
	EditSolution(ctx context.Context, in *SolutionDetail, opts ...grpc.CallOption) (*Empty, error)
}

type solutionClient struct {
	cc grpc.ClientConnInterface
}

func NewSolutionClient(cc grpc.ClientConnInterface) SolutionClient {
	return &solutionClient{cc}
}

func (c *solutionClient) GetSolutionDetail(ctx context.Context, in *SolutionByIdReq, opts ...grpc.CallOption) (*SolutionDetail, error) {
	out := new(SolutionDetail)
	err := c.cc.Invoke(ctx, "/solutionclient.Solution/getSolutionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solutionClient) GetAllSolutionByPid(ctx context.Context, in *AllSolutionByPidReq, opts ...grpc.CallOption) (*SolutionList, error) {
	out := new(SolutionList)
	err := c.cc.Invoke(ctx, "/solutionclient.Solution/getAllSolutionByPid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solutionClient) AddSolution(ctx context.Context, in *SolutionAddReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/solutionclient.Solution/addSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solutionClient) GetAllSolution(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SolutionList, error) {
	out := new(SolutionList)
	err := c.cc.Invoke(ctx, "/solutionclient.Solution/getAllSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solutionClient) DelSolution(ctx context.Context, in *SolutionByIdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/solutionclient.Solution/delSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solutionClient) EditSolution(ctx context.Context, in *SolutionDetail, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/solutionclient.Solution/editSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolutionServer is the server API for Solution service.
type SolutionServer interface {
	GetSolutionDetail(context.Context, *SolutionByIdReq) (*SolutionDetail, error)
	GetAllSolutionByPid(context.Context, *AllSolutionByPidReq) (*SolutionList, error)
	AddSolution(context.Context, *SolutionAddReq) (*Empty, error)
	GetAllSolution(context.Context, *Empty) (*SolutionList, error)
	DelSolution(context.Context, *SolutionByIdReq) (*Empty, error)
	EditSolution(context.Context, *SolutionDetail) (*Empty, error)
}

// UnimplementedSolutionServer can be embedded to have forward compatible implementations.
type UnimplementedSolutionServer struct {
}

func (*UnimplementedSolutionServer) GetSolutionDetail(context.Context, *SolutionByIdReq) (*SolutionDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSolutionDetail not implemented")
}
func (*UnimplementedSolutionServer) GetAllSolutionByPid(context.Context, *AllSolutionByPidReq) (*SolutionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSolutionByPid not implemented")
}
func (*UnimplementedSolutionServer) AddSolution(context.Context, *SolutionAddReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSolution not implemented")
}
func (*UnimplementedSolutionServer) GetAllSolution(context.Context, *Empty) (*SolutionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSolution not implemented")
}
func (*UnimplementedSolutionServer) DelSolution(context.Context, *SolutionByIdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelSolution not implemented")
}
func (*UnimplementedSolutionServer) EditSolution(context.Context, *SolutionDetail) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSolution not implemented")
}

func RegisterSolutionServer(s *grpc.Server, srv SolutionServer) {
	s.RegisterService(&_Solution_serviceDesc, srv)
}

func _Solution_GetSolutionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolutionByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolutionServer).GetSolutionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solutionclient.Solution/GetSolutionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolutionServer).GetSolutionDetail(ctx, req.(*SolutionByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solution_GetAllSolutionByPid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllSolutionByPidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolutionServer).GetAllSolutionByPid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solutionclient.Solution/GetAllSolutionByPid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolutionServer).GetAllSolutionByPid(ctx, req.(*AllSolutionByPidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solution_AddSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolutionAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolutionServer).AddSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solutionclient.Solution/AddSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolutionServer).AddSolution(ctx, req.(*SolutionAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solution_GetAllSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolutionServer).GetAllSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solutionclient.Solution/GetAllSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolutionServer).GetAllSolution(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solution_DelSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolutionByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolutionServer).DelSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solutionclient.Solution/DelSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolutionServer).DelSolution(ctx, req.(*SolutionByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Solution_EditSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolutionDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolutionServer).EditSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solutionclient.Solution/EditSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolutionServer).EditSolution(ctx, req.(*SolutionDetail))
	}
	return interceptor(ctx, in, info, handler)
}

var _Solution_serviceDesc = grpc.ServiceDesc{
	ServiceName: "solutionclient.Solution",
	HandlerType: (*SolutionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSolutionDetail",
			Handler:    _Solution_GetSolutionDetail_Handler,
		},
		{
			MethodName: "getAllSolutionByPid",
			Handler:    _Solution_GetAllSolutionByPid_Handler,
		},
		{
			MethodName: "addSolution",
			Handler:    _Solution_AddSolution_Handler,
		},
		{
			MethodName: "getAllSolution",
			Handler:    _Solution_GetAllSolution_Handler,
		},
		{
			MethodName: "delSolution",
			Handler:    _Solution_DelSolution_Handler,
		},
		{
			MethodName: "editSolution",
			Handler:    _Solution_EditSolution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "solution.proto",
}
