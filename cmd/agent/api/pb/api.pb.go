// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HostnameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameRequest) Reset()         { *m = HostnameRequest{} }
func (m *HostnameRequest) String() string { return proto.CompactTextString(m) }
func (*HostnameRequest) ProtoMessage()    {}
func (*HostnameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *HostnameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameRequest.Unmarshal(m, b)
}
func (m *HostnameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameRequest.Marshal(b, m, deterministic)
}
func (m *HostnameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameRequest.Merge(m, src)
}
func (m *HostnameRequest) XXX_Size() int {
	return xxx_messageInfo_HostnameRequest.Size(m)
}
func (m *HostnameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameRequest proto.InternalMessageInfo

// The response message containing the requested hostname
type HostnameReply struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameReply) Reset()         { *m = HostnameReply{} }
func (m *HostnameReply) String() string { return proto.CompactTextString(m) }
func (*HostnameReply) ProtoMessage()    {}
func (*HostnameReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *HostnameReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameReply.Unmarshal(m, b)
}
func (m *HostnameReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameReply.Marshal(b, m, deterministic)
}
func (m *HostnameReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameReply.Merge(m, src)
}
func (m *HostnameReply) XXX_Size() int {
	return xxx_messageInfo_HostnameReply.Size(m)
}
func (m *HostnameReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameReply.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameReply proto.InternalMessageInfo

func (m *HostnameReply) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

// The request message containing the tag list for an entity.
type TagRequest struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagRequest) Reset()         { *m = TagRequest{} }
func (m *TagRequest) String() string { return proto.CompactTextString(m) }
func (*TagRequest) ProtoMessage()    {}
func (*TagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *TagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagRequest.Unmarshal(m, b)
}
func (m *TagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagRequest.Marshal(b, m, deterministic)
}
func (m *TagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagRequest.Merge(m, src)
}
func (m *TagRequest) XXX_Size() int {
	return xxx_messageInfo_TagRequest.Size(m)
}
func (m *TagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TagRequest proto.InternalMessageInfo

func (m *TagRequest) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

// The response message containing the tagger reply
type TagReply struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagReply) Reset()         { *m = TagReply{} }
func (m *TagReply) String() string { return proto.CompactTextString(m) }
func (*TagReply) ProtoMessage()    {}
func (*TagReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *TagReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagReply.Unmarshal(m, b)
}
func (m *TagReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagReply.Marshal(b, m, deterministic)
}
func (m *TagReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagReply.Merge(m, src)
}
func (m *TagReply) XXX_Size() int {
	return xxx_messageInfo_TagReply.Size(m)
}
func (m *TagReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TagReply.DiscardUnknown(m)
}

var xxx_messageInfo_TagReply proto.InternalMessageInfo

