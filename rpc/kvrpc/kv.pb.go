// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: kv.proto

package kvrpc

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

type GetInCasualRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Vectorclock map[string]int32 `protobuf:"bytes,2,rep,name=vectorclock,proto3" json:"vectorclock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Timestamp   int64            `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetInCasualRequest) Reset() {
	*x = GetInCasualRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInCasualRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInCasualRequest) ProtoMessage() {}

func (x *GetInCasualRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInCasualRequest.ProtoReflect.Descriptor instead.
func (*GetInCasualRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{0}
}

func (x *GetInCasualRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetInCasualRequest) GetVectorclock() map[string]int32 {
	if x != nil {
		return x.Vectorclock
	}
	return nil
}

func (x *GetInCasualRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetInCasualResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string           `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Vectorclock map[string]int32 `protobuf:"bytes,2,rep,name=vectorclock,proto3" json:"vectorclock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Success     bool             `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GetInCasualResponse) Reset() {
	*x = GetInCasualResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInCasualResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInCasualResponse) ProtoMessage() {}

func (x *GetInCasualResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInCasualResponse.ProtoReflect.Descriptor instead.
func (*GetInCasualResponse) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{1}
}

func (x *GetInCasualResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetInCasualResponse) GetVectorclock() map[string]int32 {
	if x != nil {
		return x.Vectorclock
	}
	return nil
}

func (x *GetInCasualResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PutInCasualRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value       string           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Vectorclock map[string]int32 `protobuf:"bytes,3,rep,name=vectorclock,proto3" json:"vectorclock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Timestamp   int64            `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PutInCasualRequest) Reset() {
	*x = PutInCasualRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutInCasualRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutInCasualRequest) ProtoMessage() {}

func (x *PutInCasualRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutInCasualRequest.ProtoReflect.Descriptor instead.
func (*PutInCasualRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{2}
}

func (x *PutInCasualRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutInCasualRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PutInCasualRequest) GetVectorclock() map[string]int32 {
	if x != nil {
		return x.Vectorclock
	}
	return nil
}

func (x *PutInCasualRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PutInCasualResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool             `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Vectorclock map[string]int32 `protobuf:"bytes,2,rep,name=vectorclock,proto3" json:"vectorclock,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PutInCasualResponse) Reset() {
	*x = PutInCasualResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutInCasualResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutInCasualResponse) ProtoMessage() {}

func (x *PutInCasualResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutInCasualResponse.ProtoReflect.Descriptor instead.
func (*PutInCasualResponse) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{3}
}

func (x *PutInCasualResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PutInCasualResponse) GetVectorclock() map[string]int32 {
	if x != nil {
		return x.Vectorclock
	}
	return nil
}

var File_kv_proto protoreflect.FileDescriptor

var file_kv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3e, 0x0a, 0x10, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xce, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe2, 0x01, 0x0a, 0x12, 0x50,
	0x75, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a,
	0x3e, 0x0a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb8, 0x01, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x47, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x43, 0x61,
	0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x3e, 0x0a, 0x10, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x7c, 0x0a, 0x02, 0x4b, 0x56,
	0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x12,
	0x13, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b,
	0x50, 0x75, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x12, 0x13, 0x2e, 0x50, 0x75,
	0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x50, 0x75, 0x74, 0x49, 0x6e, 0x43, 0x61, 0x73, 0x75, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x6b,
	0x76, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_proto_rawDescOnce sync.Once
	file_kv_proto_rawDescData = file_kv_proto_rawDesc
)

func file_kv_proto_rawDescGZIP() []byte {
	file_kv_proto_rawDescOnce.Do(func() {
		file_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_proto_rawDescData)
	})
	return file_kv_proto_rawDescData
}

var file_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_kv_proto_goTypes = []interface{}{
	(*GetInCasualRequest)(nil),  // 0: GetInCasualRequest
	(*GetInCasualResponse)(nil), // 1: GetInCasualResponse
	(*PutInCasualRequest)(nil),  // 2: PutInCasualRequest
	(*PutInCasualResponse)(nil), // 3: PutInCasualResponse
	nil,                         // 4: GetInCasualRequest.VectorclockEntry
	nil,                         // 5: GetInCasualResponse.VectorclockEntry
	nil,                         // 6: PutInCasualRequest.VectorclockEntry
	nil,                         // 7: PutInCasualResponse.VectorclockEntry
}
var file_kv_proto_depIdxs = []int32{
	4, // 0: GetInCasualRequest.vectorclock:type_name -> GetInCasualRequest.VectorclockEntry
	5, // 1: GetInCasualResponse.vectorclock:type_name -> GetInCasualResponse.VectorclockEntry
	6, // 2: PutInCasualRequest.vectorclock:type_name -> PutInCasualRequest.VectorclockEntry
	7, // 3: PutInCasualResponse.vectorclock:type_name -> PutInCasualResponse.VectorclockEntry
	0, // 4: KV.GetInCasual:input_type -> GetInCasualRequest
	2, // 5: KV.PutInCasual:input_type -> PutInCasualRequest
	1, // 6: KV.GetInCasual:output_type -> GetInCasualResponse
	3, // 7: KV.PutInCasual:output_type -> PutInCasualResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kv_proto_init() }
func file_kv_proto_init() {
	if File_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInCasualRequest); i {
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
		file_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInCasualResponse); i {
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
		file_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutInCasualRequest); i {
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
		file_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutInCasualResponse); i {
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
			RawDescriptor: file_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_proto_goTypes,
		DependencyIndexes: file_kv_proto_depIdxs,
		MessageInfos:      file_kv_proto_msgTypes,
	}.Build()
	File_kv_proto = out.File
	file_kv_proto_rawDesc = nil
	file_kv_proto_goTypes = nil
	file_kv_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	GetInCasual(ctx context.Context, in *GetInCasualRequest, opts ...grpc.CallOption) (*GetInCasualResponse, error)
	PutInCasual(ctx context.Context, in *PutInCasualRequest, opts ...grpc.CallOption) (*PutInCasualResponse, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) GetInCasual(ctx context.Context, in *GetInCasualRequest, opts ...grpc.CallOption) (*GetInCasualResponse, error) {
	out := new(GetInCasualResponse)
	err := c.cc.Invoke(ctx, "/KV/GetInCasual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) PutInCasual(ctx context.Context, in *PutInCasualRequest, opts ...grpc.CallOption) (*PutInCasualResponse, error) {
	out := new(PutInCasualResponse)
	err := c.cc.Invoke(ctx, "/KV/PutInCasual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	GetInCasual(context.Context, *GetInCasualRequest) (*GetInCasualResponse, error)
	PutInCasual(context.Context, *PutInCasualRequest) (*PutInCasualResponse, error)
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) GetInCasual(context.Context, *GetInCasualRequest) (*GetInCasualResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInCasual not implemented")
}
func (*UnimplementedKVServer) PutInCasual(context.Context, *PutInCasualRequest) (*PutInCasualResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutInCasual not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_GetInCasual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInCasualRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).GetInCasual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/GetInCasual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).GetInCasual(ctx, req.(*GetInCasualRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_PutInCasual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutInCasualRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).PutInCasual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/PutInCasual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).PutInCasual(ctx, req.(*PutInCasualRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInCasual",
			Handler:    _KV_GetInCasual_Handler,
		},
		{
			MethodName: "PutInCasual",
			Handler:    _KV_PutInCasual_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv.proto",
}
