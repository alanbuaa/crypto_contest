// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: service.proto

package service

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

type RequestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestMsg) Reset() {
	*x = RequestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMsg) ProtoMessage() {}

func (x *RequestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMsg.ProtoReflect.Descriptor instead.
func (*RequestMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type ResponseMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResponseMsg) Reset() {
	*x = ResponseMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMsg) ProtoMessage() {}

func (x *ResponseMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMsg.ProtoReflect.Descriptor instead.
func (*ResponseMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

type PointMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	X       []byte `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y       []byte `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	Witness []byte `protobuf:"bytes,4,opt,name=witness,proto3" json:"witness,omitempty"`
}

func (x *PointMsg) Reset() {
	*x = PointMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointMsg) ProtoMessage() {}

func (x *PointMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointMsg.ProtoReflect.Descriptor instead.
func (*PointMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *PointMsg) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PointMsg) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *PointMsg) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *PointMsg) GetWitness() []byte {
	if x != nil {
		return x.Witness
	}
	return nil
}

type ZeroMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Share []byte `protobuf:"bytes,2,opt,name=share,proto3" json:"share,omitempty"`
}

func (x *ZeroMsg) Reset() {
	*x = ZeroMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZeroMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZeroMsg) ProtoMessage() {}

func (x *ZeroMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZeroMsg.ProtoReflect.Descriptor instead.
func (*ZeroMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ZeroMsg) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ZeroMsg) GetShare() []byte {
	if x != nil {
		return x.Share
	}
	return nil
}

type CommitMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       int32  `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	ShareCommit []byte `protobuf:"bytes,2,opt,name=ShareCommit,proto3" json:"ShareCommit,omitempty"`
	PolyCommit  []byte `protobuf:"bytes,3,opt,name=PolyCommit,proto3" json:"PolyCommit,omitempty"`
	ZeroWitness []byte `protobuf:"bytes,4,opt,name=ZeroWitness,proto3" json:"ZeroWitness,omitempty"`
}

func (x *CommitMsg) Reset() {
	*x = CommitMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitMsg) ProtoMessage() {}

func (x *CommitMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitMsg.ProtoReflect.Descriptor instead.
func (*CommitMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *CommitMsg) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CommitMsg) GetShareCommit() []byte {
	if x != nil {
		return x.ShareCommit
	}
	return nil
}

func (x *CommitMsg) GetPolyCommit() []byte {
	if x != nil {
		return x.PolyCommit
	}
	return nil
}

func (x *CommitMsg) GetZeroWitness() []byte {
	if x != nil {
		return x.ZeroWitness
	}
	return nil
}

type Cmt1Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index   int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Polycmt []byte `protobuf:"bytes,2,opt,name=polycmt,proto3" json:"polycmt,omitempty"`
}

func (x *Cmt1Msg) Reset() {
	*x = Cmt1Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmt1Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmt1Msg) ProtoMessage() {}

func (x *Cmt1Msg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmt1Msg.ProtoReflect.Descriptor instead.
func (*Cmt1Msg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *Cmt1Msg) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Cmt1Msg) GetPolycmt() []byte {
	if x != nil {
		return x.Polycmt
	}
	return nil
}

type RequestForCoeffMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label int32 `protobuf:"varint,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *RequestForCoeffMsg) Reset() {
	*x = RequestForCoeffMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestForCoeffMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestForCoeffMsg) ProtoMessage() {}

func (x *RequestForCoeffMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestForCoeffMsg.ProtoReflect.Descriptor instead.
func (*RequestForCoeffMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *RequestForCoeffMsg) GetLabel() int32 {
	if x != nil {
		return x.Label
	}
	return 0
}

type CoeffMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//    int32 label = 1;
	Coeff [][]byte `protobuf:"bytes,2,rep,name=coeff,proto3" json:"coeff,omitempty"`
}

func (x *CoeffMsg) Reset() {
	*x = CoeffMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoeffMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoeffMsg) ProtoMessage() {}

