// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/core.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core/core.proto

It has these top-level messages:
	GetDeviceRequest
	GetDataDeviceRequest
	GetRequestName
	QueryDataRequest
	GetOrDeleteDataRequest
	UpdateDataRequest
	MultiUpdateDataRequest
	CreateDataRequest
	LoginUserRequest
	UploadFileRequest
	DownloadFileResponse
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/air-iot/api-client-go/v4/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetDeviceRequest struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	Group  string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
}

func (m *GetDeviceRequest) Reset()                    { *m = GetDeviceRequest{} }
func (m *GetDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceRequest) ProtoMessage()               {}
func (*GetDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetDeviceRequest) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *GetDeviceRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type GetDataDeviceRequest struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	Group  string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	Table  string `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	Id     string `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
}

func (m *GetDataDeviceRequest) Reset()                    { *m = GetDataDeviceRequest{} }
func (m *GetDataDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDataDeviceRequest) ProtoMessage()               {}
func (*GetDataDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetDataDeviceRequest) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *GetDataDeviceRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *GetDataDeviceRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *GetDataDeviceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetRequestName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequestName) Reset()                    { *m = GetRequestName{} }
func (m *GetRequestName) String() string            { return proto.CompactTextString(m) }
func (*GetRequestName) ProtoMessage()               {}
func (*GetRequestName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRequestName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type QueryDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Query []byte `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (m *QueryDataRequest) Reset()                    { *m = QueryDataRequest{} }
func (m *QueryDataRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryDataRequest) ProtoMessage()               {}
func (*QueryDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *QueryDataRequest) GetQuery() []byte {
	if m != nil {
		return m.Query
	}
	return nil
}

type GetOrDeleteDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetOrDeleteDataRequest) Reset()                    { *m = GetOrDeleteDataRequest{} }
func (m *GetOrDeleteDataRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrDeleteDataRequest) ProtoMessage()               {}
func (*GetOrDeleteDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetOrDeleteDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *GetOrDeleteDataRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateDataRequest struct {
	Table        string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Id           string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Data         []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	CloseRequire bool   `protobuf:"varint,4,opt,name=closeRequire" json:"closeRequire,omitempty"`
}

func (m *UpdateDataRequest) Reset()                    { *m = UpdateDataRequest{} }
func (m *UpdateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDataRequest) ProtoMessage()               {}
func (*UpdateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UpdateDataRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdateDataRequest) GetCloseRequire() bool {
	if m != nil {
		return m.CloseRequire
	}
	return false
}

type MultiUpdateDataRequest struct {
	Table        string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Query        []byte `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Data         []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	CloseRequire bool   `protobuf:"varint,4,opt,name=closeRequire" json:"closeRequire,omitempty"`
}

func (m *MultiUpdateDataRequest) Reset()                    { *m = MultiUpdateDataRequest{} }
func (m *MultiUpdateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiUpdateDataRequest) ProtoMessage()               {}
func (*MultiUpdateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MultiUpdateDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *MultiUpdateDataRequest) GetQuery() []byte {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *MultiUpdateDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MultiUpdateDataRequest) GetCloseRequire() bool {
	if m != nil {
		return m.CloseRequire
	}
	return false
}

type CreateDataRequest struct {
	Table        string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Data         []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	CloseRequire bool   `protobuf:"varint,3,opt,name=closeRequire" json:"closeRequire,omitempty"`
}

func (m *CreateDataRequest) Reset()                    { *m = CreateDataRequest{} }
func (m *CreateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDataRequest) ProtoMessage()               {}
func (*CreateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *CreateDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateDataRequest) GetCloseRequire() bool {
	if m != nil {
		return m.CloseRequire
	}
	return false
}

type LoginUserRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginUserRequest) Reset()                    { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()               {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LoginUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UploadFileRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UploadFileRequest) Reset()                    { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()               {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UploadFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DownloadFileResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DownloadFileResponse) Reset()                    { *m = DownloadFileResponse{} }
func (m *DownloadFileResponse) String() string            { return proto.CompactTextString(m) }
func (*DownloadFileResponse) ProtoMessage()               {}
func (*DownloadFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DownloadFileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDeviceRequest)(nil), "core.GetDeviceRequest")
	proto.RegisterType((*GetDataDeviceRequest)(nil), "core.GetDataDeviceRequest")
	proto.RegisterType((*GetRequestName)(nil), "core.GetRequestName")
	proto.RegisterType((*QueryDataRequest)(nil), "core.QueryDataRequest")
	proto.RegisterType((*GetOrDeleteDataRequest)(nil), "core.GetOrDeleteDataRequest")
	proto.RegisterType((*UpdateDataRequest)(nil), "core.UpdateDataRequest")
	proto.RegisterType((*MultiUpdateDataRequest)(nil), "core.MultiUpdateDataRequest")
	proto.RegisterType((*CreateDataRequest)(nil), "core.CreateDataRequest")
	proto.RegisterType((*LoginUserRequest)(nil), "core.LoginUserRequest")
	proto.RegisterType((*UploadFileRequest)(nil), "core.UploadFileRequest")
	proto.RegisterType((*DownloadFileResponse)(nil), "core.DownloadFileResponse")
}