func (m *TagReply) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*HostnameRequest)(nil), "pb.HostnameRequest")
	proto.RegisterType((*HostnameReply)(nil), "pb.HostnameReply")
	proto.RegisterType((*TagRequest)(nil), "pb.TagRequest")
	proto.RegisterType((*TagReply)(nil), "pb.TagReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x49, 0xd5, 0xda, 0xbc, 0x5a, 0x4b, 0x9f, 0x28, 0x21, 0x88, 0x94, 0xc1, 0x45, 0x51,
	0x48, 0xb0, 0xee, 0xdc, 0x75, 0x55, 0x17, 0x2e, 0x24, 0xd6, 0x03, 0x4c, 0xca, 0x63, 0x0c, 0xc4,
	0x99, 0x31, 0x79, 0x15, 0xb2, 0xf5, 0x0a, 0x1e, 0xcd, 0x2b, 0x78, 0x10, 0x99, 0x31, 0x36, 0xd8,
	0xdd, 0xfc, 0x3f, 0xff, 0x7c, 0xf3, 0x31, 0x10, 0x4a, 0x5b, 0x24, 0xb6, 0x32, 0x6c, 0xb0, 0x67,
	0xf3, 0xf8, 0x5c, 0x19, 0xa3, 0x4a, 0x4a, 0xa5, 0x2d, 0x52, 0xa9, 0xb5, 0x61, 0xc9, 0x85, 0xd1,
	0xf5, 0xef, 0x42, 0x4c, 0x60, 0x7c, 0x6f, 0x6a, 0xd6, 0xf2, 0x95, 0x32, 0x7a, 0xdb, 0x50, 0xcd,
	0xe2, 0x1a, 0x46, 0x5d, 0x65, 0xcb, 0x06, 0x63, 0x18, 0xbc, 0xb4, 0x45, 0x14, 0x4c, 0x83, 0x59,
	0x98, 0x6d, 0xb3, 0xb8, 0x04, 0x58, 0x49, 0xd5, 0x5e, 0xc5, 0x33, 0xe8, 0x93, 0xe6, 0x82, 0x9b,
	0x76, 0xd7, 0x26, 0x71, 0x01, 0x03, 0xbf, 0x72, 0x34, 0x84, 0x7d, 0x96, 0xaa, 0x8e, 0x82, 0xe9,
	0xde, 0x2c, 0xcc, 0xfc, 0x79, 0xfe, 0x0c, 0x07, 0x0b, 0x45, 0x9a, 0xf1, 0x01, 0x86, 0x4b, 0xe2,
	0xbf, 0xe7, 0xf1, 0x24, 0xb1, 0x79, 0xb2, 0xe3, 0x17, 0x4f, 0xfe, 0x97, 0xb6, 0x6c, 0xc4, 0xe9,
	0xc7, 0xd7, 0xf7, 0x67, 0x6f, 0x8c, 0xa3, 0xf4, 0xfd, 0x26, 0x55, 0x95, 0x5d, 0xa7, 0x4e, 0x70,
	0xfe, 0x08, 0x43, 0x8f, 0x7d, 0xa2, 0xf5, 0xa6, 0x22, 0x5c, 0xc0, 0xe1, 0x92, 0x78, 0x25, 0x55,
	0x8d, 0xc7, 0x8e, 0xd1, 0x89, 0xc7, 0x47, 0xdb, 0xec, 0x70, 0x91, 0xc7, 0xa1, 0xe8, 0x70, 0xce,
	0xf2, 0x2e, 0xb8, 0xca, 0xfb, 0xfe, 0xd7, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0x07,
	0x12, 0xc5, 0x64, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	// get the hostname
	GetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) GetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error) {
	out := new(HostnameReply)
	err := c.cc.Invoke(ctx, "/pb.Agent/GetHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	// get the hostname
	GetHostname(context.Context, *HostnameRequest) (*HostnameReply, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) GetHostname(ctx context.Context, req *HostnameRequest) (*HostnameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostname not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_GetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/GetHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetHostname(ctx, req.(*HostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHostname",
			Handler:    _Agent_GetHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// AgentSecureClient is the client API for AgentSecure service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentSecureClient interface {
	GetTags(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*TagReply, error)
}

type agentSecureClient struct {
	cc *grpc.ClientConn
}

func NewAgentSecureClient(cc *grpc.ClientConn) AgentSecureClient {
	return &agentSecureClient{cc}
}

func (c *agentSecureClient) GetTags(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*TagReply, error) {
	out := new(TagReply)
	err := c.cc.Invoke(ctx, "/pb.AgentSecure/GetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentSecureServer is the server API for AgentSecure service.
type AgentSecureServer interface {
	GetTags(context.Context, *TagRequest) (*TagReply, error)
}

// UnimplementedAgentSecureServer can be embedded to have forward compatible implementations.
type UnimplementedAgentSecureServer struct {
}

func (*UnimplementedAgentSecureServer) GetTags(ctx context.Context, req *TagRequest) (*TagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}

func RegisterAgentSecureServer(s *grpc.Server, srv AgentSecureServer) {
	s.RegisterService(&_AgentSecure_serviceDesc, srv)
}

func _AgentSecure_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSecureServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AgentSecure/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSecureServer).GetTags(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentSecure_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AgentSecure",
	HandlerType: (*AgentSecureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTags",
			Handler:    _AgentSecure_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
