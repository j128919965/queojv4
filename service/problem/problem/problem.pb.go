// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: problem.proto

package problem

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Integer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Integer) Reset() {
	*x = Integer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integer) ProtoMessage() {}

func (x *Integer) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integer.ProtoReflect.Descriptor instead.
func (*Integer) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{1}
}

func (x *Integer) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ProblemDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Point       int32  `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
	Level       int32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Tags        string `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	TimeLimit   uint64 `protobuf:"varint,7,opt,name=timeLimit,proto3" json:"timeLimit,omitempty"`
	SpaceLimit  uint64 `protobuf:"varint,8,opt,name=spaceLimit,proto3" json:"spaceLimit,omitempty"`
}

func (x *ProblemDetail) Reset() {
	*x = ProblemDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemDetail) ProtoMessage() {}

func (x *ProblemDetail) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemDetail.ProtoReflect.Descriptor instead.
func (*ProblemDetail) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{2}
}

func (x *ProblemDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProblemDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProblemDetail) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *ProblemDetail) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ProblemDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProblemDetail) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ProblemDetail) GetTimeLimit() uint64 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

func (x *ProblemDetail) GetSpaceLimit() uint64 {
	if x != nil {
		return x.SpaceLimit
	}
	return 0
}

type ProblemSynopsis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Point int32  `protobuf:"varint,3,opt,name=point,proto3" json:"point,omitempty"`
	Level int32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Tags  string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ProblemSynopsis) Reset() {
	*x = ProblemSynopsis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemSynopsis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemSynopsis) ProtoMessage() {}

func (x *ProblemSynopsis) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemSynopsis.ProtoReflect.Descriptor instead.
func (*ProblemSynopsis) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{3}
}

func (x *ProblemSynopsis) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProblemSynopsis) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProblemSynopsis) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *ProblemSynopsis) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ProblemSynopsis) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type ProblemPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32              `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Problems   []*ProblemSynopsis `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
}

func (x *ProblemPage) Reset() {
	*x = ProblemPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemPage) ProtoMessage() {}

func (x *ProblemPage) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemPage.ProtoReflect.Descriptor instead.
func (*ProblemPage) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{4}
}

func (x *ProblemPage) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ProblemPage) GetProblems() []*ProblemSynopsis {
	if x != nil {
		return x.Problems
	}
	return nil
}

type ProblemsByTagReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []int32 `protobuf:"varint,1,rep,packed,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ProblemsByTagReq) Reset() {
	*x = ProblemsByTagReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemsByTagReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemsByTagReq) ProtoMessage() {}

func (x *ProblemsByTagReq) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemsByTagReq.ProtoReflect.Descriptor instead.
func (*ProblemsByTagReq) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{5}
}

func (x *ProblemsByTagReq) GetTags() []int32 {
	if x != nil {
		return x.Tags
	}
	return nil
}

type ProblemIO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid    int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	InTxt  string `protobuf:"bytes,3,opt,name=inTxt,proto3" json:"inTxt,omitempty"`
	OutTxt string `protobuf:"bytes,4,opt,name=outTxt,proto3" json:"outTxt,omitempty"`
}

func (x *ProblemIO) Reset() {
	*x = ProblemIO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemIO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemIO) ProtoMessage() {}

func (x *ProblemIO) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemIO.ProtoReflect.Descriptor instead.
func (*ProblemIO) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{6}
}

func (x *ProblemIO) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProblemIO) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProblemIO) GetInTxt() string {
	if x != nil {
		return x.InTxt
	}
	return ""
}

func (x *ProblemIO) GetOutTxt() string {
	if x != nil {
		return x.OutTxt
	}
	return ""
}

type AddOrUpdateProblemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail *ProblemDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	Io     *ProblemIO     `protobuf:"bytes,2,opt,name=io,proto3" json:"io,omitempty"`
}

func (x *AddOrUpdateProblemReq) Reset() {
	*x = AddOrUpdateProblemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateProblemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateProblemReq) ProtoMessage() {}

func (x *AddOrUpdateProblemReq) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateProblemReq.ProtoReflect.Descriptor instead.
func (*AddOrUpdateProblemReq) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{7}
}

func (x *AddOrUpdateProblemReq) GetDetail() *ProblemDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *AddOrUpdateProblemReq) GetIo() *ProblemIO {
	if x != nil {
		return x.Io
	}
	return nil
}

type AllProblemStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Easy   int32 `protobuf:"varint,1,opt,name=easy,proto3" json:"easy,omitempty"`
	Medium int32 `protobuf:"varint,2,opt,name=medium,proto3" json:"medium,omitempty"`
	Hard   int32 `protobuf:"varint,3,opt,name=hard,proto3" json:"hard,omitempty"`
}

func (x *AllProblemStatistic) Reset() {
	*x = AllProblemStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProblemStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProblemStatistic) ProtoMessage() {}

func (x *AllProblemStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProblemStatistic.ProtoReflect.Descriptor instead.
func (*AllProblemStatistic) Descriptor() ([]byte, []int) {
	return file_problem_proto_rawDescGZIP(), []int{8}
}

func (x *AllProblemStatistic) GetEasy() int32 {
	if x != nil {
		return x.Easy
	}
	return 0
}

func (x *AllProblemStatistic) GetMedium() int32 {
	if x != nil {
		return x.Medium
	}
	return 0
}

func (x *AllProblemStatistic) GetHard() int32 {
	if x != nil {
		return x.Hard
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problem_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_problem_proto_msgTypes[9]
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
	return file_problem_proto_rawDescGZIP(), []int{9}
}

var File_problem_proto protoreflect.FileDescriptor

var file_problem_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x22,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x1f, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x53, 0x79, 0x6e, 0x6f, 0x70, 0x73, 0x69, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x22, 0x69, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x79, 0x6e, 0x6f, 0x70, 0x73, 0x69,
	0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x5b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x4f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x54, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6e, 0x54, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x54,
	0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x54, 0x78, 0x74,
	0x22, 0x77, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x28, 0x0a, 0x02, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x4f, 0x52, 0x02, 0x69, 0x6f, 0x22, 0x55, 0x0a, 0x13, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x61, 0x73, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x65, 0x61, 0x73, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x61, 0x72, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd0, 0x03, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x48, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x41, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x49, 0x4f, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x4f, 0x12, 0x50, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x73, 0x42, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x50, 0x61, 0x67, 0x65, 0x12, 0x52, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x50, 0x0a, 0x12, 0x61, 0x64,
	0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_problem_proto_rawDescOnce sync.Once
	file_problem_proto_rawDescData = file_problem_proto_rawDesc
)

func file_problem_proto_rawDescGZIP() []byte {
	file_problem_proto_rawDescOnce.Do(func() {
		file_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_problem_proto_rawDescData)
	})
	return file_problem_proto_rawDescData
}

var file_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_problem_proto_goTypes = []interface{}{
	(*Result)(nil),                // 0: problemclient.Result
	(*Integer)(nil),               // 1: problemclient.Integer
	(*ProblemDetail)(nil),         // 2: problemclient.ProblemDetail
	(*ProblemSynopsis)(nil),       // 3: problemclient.ProblemSynopsis
	(*ProblemPage)(nil),           // 4: problemclient.ProblemPage
	(*ProblemsByTagReq)(nil),      // 5: problemclient.ProblemsByTagReq
	(*ProblemIO)(nil),             // 6: problemclient.ProblemIO
	(*AddOrUpdateProblemReq)(nil), // 7: problemclient.AddOrUpdateProblemReq
	(*AllProblemStatistic)(nil),   // 8: problemclient.AllProblemStatistic
	(*Empty)(nil),                 // 9: problemclient.Empty
}
var file_problem_proto_depIdxs = []int32{
	3, // 0: problemclient.ProblemPage.problems:type_name -> problemclient.ProblemSynopsis
	2, // 1: problemclient.AddOrUpdateProblemReq.detail:type_name -> problemclient.ProblemDetail
	6, // 2: problemclient.AddOrUpdateProblemReq.io:type_name -> problemclient.ProblemIO
	1, // 3: problemclient.Problem.getProblemDetail:input_type -> problemclient.Integer
	9, // 4: problemclient.Problem.getAllProblem:input_type -> problemclient.Empty
	1, // 5: problemclient.Problem.getProblemIO:input_type -> problemclient.Integer
	5, // 6: problemclient.Problem.getProblemsByTags:input_type -> problemclient.ProblemsByTagReq
	9, // 7: problemclient.Problem.getAllProblemStatistic:input_type -> problemclient.Empty
	7, // 8: problemclient.Problem.addOrUpdateProblem:input_type -> problemclient.AddOrUpdateProblemReq
	2, // 9: problemclient.Problem.getProblemDetail:output_type -> problemclient.ProblemDetail
	4, // 10: problemclient.Problem.getAllProblem:output_type -> problemclient.ProblemPage
	6, // 11: problemclient.Problem.getProblemIO:output_type -> problemclient.ProblemIO
	4, // 12: problemclient.Problem.getProblemsByTags:output_type -> problemclient.ProblemPage
	8, // 13: problemclient.Problem.getAllProblemStatistic:output_type -> problemclient.AllProblemStatistic
	9, // 14: problemclient.Problem.addOrUpdateProblem:output_type -> problemclient.Empty
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_problem_proto_init() }
func file_problem_proto_init() {
	if File_problem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_problem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integer); i {
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
		file_problem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemDetail); i {
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
		file_problem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemSynopsis); i {
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
		file_problem_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemPage); i {
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
		file_problem_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemsByTagReq); i {
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
		file_problem_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemIO); i {
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
		file_problem_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateProblemReq); i {
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
		file_problem_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProblemStatistic); i {
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
		file_problem_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_problem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_problem_proto_goTypes,
		DependencyIndexes: file_problem_proto_depIdxs,
		MessageInfos:      file_problem_proto_msgTypes,
	}.Build()
	File_problem_proto = out.File
	file_problem_proto_rawDesc = nil
	file_problem_proto_goTypes = nil
	file_problem_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProblemClient is the client API for Problem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProblemClient interface {
	GetProblemDetail(ctx context.Context, in *Integer, opts ...grpc.CallOption) (*ProblemDetail, error)
	GetAllProblem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProblemPage, error)
	GetProblemIO(ctx context.Context, in *Integer, opts ...grpc.CallOption) (*ProblemIO, error)
	GetProblemsByTags(ctx context.Context, in *ProblemsByTagReq, opts ...grpc.CallOption) (*ProblemPage, error)
	GetAllProblemStatistic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllProblemStatistic, error)
	AddOrUpdateProblem(ctx context.Context, in *AddOrUpdateProblemReq, opts ...grpc.CallOption) (*Empty, error)
}

type problemClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemClient(cc grpc.ClientConnInterface) ProblemClient {
	return &problemClient{cc}
}

func (c *problemClient) GetProblemDetail(ctx context.Context, in *Integer, opts ...grpc.CallOption) (*ProblemDetail, error) {
	out := new(ProblemDetail)
	err := c.cc.Invoke(ctx, "/problemclient.Problem/getProblemDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetAllProblem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProblemPage, error) {
	out := new(ProblemPage)
	err := c.cc.Invoke(ctx, "/problemclient.Problem/getAllProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetProblemIO(ctx context.Context, in *Integer, opts ...grpc.CallOption) (*ProblemIO, error) {
	out := new(ProblemIO)
	err := c.cc.Invoke(ctx, "/problemclient.Problem/getProblemIO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetProblemsByTags(ctx context.Context, in *ProblemsByTagReq, opts ...grpc.CallOption) (*ProblemPage, error) {
	out := new(ProblemPage)
	err := c.cc.Invoke(ctx, "/problemclient.Problem/getProblemsByTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetAllProblemStatistic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllProblemStatistic, error) {
	out := new(AllProblemStatistic)
	err := c.cc.Invoke(ctx, "/problemclient.Problem/getAllProblemStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) AddOrUpdateProblem(ctx context.Context, in *AddOrUpdateProblemReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/problemclient.Problem/addOrUpdateProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServer is the server API for Problem service.
type ProblemServer interface {
	GetProblemDetail(context.Context, *Integer) (*ProblemDetail, error)
	GetAllProblem(context.Context, *Empty) (*ProblemPage, error)
	GetProblemIO(context.Context, *Integer) (*ProblemIO, error)
	GetProblemsByTags(context.Context, *ProblemsByTagReq) (*ProblemPage, error)
	GetAllProblemStatistic(context.Context, *Empty) (*AllProblemStatistic, error)
	AddOrUpdateProblem(context.Context, *AddOrUpdateProblemReq) (*Empty, error)
}

// UnimplementedProblemServer can be embedded to have forward compatible implementations.
type UnimplementedProblemServer struct {
}

func (*UnimplementedProblemServer) GetProblemDetail(context.Context, *Integer) (*ProblemDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemDetail not implemented")
}
func (*UnimplementedProblemServer) GetAllProblem(context.Context, *Empty) (*ProblemPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProblem not implemented")
}
func (*UnimplementedProblemServer) GetProblemIO(context.Context, *Integer) (*ProblemIO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemIO not implemented")
}
func (*UnimplementedProblemServer) GetProblemsByTags(context.Context, *ProblemsByTagReq) (*ProblemPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemsByTags not implemented")
}
func (*UnimplementedProblemServer) GetAllProblemStatistic(context.Context, *Empty) (*AllProblemStatistic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProblemStatistic not implemented")
}
func (*UnimplementedProblemServer) AddOrUpdateProblem(context.Context, *AddOrUpdateProblemReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateProblem not implemented")
}

func RegisterProblemServer(s *grpc.Server, srv ProblemServer) {
	s.RegisterService(&_Problem_serviceDesc, srv)
}

func _Problem_GetProblemDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Integer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblemDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/problemclient.Problem/GetProblemDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblemDetail(ctx, req.(*Integer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetAllProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetAllProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/problemclient.Problem/GetAllProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetAllProblem(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetProblemIO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Integer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblemIO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/problemclient.Problem/GetProblemIO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblemIO(ctx, req.(*Integer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetProblemsByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemsByTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblemsByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/problemclient.Problem/GetProblemsByTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblemsByTags(ctx, req.(*ProblemsByTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetAllProblemStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetAllProblemStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/problemclient.Problem/GetAllProblemStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetAllProblemStatistic(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_AddOrUpdateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateProblemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).AddOrUpdateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/problemclient.Problem/AddOrUpdateProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).AddOrUpdateProblem(ctx, req.(*AddOrUpdateProblemReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Problem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "problemclient.Problem",
	HandlerType: (*ProblemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getProblemDetail",
			Handler:    _Problem_GetProblemDetail_Handler,
		},
		{
			MethodName: "getAllProblem",
			Handler:    _Problem_GetAllProblem_Handler,
		},
		{
			MethodName: "getProblemIO",
			Handler:    _Problem_GetProblemIO_Handler,
		},
		{
			MethodName: "getProblemsByTags",
			Handler:    _Problem_GetProblemsByTags_Handler,
		},
		{
			MethodName: "getAllProblemStatistic",
			Handler:    _Problem_GetAllProblemStatistic_Handler,
		},
		{
			MethodName: "addOrUpdateProblem",
			Handler:    _Problem_AddOrUpdateProblem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "problem.proto",
}