func init() { proto.RegisterFile("core/core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x59, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xee, 0x3a, 0x89, 0xdb, 0x3c, 0x27, 0xae, 0xb3, 0x84, 0x10, 0x05, 0x0e, 0x95, 0x55, 0x89,
	0x08, 0x5a, 0x27, 0x75, 0xd3, 0x96, 0x62, 0x5a, 0x1a, 0xdb, 0x89, 0x15, 0x91, 0x40, 0x70, 0x1a,
	0x90, 0xb8, 0x4d, 0x76, 0x1f, 0xce, 0xe0, 0xdd, 0x99, 0xcd, 0xee, 0xb8, 0x10, 0x89, 0x1f, 0xc0,
	0x5f, 0x80, 0x13, 0x9c, 0x8a, 0x38, 0x23, 0xf1, 0x03, 0x38, 0xf3, 0x9f, 0xd0, 0xcc, 0xec, 0xae,
	0xdd, 0x78, 0xd3, 0xce, 0xba, 0xa4, 0x07, 0x94, 0xcb, 0x6a, 0x67, 0x3c, 0xdf, 0x7c, 0x6f, 0xde,
	0x7c, 0xf3, 0xde, 0x9b, 0x35, 0x5c, 0x77, 0x78, 0x88, 0x6b, 0xf2, 0x51, 0x0b, 0x42, 0x2e, 0xb8,
	0x3d, 0x2d, 0xdf, 0x57, 0xe6, 0x49, 0x40, 0xd7, 0x48, 0x40, 0x75, 0x67, 0xf5, 0x09, 0x54, 0x3a,
	0x28, 0xda, 0xf8, 0x8c, 0x3a, 0xd8, 0xc5, 0x93, 0x01, 0x46, 0xc2, 0x5e, 0x82, 0xa2, 0x1b, 0xd2,
	0x67, 0x18, 0x2e, 0x5b, 0x37, 0xac, 0xd5, 0xd9, 0x6e, 0xdc, 0xb2, 0x17, 0x61, 0xa6, 0x17, 0xf2,
	0x41, 0xb0, 0x5c, 0x50, 0xdd, 0xba, 0x51, 0xfd, 0x0e, 0x16, 0xe5, 0x0c, 0x44, 0x90, 0xd7, 0x98,
	0x45, 0xf6, 0x0a, 0x72, 0xe4, 0xe1, 0xf2, 0x94, 0xee, 0x55, 0x0d, 0xbb, 0x0c, 0x05, 0xea, 0x2e,
	0x4f, 0xab, 0xae, 0x02, 0x75, 0xab, 0x37, 0xa1, 0xdc, 0x41, 0x11, 0x33, 0x7c, 0x4e, 0x7c, 0xb4,
	0x6d, 0x98, 0x66, 0xc4, 0xc7, 0x98, 0x43, 0xbd, 0x57, 0x1f, 0x43, 0xe5, 0xcb, 0x01, 0x86, 0xa7,
	0xd2, 0xa6, 0xc4, 0x9a, 0x74, 0x7e, 0x6b, 0x74, 0xfe, 0x45, 0x98, 0x39, 0x91, 0x23, 0x95, 0x2d,
	0x73, 0x5d, 0xdd, 0xa8, 0x3e, 0x86, 0xa5, 0x0e, 0x8a, 0x2f, 0xc2, 0x36, 0x7a, 0x28, 0xf0, 0xd5,
	0xb3, 0x68, 0x2b, 0x0b, 0xa9, 0x95, 0x27, 0xb0, 0x70, 0x18, 0xb8, 0x64, 0x02, 0xa8, 0x5c, 0x8e,
	0x4b, 0x04, 0x51, 0x5e, 0x98, 0xeb, 0xaa, 0x77, 0xbb, 0x0a, 0x73, 0x8e, 0xc7, 0x23, 0xe5, 0x58,
	0x1a, 0xa2, 0x72, 0xc7, 0xb5, 0xee, 0x0b, 0x7d, 0xd5, 0x1f, 0x61, 0x69, 0x6f, 0xe0, 0x09, 0x6a,
	0xca, 0x9b, 0xb9, 0xf0, 0x89, 0xd9, 0x09, 0x2c, 0xb4, 0x42, 0x34, 0x22, 0x4e, 0x28, 0x0a, 0x2f,
	0xa1, 0x98, 0xca, 0xa0, 0x58, 0x85, 0xca, 0x2e, 0xef, 0x51, 0x76, 0x18, 0x61, 0x38, 0xca, 0xc0,
	0xfb, 0xc8, 0x52, 0x06, 0xd9, 0xa8, 0xbe, 0x2f, 0xbd, 0xef, 0x71, 0xe2, 0x6e, 0x53, 0x2f, 0x15,
	0x63, 0x42, 0x6b, 0x0d, 0x69, 0xab, 0x1f, 0xc0, 0x62, 0x9b, 0x7f, 0xcf, 0x86, 0x43, 0xa3, 0x80,
	0xb3, 0x08, 0xb3, 0xc6, 0xd6, 0xff, 0x28, 0x40, 0x79, 0x97, 0x3a, 0xc8, 0x22, 0x3c, 0xc0, 0x50,
	0xca, 0xdc, 0xde, 0x50, 0x5a, 0x94, 0xc8, 0xf8, 0x07, 0x7b, 0xa1, 0x26, 0xcf, 0x95, 0x92, 0x5e,
	0xcc, 0xbb, 0x32, 0xaf, 0xba, 0x92, 0xa9, 0xab, 0x57, 0xec, 0x75, 0x80, 0xc3, 0x28, 0x17, 0xe2,
	0x1e, 0x5c, 0xdf, 0xa6, 0xcc, 0xdd, 0x23, 0xce, 0x31, 0x65, 0xd8, 0xe2, 0xae, 0x19, 0xec, 0x13,
	0x7d, 0xb0, 0xd5, 0x99, 0x4b, 0xe8, 0xde, 0x51, 0x83, 0x46, 0xb4, 0x7d, 0x2e, 0xba, 0x01, 0xf3,
	0xda, 0x89, 0x43, 0xa8, 0x8a, 0x24, 0x63, 0x9e, 0x1d, 0x83, 0xae, 0x5a, 0xf5, 0x5f, 0x2c, 0x80,
	0xcd, 0x20, 0x48, 0x1c, 0x55, 0x83, 0x6b, 0x1d, 0x14, 0x4f, 0xe5, 0xe6, 0xc4, 0x96, 0xab, 0xf7,
	0x73, 0xb9, 0x3f, 0x84, 0x19, 0xb5, 0x34, 0xa3, 0x65, 0xae, 0xc1, 0x54, 0x07, 0x85, 0xf9, 0xca,
	0xea, 0x3f, 0x5b, 0x00, 0xbb, 0xbc, 0x97, 0x18, 0x77, 0x1b, 0x8a, 0x5a, 0xba, 0xb6, 0xad, 0x46,
	0xea, 0xc6, 0x7f, 0x46, 0x97, 0x6b, 0x31, 0xf5, 0x7f, 0xa6, 0xa1, 0x24, 0x05, 0x9e, 0x18, 0x77,
	0xa1, 0x9e, 0xb0, 0xeb, 0x50, 0xd4, 0x23, 0x72, 0x60, 0x6e, 0x43, 0x51, 0x87, 0x98, 0xd8, 0x5d,
	0xba, 0x71, 0xee, 0xf0, 0x1a, 0x5c, 0xed, 0x62, 0xe0, 0x11, 0xc7, 0x70, 0x7c, 0xce, 0xdd, 0x78,
	0x04, 0x76, 0x07, 0x45, 0x6b, 0x10, 0x86, 0xc8, 0x84, 0x74, 0xdc, 0x0e, 0xfb, 0x96, 0xdb, 0x4b,
	0x5a, 0xaa, 0x67, 0xc3, 0x45, 0x96, 0xc8, 0x17, 0xe5, 0xef, 0xfb, 0x18, 0xfa, 0x34, 0x8a, 0x28,
	0x67, 0x2f, 0x2c, 0xed, 0x15, 0xdc, 0x77, 0xa0, 0xa4, 0xf6, 0xa3, 0x49, 0x9c, 0xfe, 0x20, 0x30,
	0xda, 0xa1, 0xfb, 0x50, 0xd1, 0xee, 0xdd, 0x23, 0x2c, 0x0f, 0xee, 0x01, 0x54, 0xb4, 0x31, 0x23,
	0x38, 0x13, 0x1b, 0xeb, 0xbf, 0xcf, 0x80, 0xfd, 0x54, 0x86, 0xdd, 0x03, 0xe7, 0x18, 0x7d, 0x32,
	0x91, 0xac, 0xb6, 0xe1, 0x5d, 0x9d, 0x4c, 0x55, 0x72, 0x6f, 0x9e, 0xea, 0x88, 0xb2, 0xc9, 0xdc,
	0x8e, 0xca, 0xdb, 0xb1, 0xb3, 0xcf, 0xd6, 0x10, 0xe3, 0xf3, 0x7c, 0x06, 0x37, 0xd4, 0x3c, 0xca,
	0x9e, 0xd7, 0x9d, 0xec, 0x21, 0x80, 0x8c, 0x89, 0x7a, 0x94, 0xbd, 0x32, 0x84, 0x9d, 0xad, 0x42,
	0xc6, 0xa1, 0x77, 0x61, 0x5e, 0xd9, 0xb1, 0xe5, 0x0f, 0x3c, 0x22, 0x78, 0x78, 0x79, 0xb6, 0xce,
	0x4d, 0x54, 0x07, 0x82, 0x88, 0xc8, 0x5c, 0x29, 0x0d, 0x9d, 0xa8, 0x5a, 0xdc, 0xf7, 0x09, 0x73,
	0x9b, 0xa7, 0x3b, 0xed, 0x1c, 0x61, 0xf9, 0x6f, 0x80, 0x8a, 0x96, 0x06, 0x11, 0xa9, 0x50, 0xd7,
	0x13, 0xa1, 0xc6, 0xc2, 0x38, 0x5b, 0xd5, 0x65, 0x09, 0xa3, 0xac, 0x4f, 0xa5, 0xd6, 0xd9, 0x8e,
	0x6b, 0x0e, 0xdd, 0xd0, 0x7b, 0xfc, 0x5e, 0x2a, 0xa6, 0x8c, 0x02, 0x30, 0xeb, 0x6c, 0x26, 0x1b,
	0x9d, 0x13, 0x78, 0x0f, 0x60, 0x18, 0x0c, 0xcc, 0xad, 0xac, 0xa7, 0x22, 0x49, 0x33, 0xf2, 0x99,
	0x8a, 0x2f, 0x4b, 0xf2, 0xa9, 0x52, 0xcc, 0x41, 0xf5, 0x54, 0x2e, 0x31, 0x66, 0xac, 0xc2, 0xcb,
	0x0a, 0x70, 0x30, 0x0c, 0x54, 0x39, 0x70, 0x9b, 0xb0, 0x22, 0x95, 0x93, 0xee, 0x7f, 0x1b, 0x03,
	0xd1, 0x3c, 0x95, 0xcf, 0x9d, 0x76, 0x64, 0x26, 0xd7, 0x06, 0x80, 0x5e, 0x94, 0xa2, 0x8e, 0xf7,
	0x22, 0xbb, 0x24, 0xce, 0x52, 0xcd, 0x5b, 0x1d, 0x14, 0x5f, 0x93, 0x90, 0x51, 0xd6, 0xdb, 0xa6,
	0x9e, 0xc0, 0x50, 0x12, 0x6b, 0xd1, 0x6f, 0xf9, 0x81, 0x78, 0x59, 0x99, 0x55, 0xd2, 0xa6, 0xf7,
	0x94, 0xe0, 0x73, 0x8a, 0xe0, 0xc9, 0xf8, 0x91, 0xc9, 0x39, 0xc3, 0x06, 0xcc, 0xc6, 0x82, 0x6f,
	0x37, 0xcd, 0x55, 0xf4, 0x11, 0x5c, 0xed, 0xa0, 0x50, 0x98, 0x9c, 0x7c, 0x8d, 0x44, 0xb6, 0x93,
	0x80, 0x1f, 0x42, 0x79, 0x24, 0x01, 0xe6, 0xb2, 0xf8, 0x7e, 0xb2, 0xbf, 0x0a, 0x66, 0x2e, 0xe3,
	0x07, 0x50, 0x8a, 0xb5, 0x9f, 0x13, 0x98, 0x6a, 0x79, 0x14, 0x67, 0xa0, 0xe5, 0x8f, 0xa1, 0x3c,
	0x92, 0xac, 0xf3, 0x61, 0x3f, 0x85, 0xf2, 0x50, 0xc4, 0xa3, 0x0e, 0x36, 0x14, 0x72, 0xfd, 0xcf,
	0x42, 0x9c, 0xf0, 0xbb, 0xe8, 0xf0, 0xd0, 0xbd, 0xac, 0x23, 0xcd, 0xea, 0x24, 0x0f, 0xca, 0x7b,
	0x18, 0x45, 0xa4, 0x87, 0x13, 0x5e, 0x0b, 0x72, 0x55, 0xf9, 0x0c, 0x2a, 0x6d, 0x12, 0x1d, 0x1f,
	0x71, 0x32, 0xdc, 0xa1, 0x8b, 0xe4, 0x7b, 0x6e, 0x49, 0x42, 0x41, 0xd4, 0xa8, 0x84, 0xf0, 0x0e,
	0xc0, 0x3e, 0x8f, 0xc4, 0x2e, 0x11, 0xea, 0x7a, 0x6c, 0x42, 0xaa, 0xef, 0x71, 0xe6, 0x42, 0x5a,
	0x87, 0x59, 0x49, 0xa1, 0x01, 0x46, 0xfb, 0xf0, 0x9b, 0x05, 0xa5, 0x2e, 0xf7, 0xf0, 0xcd, 0xe8,
	0x76, 0x03, 0xca, 0x9b, 0xae, 0x4f, 0x99, 0x64, 0x6c, 0x1d, 0xa3, 0xd3, 0x37, 0x09, 0xf8, 0x75,
	0x06, 0xe5, 0x16, 0x11, 0xc4, 0x1b, 0x5e, 0x21, 0x2f, 0xf6, 0xbe, 0xda, 0x87, 0x92, 0x4c, 0x84,
	0x6f, 0x86, 0xec, 0x11, 0x94, 0x0f, 0x50, 0x08, 0xca, 0x26, 0x5a, 0x5c, 0xfd, 0xaf, 0x02, 0xbc,
	0x7d, 0x70, 0x1a, 0x09, 0xf4, 0xbf, 0x22, 0x21, 0x55, 0x17, 0x8f, 0xcb, 0x08, 0x64, 0xa4, 0xfc,
	0xe7, 0x53, 0x30, 0xaf, 0x6f, 0x76, 0xff, 0x1b, 0x8f, 0xdd, 0x82, 0xe2, 0x8e, 0x1f, 0xf0, 0x50,
	0x18, 0xad, 0xe0, 0x16, 0x14, 0xb7, 0x7e, 0x30, 0x1e, 0xbd, 0x21, 0x4d, 0xf1, 0x38, 0x71, 0xf3,
	0x7c, 0x97, 0xb2, 0x5b, 0x70, 0x2d, 0xf9, 0xe0, 0x77, 0xfe, 0xb2, 0xe3, 0xcb, 0x64, 0xd6, 0x97,
	0xc1, 0xea, 0x95, 0x75, 0x2b, 0x4e, 0xb1, 0x51, 0x7f, 0x8f, 0x30, 0xd2, 0xbb, 0xfc, 0x54, 0x63,
	0x28, 0xf0, 0x66, 0x03, 0x6e, 0x3a, 0xac, 0x46, 0x68, 0x48, 0xb9, 0xa8, 0x45, 0x6e, 0xbf, 0xe6,
	0x78, 0x14, 0x99, 0xa8, 0xb9, 0x83, 0xa3, 0x23, 0x5e, 0xeb, 0x85, 0x81, 0xa3, 0xfc, 0xbe, 0x6f,
	0x7d, 0x53, 0xaa, 0xa9, 0x3f, 0x2d, 0x1a, 0xf2, 0xf1, 0x93, 0x65, 0xfd, 0x6a, 0x59, 0x47, 0x45,
	0xf5, 0x5f, 0xc5, 0xdd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x36, 0x17, 0xab, 0x02, 0xd3, 0x18,
	0x00, 0x00,
}
