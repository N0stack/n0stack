// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/provisioning/block_storage_agent.proto

package provisioning // import "github.com/n0stack/n0stack/n0core/pkg/api/provisioning"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateEmptyBlockStorageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmptyBlockStorageRequest) Reset()         { *m = CreateEmptyBlockStorageRequest{} }
func (m *CreateEmptyBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEmptyBlockStorageRequest) ProtoMessage()    {}
func (*CreateEmptyBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_storage_agent_358f857d83c2f86f, []int{0}
}
func (m *CreateEmptyBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmptyBlockStorageRequest.Unmarshal(m, b)
}
func (m *CreateEmptyBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmptyBlockStorageRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEmptyBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmptyBlockStorageRequest.Merge(dst, src)
}
func (m *CreateEmptyBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEmptyBlockStorageRequest.Size(m)
}
func (m *CreateEmptyBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmptyBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmptyBlockStorageRequest proto.InternalMessageInfo

func (m *CreateEmptyBlockStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEmptyBlockStorageRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type CreateEmptyBlockStorageResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmptyBlockStorageResponse) Reset()         { *m = CreateEmptyBlockStorageResponse{} }
func (m *CreateEmptyBlockStorageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEmptyBlockStorageResponse) ProtoMessage()    {}
func (*CreateEmptyBlockStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_storage_agent_358f857d83c2f86f, []int{1}
}
func (m *CreateEmptyBlockStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmptyBlockStorageResponse.Unmarshal(m, b)
}
func (m *CreateEmptyBlockStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmptyBlockStorageResponse.Marshal(b, m, deterministic)
}
func (dst *CreateEmptyBlockStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmptyBlockStorageResponse.Merge(dst, src)
}
func (m *CreateEmptyBlockStorageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEmptyBlockStorageResponse.Size(m)
}
func (m *CreateEmptyBlockStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmptyBlockStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmptyBlockStorageResponse proto.InternalMessageInfo

func (m *CreateEmptyBlockStorageResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEmptyBlockStorageResponse) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *CreateEmptyBlockStorageResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type FetchBlockStorageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	SourceUrl            string   `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl" json:"source_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockStorageRequest) Reset()         { *m = FetchBlockStorageRequest{} }
func (m *FetchBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*FetchBlockStorageRequest) ProtoMessage()    {}
func (*FetchBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_storage_agent_358f857d83c2f86f, []int{2}
}
func (m *FetchBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchBlockStorageRequest.Unmarshal(m, b)
}
func (m *FetchBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchBlockStorageRequest.Marshal(b, m, deterministic)
}
func (dst *FetchBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockStorageRequest.Merge(dst, src)
}
func (m *FetchBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_FetchBlockStorageRequest.Size(m)
}
func (m *FetchBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockStorageRequest proto.InternalMessageInfo

func (m *FetchBlockStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FetchBlockStorageRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *FetchBlockStorageRequest) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

type FetchBlockStorageResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockStorageResponse) Reset()         { *m = FetchBlockStorageResponse{} }
func (m *FetchBlockStorageResponse) String() string { return proto.CompactTextString(m) }
func (*FetchBlockStorageResponse) ProtoMessage()    {}
func (*FetchBlockStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_storage_agent_358f857d83c2f86f, []int{3}
}
func (m *FetchBlockStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchBlockStorageResponse.Unmarshal(m, b)
}
func (m *FetchBlockStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchBlockStorageResponse.Marshal(b, m, deterministic)
}
func (dst *FetchBlockStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockStorageResponse.Merge(dst, src)
}
func (m *FetchBlockStorageResponse) XXX_Size() int {
	return xxx_messageInfo_FetchBlockStorageResponse.Size(m)
}
func (m *FetchBlockStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockStorageResponse proto.InternalMessageInfo

func (m *FetchBlockStorageResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FetchBlockStorageResponse) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *FetchBlockStorageResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type DeleteBlockStorageRequest struct {
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlockStorageRequest) Reset()         { *m = DeleteBlockStorageRequest{} }
func (m *DeleteBlockStorageRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockStorageRequest) ProtoMessage()    {}
func (*DeleteBlockStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_storage_agent_358f857d83c2f86f, []int{4}
}
func (m *DeleteBlockStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlockStorageRequest.Unmarshal(m, b)
}
func (m *DeleteBlockStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlockStorageRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBlockStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlockStorageRequest.Merge(dst, src)
}
func (m *DeleteBlockStorageRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlockStorageRequest.Size(m)
}
func (m *DeleteBlockStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlockStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlockStorageRequest proto.InternalMessageInfo

func (m *DeleteBlockStorageRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateEmptyBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.CreateEmptyBlockStorageRequest")
	proto.RegisterType((*CreateEmptyBlockStorageResponse)(nil), "n0stack.internal.n0core.provisioning.CreateEmptyBlockStorageResponse")
	proto.RegisterType((*FetchBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.FetchBlockStorageRequest")
	proto.RegisterType((*FetchBlockStorageResponse)(nil), "n0stack.internal.n0core.provisioning.FetchBlockStorageResponse")
	proto.RegisterType((*DeleteBlockStorageRequest)(nil), "n0stack.internal.n0core.provisioning.DeleteBlockStorageRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlockStorageAgentService service

type BlockStorageAgentServiceClient interface {
	CreateEmptyBlockStorage(ctx context.Context, in *CreateEmptyBlockStorageRequest, opts ...grpc.CallOption) (*CreateEmptyBlockStorageResponse, error)
	FetchBlockStorage(ctx context.Context, in *FetchBlockStorageRequest, opts ...grpc.CallOption) (*FetchBlockStorageResponse, error)
	DeleteBlockStorage(ctx context.Context, in *DeleteBlockStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type blockStorageAgentServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlockStorageAgentServiceClient(cc *grpc.ClientConn) BlockStorageAgentServiceClient {
	return &blockStorageAgentServiceClient{cc}
}

func (c *blockStorageAgentServiceClient) CreateEmptyBlockStorage(ctx context.Context, in *CreateEmptyBlockStorageRequest, opts ...grpc.CallOption) (*CreateEmptyBlockStorageResponse, error) {
	out := new(CreateEmptyBlockStorageResponse)
	err := grpc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.BlockStorageAgentService/CreateEmptyBlockStorage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStorageAgentServiceClient) FetchBlockStorage(ctx context.Context, in *FetchBlockStorageRequest, opts ...grpc.CallOption) (*FetchBlockStorageResponse, error) {
	out := new(FetchBlockStorageResponse)
	err := grpc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.BlockStorageAgentService/FetchBlockStorage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStorageAgentServiceClient) DeleteBlockStorage(ctx context.Context, in *DeleteBlockStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.BlockStorageAgentService/DeleteBlockStorage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlockStorageAgentService service

type BlockStorageAgentServiceServer interface {
	CreateEmptyBlockStorage(context.Context, *CreateEmptyBlockStorageRequest) (*CreateEmptyBlockStorageResponse, error)
	FetchBlockStorage(context.Context, *FetchBlockStorageRequest) (*FetchBlockStorageResponse, error)
	DeleteBlockStorage(context.Context, *DeleteBlockStorageRequest) (*empty.Empty, error)
}

func RegisterBlockStorageAgentServiceServer(s *grpc.Server, srv BlockStorageAgentServiceServer) {
	s.RegisterService(&_BlockStorageAgentService_serviceDesc, srv)
}

func _BlockStorageAgentService_CreateEmptyBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmptyBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).CreateEmptyBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.BlockStorageAgentService/CreateEmptyBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).CreateEmptyBlockStorage(ctx, req.(*CreateEmptyBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStorageAgentService_FetchBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).FetchBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.BlockStorageAgentService/FetchBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).FetchBlockStorage(ctx, req.(*FetchBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStorageAgentService_DeleteBlockStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStorageAgentServiceServer).DeleteBlockStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.BlockStorageAgentService/DeleteBlockStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStorageAgentServiceServer).DeleteBlockStorage(ctx, req.(*DeleteBlockStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockStorageAgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.internal.n0core.provisioning.BlockStorageAgentService",
	HandlerType: (*BlockStorageAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmptyBlockStorage",
			Handler:    _BlockStorageAgentService_CreateEmptyBlockStorage_Handler,
		},
		{
			MethodName: "FetchBlockStorage",
			Handler:    _BlockStorageAgentService_FetchBlockStorage_Handler,
		},
		{
			MethodName: "DeleteBlockStorage",
			Handler:    _BlockStorageAgentService_DeleteBlockStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/provisioning/block_storage_agent.proto",
}

func init() {
	proto.RegisterFile("pkg/api/provisioning/block_storage_agent.proto", fileDescriptor_block_storage_agent_358f857d83c2f86f)
}

var fileDescriptor_block_storage_agent_358f857d83c2f86f = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x9b, 0xaf, 0xfd, 0x90, 0xea, 0x0d, 0x0b, 0x41, 0x1a, 0x04, 0x54, 0x11, 0x43, 0x27,
	0xbb, 0x82, 0x11, 0x09, 0x44, 0xff, 0x30, 0x30, 0xb6, 0x62, 0x80, 0x25, 0x72, 0xac, 0x4b, 0x1a,
	0x35, 0xb5, 0x83, 0xed, 0x54, 0xea, 0x8b, 0xf0, 0x0c, 0xbc, 0x0d, 0xaf, 0x84, 0x12, 0x17, 0x14,
	0xd1, 0x06, 0x45, 0x94, 0x29, 0xb6, 0x75, 0xef, 0xef, 0xdc, 0x9c, 0x13, 0x07, 0x91, 0x74, 0x1e,
	0x51, 0x96, 0xc6, 0x34, 0x55, 0x72, 0x19, 0xeb, 0x58, 0x8a, 0x58, 0x44, 0x34, 0x4c, 0x24, 0x9f,
	0x07, 0xda, 0x48, 0xc5, 0x22, 0x08, 0x58, 0x04, 0xc2, 0x90, 0x54, 0x49, 0x23, 0xf1, 0xb9, 0xe8,
	0x6b, 0xc3, 0xf8, 0x9c, 0xc4, 0xc2, 0x80, 0x12, 0x2c, 0x21, 0xa2, 0xcf, 0xa5, 0x02, 0x52, 0xee,
	0xf7, 0x8e, 0x23, 0x29, 0xa3, 0x04, 0x68, 0xd1, 0x13, 0x66, 0xcf, 0x14, 0x16, 0xa9, 0x59, 0x59,
	0x84, 0x7f, 0x8f, 0x4e, 0x87, 0x0a, 0x98, 0x81, 0x71, 0x7e, 0x38, 0xc8, 0xa5, 0xa6, 0x56, 0x69,
	0x02, 0x2f, 0x19, 0x68, 0x83, 0x31, 0x6a, 0x09, 0xb6, 0x00, 0xd7, 0xe9, 0x3a, 0xbd, 0xf6, 0xa4,
	0x58, 0xe3, 0x03, 0xf4, 0x3f, 0x5c, 0x19, 0xd0, 0xee, 0xbf, 0xae, 0xd3, 0x6b, 0x4d, 0xec, 0xc6,
	0x0f, 0xd0, 0x59, 0x25, 0x4b, 0xa7, 0x52, 0x68, 0xa8, 0x0f, 0xcb, 0x2b, 0x53, 0x66, 0x66, 0x6e,
	0xd3, 0x56, 0xe6, 0x6b, 0x9f, 0x23, 0xf7, 0x0e, 0x0c, 0x9f, 0xed, 0x34, 0x26, 0x3e, 0x41, 0x48,
	0xcb, 0x4c, 0x71, 0x08, 0x32, 0x95, 0xac, 0xf9, 0x6d, 0x7b, 0xf2, 0xa0, 0x12, 0xff, 0x11, 0x75,
	0xb6, 0x88, 0xfc, 0xc9, 0xfc, 0x14, 0x75, 0x46, 0x90, 0x80, 0x81, 0x8a, 0x17, 0xf8, 0xde, 0x70,
	0xf1, 0xde, 0x44, 0x6e, 0xb9, 0xf6, 0x36, 0x0f, 0x7f, 0x0a, 0x6a, 0x19, 0x73, 0xc0, 0x6f, 0x0e,
	0x3a, 0xaa, 0xf0, 0x1b, 0x8f, 0x48, 0x9d, 0x4f, 0x83, 0xfc, 0x1c, 0xbd, 0x37, 0xde, 0x91, 0x62,
	0x4d, 0xf3, 0x1b, 0xf8, 0xd5, 0x41, 0xfb, 0x1b, 0xa6, 0xe2, 0xeb, 0x7a, 0xf8, 0xaa, 0xc8, 0xbd,
	0x9b, 0x5f, 0xf7, 0x7f, 0x0d, 0x26, 0x11, 0xde, 0x4c, 0x04, 0xd7, 0x04, 0x57, 0x66, 0xe9, 0x1d,
	0x12, 0x7b, 0xe7, 0xc8, 0xe7, 0x9d, 0x23, 0x85, 0x3b, 0x7e, 0x63, 0x30, 0x7e, 0x1a, 0x46, 0xb1,
	0x99, 0x65, 0x21, 0xe1, 0x72, 0x41, 0xd7, 0x32, 0xa5, 0x67, 0xae, 0x42, 0xb7, 0xfd, 0x06, 0xae,
	0xca, 0x9b, 0x70, 0xaf, 0x00, 0x5f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x04, 0x5c, 0xca, 0x4e,
	0x32, 0x04, 0x00, 0x00,
}
