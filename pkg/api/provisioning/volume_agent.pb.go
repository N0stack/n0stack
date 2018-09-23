// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/provisioning/volume_agent.proto

package provisioning // import "github.com/n0stack/n0core/pkg/api/provisioning"

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

type VolumeAgent struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeAgent) Reset()         { *m = VolumeAgent{} }
func (m *VolumeAgent) String() string { return proto.CompactTextString(m) }
func (*VolumeAgent) ProtoMessage()    {}
func (*VolumeAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_agent_4c16863b3472c44e, []int{0}
}
func (m *VolumeAgent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAgent.Unmarshal(m, b)
}
func (m *VolumeAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAgent.Marshal(b, m, deterministic)
}
func (dst *VolumeAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAgent.Merge(dst, src)
}
func (m *VolumeAgent) XXX_Size() int {
	return xxx_messageInfo_VolumeAgent.Size(m)
}
func (m *VolumeAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAgent.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAgent proto.InternalMessageInfo

func (m *VolumeAgent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VolumeAgent) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *VolumeAgent) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type CreateEmptyVolumeAgentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmptyVolumeAgentRequest) Reset()         { *m = CreateEmptyVolumeAgentRequest{} }
func (m *CreateEmptyVolumeAgentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEmptyVolumeAgentRequest) ProtoMessage()    {}
func (*CreateEmptyVolumeAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_agent_4c16863b3472c44e, []int{1}
}
func (m *CreateEmptyVolumeAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmptyVolumeAgentRequest.Unmarshal(m, b)
}
func (m *CreateEmptyVolumeAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmptyVolumeAgentRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEmptyVolumeAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmptyVolumeAgentRequest.Merge(dst, src)
}
func (m *CreateEmptyVolumeAgentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEmptyVolumeAgentRequest.Size(m)
}
func (m *CreateEmptyVolumeAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmptyVolumeAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmptyVolumeAgentRequest proto.InternalMessageInfo

func (m *CreateEmptyVolumeAgentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEmptyVolumeAgentRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type CreateVolumeAgentWithDownloadingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	SourceUrl            string   `protobuf:"bytes,4,opt,name=source_url,json=sourceUrl" json:"source_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVolumeAgentWithDownloadingRequest) Reset() {
	*m = CreateVolumeAgentWithDownloadingRequest{}
}
func (m *CreateVolumeAgentWithDownloadingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVolumeAgentWithDownloadingRequest) ProtoMessage()    {}
func (*CreateVolumeAgentWithDownloadingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_agent_4c16863b3472c44e, []int{2}
}
func (m *CreateVolumeAgentWithDownloadingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVolumeAgentWithDownloadingRequest.Unmarshal(m, b)
}
func (m *CreateVolumeAgentWithDownloadingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVolumeAgentWithDownloadingRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVolumeAgentWithDownloadingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVolumeAgentWithDownloadingRequest.Merge(dst, src)
}
func (m *CreateVolumeAgentWithDownloadingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVolumeAgentWithDownloadingRequest.Size(m)
}
func (m *CreateVolumeAgentWithDownloadingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVolumeAgentWithDownloadingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVolumeAgentWithDownloadingRequest proto.InternalMessageInfo

func (m *CreateVolumeAgentWithDownloadingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVolumeAgentWithDownloadingRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *CreateVolumeAgentWithDownloadingRequest) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

type DeleteVolumeAgentRequest struct {
	Path                 string   `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteVolumeAgentRequest) Reset()         { *m = DeleteVolumeAgentRequest{} }
func (m *DeleteVolumeAgentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVolumeAgentRequest) ProtoMessage()    {}
func (*DeleteVolumeAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_agent_4c16863b3472c44e, []int{3}
}
func (m *DeleteVolumeAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVolumeAgentRequest.Unmarshal(m, b)
}
func (m *DeleteVolumeAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVolumeAgentRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteVolumeAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVolumeAgentRequest.Merge(dst, src)
}
func (m *DeleteVolumeAgentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteVolumeAgentRequest.Size(m)
}
func (m *DeleteVolumeAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVolumeAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVolumeAgentRequest proto.InternalMessageInfo

func (m *DeleteVolumeAgentRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*VolumeAgent)(nil), "n0stack.internal.n0core.provisioning.VolumeAgent")
	proto.RegisterType((*CreateEmptyVolumeAgentRequest)(nil), "n0stack.internal.n0core.provisioning.CreateEmptyVolumeAgentRequest")
	proto.RegisterType((*CreateVolumeAgentWithDownloadingRequest)(nil), "n0stack.internal.n0core.provisioning.CreateVolumeAgentWithDownloadingRequest")
	proto.RegisterType((*DeleteVolumeAgentRequest)(nil), "n0stack.internal.n0core.provisioning.DeleteVolumeAgentRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VolumeAgentService service

type VolumeAgentServiceClient interface {
	CreateEmptyVolumeAgent(ctx context.Context, in *CreateEmptyVolumeAgentRequest, opts ...grpc.CallOption) (*VolumeAgent, error)
	CreateVolumeAgentWithDownloading(ctx context.Context, in *CreateVolumeAgentWithDownloadingRequest, opts ...grpc.CallOption) (*VolumeAgent, error)
	DeleteVolumeAgent(ctx context.Context, in *DeleteVolumeAgentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type volumeAgentServiceClient struct {
	cc *grpc.ClientConn
}

func NewVolumeAgentServiceClient(cc *grpc.ClientConn) VolumeAgentServiceClient {
	return &volumeAgentServiceClient{cc}
}

func (c *volumeAgentServiceClient) CreateEmptyVolumeAgent(ctx context.Context, in *CreateEmptyVolumeAgentRequest, opts ...grpc.CallOption) (*VolumeAgent, error) {
	out := new(VolumeAgent)
	err := grpc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.VolumeAgentService/CreateEmptyVolumeAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeAgentServiceClient) CreateVolumeAgentWithDownloading(ctx context.Context, in *CreateVolumeAgentWithDownloadingRequest, opts ...grpc.CallOption) (*VolumeAgent, error) {
	out := new(VolumeAgent)
	err := grpc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.VolumeAgentService/CreateVolumeAgentWithDownloading", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeAgentServiceClient) DeleteVolumeAgent(ctx context.Context, in *DeleteVolumeAgentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.internal.n0core.provisioning.VolumeAgentService/DeleteVolumeAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VolumeAgentService service

type VolumeAgentServiceServer interface {
	CreateEmptyVolumeAgent(context.Context, *CreateEmptyVolumeAgentRequest) (*VolumeAgent, error)
	CreateVolumeAgentWithDownloading(context.Context, *CreateVolumeAgentWithDownloadingRequest) (*VolumeAgent, error)
	DeleteVolumeAgent(context.Context, *DeleteVolumeAgentRequest) (*empty.Empty, error)
}

func RegisterVolumeAgentServiceServer(s *grpc.Server, srv VolumeAgentServiceServer) {
	s.RegisterService(&_VolumeAgentService_serviceDesc, srv)
}

func _VolumeAgentService_CreateEmptyVolumeAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmptyVolumeAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeAgentServiceServer).CreateEmptyVolumeAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.VolumeAgentService/CreateEmptyVolumeAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeAgentServiceServer).CreateEmptyVolumeAgent(ctx, req.(*CreateEmptyVolumeAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeAgentService_CreateVolumeAgentWithDownloading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeAgentWithDownloadingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeAgentServiceServer).CreateVolumeAgentWithDownloading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.VolumeAgentService/CreateVolumeAgentWithDownloading",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeAgentServiceServer).CreateVolumeAgentWithDownloading(ctx, req.(*CreateVolumeAgentWithDownloadingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeAgentService_DeleteVolumeAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeAgentServiceServer).DeleteVolumeAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.internal.n0core.provisioning.VolumeAgentService/DeleteVolumeAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeAgentServiceServer).DeleteVolumeAgent(ctx, req.(*DeleteVolumeAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VolumeAgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.internal.n0core.provisioning.VolumeAgentService",
	HandlerType: (*VolumeAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmptyVolumeAgent",
			Handler:    _VolumeAgentService_CreateEmptyVolumeAgent_Handler,
		},
		{
			MethodName: "CreateVolumeAgentWithDownloading",
			Handler:    _VolumeAgentService_CreateVolumeAgentWithDownloading_Handler,
		},
		{
			MethodName: "DeleteVolumeAgent",
			Handler:    _VolumeAgentService_DeleteVolumeAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/provisioning/volume_agent.proto",
}

func init() {
	proto.RegisterFile("pkg/api/provisioning/volume_agent.proto", fileDescriptor_volume_agent_4c16863b3472c44e)
}

var fileDescriptor_volume_agent_4c16863b3472c44e = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0x4a, 0xc3, 0x30,
	0x18, 0x85, 0x57, 0x37, 0x85, 0xc5, 0x2b, 0x83, 0x8c, 0x32, 0x19, 0x8c, 0x22, 0x6c, 0x57, 0xc9,
	0xd4, 0xcb, 0xa1, 0xa0, 0x9b, 0x17, 0x22, 0xde, 0x4c, 0x54, 0xf0, 0x66, 0xa4, 0xf5, 0xb7, 0x0b,
	0x4b, 0x93, 0x9a, 0xa6, 0x93, 0xbd, 0x86, 0x0f, 0xe1, 0x7b, 0xf8, 0x66, 0xd2, 0x46, 0x21, 0xea,
	0x74, 0x9b, 0x77, 0x49, 0x39, 0xff, 0x77, 0xd2, 0x73, 0xf8, 0x51, 0x27, 0x9d, 0xc6, 0x94, 0xa5,
	0x9c, 0xa6, 0x5a, 0xcd, 0x78, 0xc6, 0x95, 0xe4, 0x32, 0xa6, 0x33, 0x25, 0xf2, 0x04, 0xc6, 0x2c,
	0x06, 0x69, 0x48, 0xaa, 0x95, 0x51, 0x78, 0x5f, 0xf6, 0x32, 0xc3, 0xa2, 0x29, 0xe1, 0xd2, 0x80,
	0x96, 0x4c, 0x10, 0xd9, 0x8b, 0x94, 0x06, 0xe2, 0x0e, 0x36, 0xf7, 0x62, 0xa5, 0x62, 0x01, 0xb4,
	0x9c, 0x09, 0xf3, 0x47, 0x0a, 0x49, 0x6a, 0xe6, 0x16, 0x11, 0x5c, 0xa2, 0xed, 0xdb, 0x12, 0x7c,
	0x5a, 0x70, 0x31, 0x46, 0x35, 0xc9, 0x12, 0xf0, 0xbd, 0xb6, 0xd7, 0xad, 0x8f, 0xca, 0x33, 0xde,
	0x45, 0x9b, 0xe1, 0xdc, 0x40, 0xe6, 0x6f, 0xb4, 0xbd, 0x6e, 0x6d, 0x64, 0x2f, 0x85, 0x32, 0x65,
	0x66, 0xe2, 0x57, 0xad, 0xb2, 0x38, 0x07, 0x17, 0xa8, 0x35, 0xd0, 0xc0, 0x0c, 0x9c, 0x17, 0x0e,
	0x0e, 0x77, 0x04, 0x4f, 0x39, 0x64, 0x6b, 0xe0, 0x03, 0x8d, 0x3a, 0x16, 0xe5, 0x50, 0xee, 0xb8,
	0x99, 0x0c, 0xd5, 0xb3, 0x14, 0x8a, 0x3d, 0x70, 0x19, 0xaf, 0x0d, 0xc5, 0x2d, 0x84, 0x32, 0x95,
	0xeb, 0x08, 0xc6, 0xb9, 0x16, 0x7e, 0xad, 0xd4, 0xd7, 0xed, 0x97, 0x1b, 0x2d, 0x02, 0x82, 0xfc,
	0x21, 0x08, 0xf8, 0xe2, 0xe9, 0x98, 0x7c, 0xff, 0xdd, 0xc3, 0xb7, 0x2a, 0xc2, 0x8e, 0xf4, 0x1a,
	0xf4, 0x8c, 0x47, 0x80, 0x5f, 0x3c, 0xd4, 0x58, 0x1c, 0x03, 0x1e, 0x90, 0x55, 0x1a, 0x23, 0x7f,
	0x86, 0xd8, 0x3c, 0x58, 0x0d, 0xe2, 0x4c, 0x06, 0x15, 0xfc, 0xea, 0xa1, 0xf6, 0xb2, 0x40, 0xf1,
	0xd5, 0x3a, 0xcf, 0x5b, 0x5a, 0xcc, 0xff, 0x1e, 0x9a, 0xa0, 0x9d, 0x1f, 0x25, 0xe0, 0x93, 0xd5,
	0x48, 0xbf, 0xb5, 0xd7, 0x6c, 0x10, 0xbb, 0x03, 0xe4, 0x73, 0x07, 0x48, 0x19, 0x6e, 0x50, 0x39,
	0x3b, 0xbe, 0xef, 0xc7, 0xdc, 0x4c, 0xf2, 0x90, 0x44, 0x2a, 0xa1, 0x1f, 0x2e, 0xd4, 0xc2, 0xe9,
	0xa2, 0x3d, 0xec, 0xbb, 0x97, 0x70, 0xab, 0x04, 0x1e, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x5b, 0x15, 0x02, 0xb3, 0x03, 0x00, 0x00,
}