func (x *CoeffMsg) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoeffMsg.ProtoReflect.Descriptor instead.
func (*CoeffMsg) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *CoeffMsg) GetCoeff() [][]byte {
	if x != nil {
		return x.Coeff
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x56, 0x0a, 0x08, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x01, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x35, 0x0a,
	0x07, 0x5a, 0x65, 0x72, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f,
	0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x50, 0x6f, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x5a, 0x65,
	0x72, 0x6f, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x5a, 0x65, 0x72, 0x6f, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x39, 0x0a, 0x07,
	0x43, 0x6d, 0x74, 0x31, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x6f, 0x6c, 0x79, 0x63, 0x6d, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x6f, 0x6c, 0x79, 0x63, 0x6d, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x22, 0x20, 0x0a, 0x08, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x4d, 0x73, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05,
	0x63, 0x6f, 0x65, 0x66, 0x66, 0x32, 0xb7, 0x03, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x31, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x73, 0x67, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x50, 0x68, 0x61, 0x73, 0x65, 0x31, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73,
	0x67, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x50, 0x68, 0x61, 0x73, 0x65, 0x31, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0b, 0x50, 0x68, 0x61, 0x73, 0x65, 0x32, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x5a, 0x65, 0x72, 0x6f, 0x4d, 0x73,
	0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x50, 0x68, 0x61,
	0x73, 0x65, 0x32, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x50, 0x68, 0x61, 0x73, 0x65, 0x33,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x50, 0x68, 0x61, 0x73, 0x65, 0x33, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x32,
	0xae, 0x04, 0x0a, 0x14, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x65, 0x66, 0x66, 0x6f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x32, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x4d, 0x73, 0x67, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a,
	0x52, 0x65, 0x61, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x31, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x74, 0x31, 0x4d, 0x73,
	0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x31, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6d, 0x74, 0x31, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x32, 0x12, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x73,
	0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x52, 0x65, 0x61,
	0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x32, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x73, 0x67,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x68, 0x61,
	0x73, 0x65, 0x33, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d,
	0x74, 0x31, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x33, 0x32, 0x12, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x74, 0x31, 0x4d, 0x73, 0x67, 0x1a,
	0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x50,
	0x68, 0x61, 0x73, 0x65, 0x33, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x74, 0x31, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_proto_goTypes = []interface{}{
	(*RequestMsg)(nil),         // 0: service.RequestMsg
	(*ResponseMsg)(nil),        // 1: service.ResponseMsg
	(*PointMsg)(nil),           // 2: service.PointMsg
	(*ZeroMsg)(nil),            // 3: service.ZeroMsg
	(*CommitMsg)(nil),          // 4: service.CommitMsg
	(*Cmt1Msg)(nil),            // 5: service.Cmt1Msg
	(*RequestForCoeffMsg)(nil), // 6: service.RequestForCoeffMsg
	(*CoeffMsg)(nil),           // 7: service.CoeffMsg
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: service.NodeService.Phase1GetStart:input_type -> service.RequestMsg
	2,  // 1: service.NodeService.Phase1ReceiveMsg:input_type -> service.PointMsg
	0,  // 2: service.NodeService.Phase1Verify:input_type -> service.RequestMsg
	3,  // 3: service.NodeService.Phase2Share:input_type -> service.ZeroMsg
	0,  // 4: service.NodeService.Phase2Verify:input_type -> service.RequestMsg
	2,  // 5: service.NodeService.Phase3SendMsg:input_type -> service.PointMsg
	0,  // 6: service.NodeService.Phase3Verify:input_type -> service.RequestMsg
	0,  // 7: service.BulletinBoardService.GetCoeffofNodeSecretShares2:input_type -> service.RequestMsg
	0,  // 8: service.BulletinBoardService.StartEpoch:input_type -> service.RequestMsg
	0,  // 9: service.BulletinBoardService.ReadPhase1:input_type -> service.RequestMsg
	5,  // 10: service.BulletinBoardService.WritePhase1:input_type -> service.Cmt1Msg
	4,  // 11: service.BulletinBoardService.WritePhase2:input_type -> service.CommitMsg
	0,  // 12: service.BulletinBoardService.ReadPhase2:input_type -> service.RequestMsg
	5,  // 13: service.BulletinBoardService.WritePhase3:input_type -> service.Cmt1Msg
	5,  // 14: service.BulletinBoardService.WritePhase32:input_type -> service.Cmt1Msg
	0,  // 15: service.BulletinBoardService.ReadPhase3:input_type -> service.RequestMsg
	1,  // 16: service.NodeService.Phase1GetStart:output_type -> service.ResponseMsg
	1,  // 17: service.NodeService.Phase1ReceiveMsg:output_type -> service.ResponseMsg
	1,  // 18: service.NodeService.Phase1Verify:output_type -> service.ResponseMsg
	1,  // 19: service.NodeService.Phase2Share:output_type -> service.ResponseMsg
	1,  // 20: service.NodeService.Phase2Verify:output_type -> service.ResponseMsg
	1,  // 21: service.NodeService.Phase3SendMsg:output_type -> service.ResponseMsg
	1,  // 22: service.NodeService.Phase3Verify:output_type -> service.ResponseMsg
	7,  // 23: service.BulletinBoardService.GetCoeffofNodeSecretShares2:output_type -> service.CoeffMsg
	1,  // 24: service.BulletinBoardService.StartEpoch:output_type -> service.ResponseMsg
	5,  // 25: service.BulletinBoardService.ReadPhase1:output_type -> service.Cmt1Msg
	1,  // 26: service.BulletinBoardService.WritePhase1:output_type -> service.ResponseMsg
	1,  // 27: service.BulletinBoardService.WritePhase2:output_type -> service.ResponseMsg
	4,  // 28: service.BulletinBoardService.ReadPhase2:output_type -> service.CommitMsg
	1,  // 29: service.BulletinBoardService.WritePhase3:output_type -> service.ResponseMsg
	1,  // 30: service.BulletinBoardService.WritePhase32:output_type -> service.ResponseMsg
	5,  // 31: service.BulletinBoardService.ReadPhase3:output_type -> service.Cmt1Msg
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMsg); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMsg); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointMsg); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZeroMsg); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitMsg); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmt1Msg); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestForCoeffMsg); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoeffMsg); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
