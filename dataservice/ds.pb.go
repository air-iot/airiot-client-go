// Copyright 2015 The gRPC Authors
// protoc -I . --go_out=plugins=grpc:. ./ds.proto

// protoc -I ./ --go_out=. ./dataService/ds.proto
// protoc -I ./ --go-grpc_out=. dataService/ds.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: dataService/ds.proto

package dataservice

import (
	api "github.com/air-iot/api-client-go/v4/api"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_ds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_ds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_dataService_ds_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Request) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Code     int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Info     string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Detail   string `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	Result   []byte `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	Count    int64  `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	HttpCode int32  `protobuf:"varint,7,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	Headers  []byte `protobuf:"bytes,8,opt,name=headers,proto3" json:"headers,omitempty"`
}

func (x *ProxyResponse) Reset() {
	*x = ProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_ds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyResponse) ProtoMessage() {}

func (x *ProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_ds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyResponse.ProtoReflect.Descriptor instead.
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return file_dataService_ds_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ProxyResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ProxyResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *ProxyResponse) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *ProxyResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ProxyResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ProxyResponse) GetHttpCode() int32 {
	if x != nil {
		return x.HttpCode
	}
	return 0
}

func (x *ProxyResponse) GetHeaders() []byte {
	if x != nil {
		return x.Headers
	}
	return nil
}

type ViewPreviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode      string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	DatesetId string `protobuf:"bytes,2,opt,name=datesetId,proto3" json:"datesetId,omitempty"`
	ViewId    string `protobuf:"bytes,3,opt,name=viewId,proto3" json:"viewId,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ViewPreviewReq) Reset() {
	*x = ViewPreviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataService_ds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewPreviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewPreviewReq) ProtoMessage() {}

func (x *ViewPreviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_dataService_ds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewPreviewReq.ProtoReflect.Descriptor instead.
func (*ViewPreviewReq) Descriptor() ([]byte, []int) {
	return file_dataService_ds_proto_rawDescGZIP(), []int{2}
}

func (x *ViewPreviewReq) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *ViewPreviewReq) GetDatesetId() string {
	if x != nil {
		return x.DatesetId
	}
	return ""
}

func (x *ViewPreviewReq) GetViewId() string {
	if x != nil {
		return x.ViewId
	}
	return ""
}

func (x *ViewPreviewReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_dataService_ds_proto protoreflect.FileDescriptor

var file_dataService_ds_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xcc, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x22, 0x6e, 0x0a, 0x0e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x97, 0x03, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x9b, 0x03, 0x0a,
	0x14, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x6e, 0x79, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x4a, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xc2, 0x02, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x57,
	0x68, 0x6f, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x8a, 0x02, 0x0a, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x07, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4f, 0x0a, 0x2a, 0x63, 0x6e, 0x2e, 0x61,
	0x69, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x19, 0x2e, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dataService_ds_proto_rawDescOnce sync.Once
	file_dataService_ds_proto_rawDescData = file_dataService_ds_proto_rawDesc
)

func file_dataService_ds_proto_rawDescGZIP() []byte {
	file_dataService_ds_proto_rawDescOnce.Do(func() {
		file_dataService_ds_proto_rawDescData = protoimpl.X.CompressGZIP(file_dataService_ds_proto_rawDescData)
	})
	return file_dataService_ds_proto_rawDescData
}

var file_dataService_ds_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dataService_ds_proto_goTypes = []interface{}{
	(*Request)(nil),                // 0: dataService.Request
	(*ProxyResponse)(nil),          // 1: dataService.ProxyResponse
	(*ViewPreviewReq)(nil),         // 2: dataService.ViewPreviewReq
	(*api.QueryRequest)(nil),       // 3: api.QueryRequest
	(*api.GetOrDeleteRequest)(nil), // 4: api.GetOrDeleteRequest
	(*api.UpdateRequest)(nil),      // 5: api.UpdateRequest
	(*api.CreateRequest)(nil),      // 6: api.CreateRequest
	(*api.Response)(nil),           // 7: api.Response
}
var file_dataService_ds_proto_depIdxs = []int32{
	3,  // 0: dataService.DataGroupService.Query:input_type -> api.QueryRequest
	4,  // 1: dataService.DataGroupService.Get:input_type -> api.GetOrDeleteRequest
	4,  // 2: dataService.DataGroupService.Delete:input_type -> api.GetOrDeleteRequest
	5,  // 3: dataService.DataGroupService.Update:input_type -> api.UpdateRequest
	5,  // 4: dataService.DataGroupService.Replace:input_type -> api.UpdateRequest
	6,  // 5: dataService.DataGroupService.Create:input_type -> api.CreateRequest
	6,  // 6: dataService.DataGroupService.CreateMany:input_type -> api.CreateRequest
	3,  // 7: dataService.DataGroupService.DeleteMany:input_type -> api.QueryRequest
	3,  // 8: dataService.DataInterfaceService.Query:input_type -> api.QueryRequest
	4,  // 9: dataService.DataInterfaceService.Get:input_type -> api.GetOrDeleteRequest
	4,  // 10: dataService.DataInterfaceService.Delete:input_type -> api.GetOrDeleteRequest
	5,  // 11: dataService.DataInterfaceService.Update:input_type -> api.UpdateRequest
	5,  // 12: dataService.DataInterfaceService.Replace:input_type -> api.UpdateRequest
	6,  // 13: dataService.DataInterfaceService.Create:input_type -> api.CreateRequest
	6,  // 14: dataService.DataInterfaceService.CreateMany:input_type -> api.CreateRequest
	3,  // 15: dataService.DataInterfaceService.DeleteMany:input_type -> api.QueryRequest
	0,  // 16: dataService.DataService.Proxy:input_type -> dataService.Request
	3,  // 17: dataService.DatasetService.Query:input_type -> api.QueryRequest
	3,  // 18: dataService.DatasetService.QueryWhole:input_type -> api.QueryRequest
	6,  // 19: dataService.DatasetService.CreateWhole:input_type -> api.CreateRequest
	5,  // 20: dataService.DatasetService.ReplateWhole:input_type -> api.UpdateRequest
	2,  // 21: dataService.DatasetService.Preview:input_type -> dataService.ViewPreviewReq
	3,  // 22: dataService.DatasetService.DeleteAll:input_type -> api.QueryRequest
	3,  // 23: dataService.DatasetViewService.Query:input_type -> api.QueryRequest
	6,  // 24: dataService.DatasetViewService.Create:input_type -> api.CreateRequest
	5,  // 25: dataService.DatasetViewService.Replace:input_type -> api.UpdateRequest
	2,  // 26: dataService.DatasetViewService.Preview:input_type -> dataService.ViewPreviewReq
	3,  // 27: dataService.DatasetViewService.DeleteAll:input_type -> api.QueryRequest
	7,  // 28: dataService.DataGroupService.Query:output_type -> api.Response
	7,  // 29: dataService.DataGroupService.Get:output_type -> api.Response
	7,  // 30: dataService.DataGroupService.Delete:output_type -> api.Response
	7,  // 31: dataService.DataGroupService.Update:output_type -> api.Response
	7,  // 32: dataService.DataGroupService.Replace:output_type -> api.Response
	7,  // 33: dataService.DataGroupService.Create:output_type -> api.Response
	7,  // 34: dataService.DataGroupService.CreateMany:output_type -> api.Response
	7,  // 35: dataService.DataGroupService.DeleteMany:output_type -> api.Response
	7,  // 36: dataService.DataInterfaceService.Query:output_type -> api.Response
	7,  // 37: dataService.DataInterfaceService.Get:output_type -> api.Response
	7,  // 38: dataService.DataInterfaceService.Delete:output_type -> api.Response
	7,  // 39: dataService.DataInterfaceService.Update:output_type -> api.Response
	7,  // 40: dataService.DataInterfaceService.Replace:output_type -> api.Response
	7,  // 41: dataService.DataInterfaceService.Create:output_type -> api.Response
	7,  // 42: dataService.DataInterfaceService.CreateMany:output_type -> api.Response
	7,  // 43: dataService.DataInterfaceService.DeleteMany:output_type -> api.Response
	1,  // 44: dataService.DataService.Proxy:output_type -> dataService.ProxyResponse
	7,  // 45: dataService.DatasetService.Query:output_type -> api.Response
	7,  // 46: dataService.DatasetService.QueryWhole:output_type -> api.Response
	7,  // 47: dataService.DatasetService.CreateWhole:output_type -> api.Response
	7,  // 48: dataService.DatasetService.ReplateWhole:output_type -> api.Response
	7,  // 49: dataService.DatasetService.Preview:output_type -> api.Response
	7,  // 50: dataService.DatasetService.DeleteAll:output_type -> api.Response
	7,  // 51: dataService.DatasetViewService.Query:output_type -> api.Response
	7,  // 52: dataService.DatasetViewService.Create:output_type -> api.Response
	7,  // 53: dataService.DatasetViewService.Replace:output_type -> api.Response
	7,  // 54: dataService.DatasetViewService.Preview:output_type -> api.Response
	7,  // 55: dataService.DatasetViewService.DeleteAll:output_type -> api.Response
	28, // [28:56] is the sub-list for method output_type
	0,  // [0:28] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_dataService_ds_proto_init() }
func file_dataService_ds_proto_init() {
	if File_dataService_ds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dataService_ds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_dataService_ds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyResponse); i {
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
		file_dataService_ds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewPreviewReq); i {
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
			RawDescriptor: file_dataService_ds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_dataService_ds_proto_goTypes,
		DependencyIndexes: file_dataService_ds_proto_depIdxs,
		MessageInfos:      file_dataService_ds_proto_msgTypes,
	}.Build()
	File_dataService_ds_proto = out.File
	file_dataService_ds_proto_rawDesc = nil
	file_dataService_ds_proto_goTypes = nil
	file_dataService_ds_proto_depIdxs = nil
}
