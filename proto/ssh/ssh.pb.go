// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssh/ssh.proto

package ssh

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Repository struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5394c6fddbe608, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Repository) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

type UploadPackRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Data                 []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UploadPackRequest) Reset()         { *m = UploadPackRequest{} }
func (m *UploadPackRequest) String() string { return proto.CompactTextString(m) }
func (*UploadPackRequest) ProtoMessage()    {}
func (*UploadPackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5394c6fddbe608, []int{1}
}

func (m *UploadPackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPackRequest.Unmarshal(m, b)
}
func (m *UploadPackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPackRequest.Marshal(b, m, deterministic)
}
func (m *UploadPackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPackRequest.Merge(m, src)
}
func (m *UploadPackRequest) XXX_Size() int {
	return xxx_messageInfo_UploadPackRequest.Size(m)
}
func (m *UploadPackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPackRequest proto.InternalMessageInfo

func (m *UploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *UploadPackRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadPackResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Err                  []byte   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPackResponse) Reset()         { *m = UploadPackResponse{} }
func (m *UploadPackResponse) String() string { return proto.CompactTextString(m) }
func (*UploadPackResponse) ProtoMessage()    {}
func (*UploadPackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5394c6fddbe608, []int{2}
}

func (m *UploadPackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPackResponse.Unmarshal(m, b)
}
func (m *UploadPackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPackResponse.Marshal(b, m, deterministic)
}
func (m *UploadPackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPackResponse.Merge(m, src)
}
func (m *UploadPackResponse) XXX_Size() int {
	return xxx_messageInfo_UploadPackResponse.Size(m)
}
func (m *UploadPackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPackResponse proto.InternalMessageInfo

func (m *UploadPackResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadPackResponse) GetErr() []byte {
	if m != nil {
		return m.Err
	}
	return nil
}

type ReceivePackRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Data                 []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReceivePackRequest) Reset()         { *m = ReceivePackRequest{} }
func (m *ReceivePackRequest) String() string { return proto.CompactTextString(m) }
func (*ReceivePackRequest) ProtoMessage()    {}
func (*ReceivePackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5394c6fddbe608, []int{3}
}

func (m *ReceivePackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivePackRequest.Unmarshal(m, b)
}
func (m *ReceivePackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivePackRequest.Marshal(b, m, deterministic)
}
func (m *ReceivePackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivePackRequest.Merge(m, src)
}
func (m *ReceivePackRequest) XXX_Size() int {
	return xxx_messageInfo_ReceivePackRequest.Size(m)
}
func (m *ReceivePackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivePackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivePackRequest proto.InternalMessageInfo

func (m *ReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *ReceivePackRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ReceivePackResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Err                  []byte   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceivePackResponse) Reset()         { *m = ReceivePackResponse{} }
func (m *ReceivePackResponse) String() string { return proto.CompactTextString(m) }
func (*ReceivePackResponse) ProtoMessage()    {}
func (*ReceivePackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5394c6fddbe608, []int{4}
}

func (m *ReceivePackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivePackResponse.Unmarshal(m, b)
}
func (m *ReceivePackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivePackResponse.Marshal(b, m, deterministic)
}
func (m *ReceivePackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivePackResponse.Merge(m, src)
}
func (m *ReceivePackResponse) XXX_Size() int {
	return xxx_messageInfo_ReceivePackResponse.Size(m)
}
func (m *ReceivePackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivePackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivePackResponse proto.InternalMessageInfo

func (m *ReceivePackResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReceivePackResponse) GetErr() []byte {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*Repository)(nil), "ssh.Repository")
	proto.RegisterType((*UploadPackRequest)(nil), "ssh.UploadPackRequest")
	proto.RegisterType((*UploadPackResponse)(nil), "ssh.UploadPackResponse")
	proto.RegisterType((*ReceivePackRequest)(nil), "ssh.ReceivePackRequest")
	proto.RegisterType((*ReceivePackResponse)(nil), "ssh.ReceivePackResponse")
}

func init() { proto.RegisterFile("ssh/ssh.proto", fileDescriptor_3b5394c6fddbe608) }

var fileDescriptor_3b5394c6fddbe608 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x89, 0x15, 0xc1, 0xf1, 0x4f, 0x75, 0x04, 0x5b, 0x7a, 0x92, 0x3d, 0x55, 0xa4, 0x5d,
	0xa9, 0x9e, 0xf4, 0x26, 0x08, 0x3d, 0x96, 0x14, 0x41, 0xbd, 0xa5, 0xe9, 0xd0, 0x5d, 0xac, 0x4d,
	0xcc, 0x64, 0x17, 0xfc, 0x40, 0x7e, 0x4f, 0x49, 0x5a, 0xba, 0xab, 0x7b, 0x12, 0xbc, 0x3d, 0x26,
	0x93, 0xdf, 0xbc, 0x79, 0x0c, 0x1c, 0x31, 0x67, 0x29, 0x73, 0x36, 0xb4, 0xce, 0x78, 0x83, 0x2d,
	0xe6, 0x2c, 0xb9, 0x05, 0x90, 0x64, 0x0d, 0xe7, 0xde, 0xb8, 0x4f, 0x44, 0xd8, 0x2d, 0x98, 0x5c,
	0x57, 0x5c, 0x88, 0xfe, 0xbe, 0x8c, 0x3a, 0xd4, 0x1c, 0x59, 0xd3, 0xdd, 0x59, 0xd7, 0x82, 0x4e,
	0x9e, 0xe1, 0xf4, 0xc9, 0x2e, 0x8d, 0x9a, 0x4f, 0x94, 0x7e, 0x93, 0xf4, 0x51, 0x10, 0x7b, 0x4c,
	0x01, 0xdc, 0x16, 0x15, 0x11, 0x07, 0xa3, 0xf6, 0x30, 0xcc, 0xab, 0x26, 0xc8, 0x5a, 0x4b, 0x20,
	0xcf, 0x95, 0x57, 0x91, 0x7c, 0x28, 0xa3, 0x4e, 0xee, 0x00, 0xeb, 0x64, 0xb6, 0x66, 0xc5, 0xb4,
	0xed, 0x14, 0x55, 0x27, 0x9e, 0x40, 0x8b, 0x9c, 0xdb, 0x7c, 0x0e, 0x32, 0x79, 0x01, 0x94, 0xa4,
	0x29, 0x2f, 0xe9, 0xdf, 0x6d, 0xdd, 0xc3, 0xd9, 0x0f, 0xf4, 0x5f, 0x7c, 0x8d, 0xbe, 0x04, 0xe0,
	0x74, 0x3a, 0x9e, 0x84, 0xd4, 0xb5, 0x59, 0x4e, 0xc9, 0x95, 0xb9, 0x26, 0x7c, 0x84, 0xe3, 0x89,
	0x61, 0x5f, 0xad, 0x8b, 0xe7, 0xd1, 0x56, 0x23, 0xd9, 0x5e, 0xa7, 0x51, 0x5f, 0xcf, 0xef, 0x8b,
	0x6b, 0x81, 0x63, 0x68, 0x07, 0x4c, 0xcd, 0x1e, 0x76, 0x36, 0xeb, 0xfd, 0xce, 0xa2, 0xd7, 0x6d,
	0x3e, 0x54, 0xa4, 0x87, 0xab, 0xd7, 0xcb, 0x45, 0xee, 0xb3, 0x62, 0x36, 0xd4, 0xe6, 0x3d, 0x2d,
	0x4b, 0xe5, 0xd5, 0x4a, 0xcd, 0x28, 0x5d, 0xe4, 0x7e, 0x90, 0xa9, 0x81, 0x35, 0x3a, 0x8d, 0x97,
	0x13, 0x6e, 0x68, 0xb6, 0x17, 0xe5, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xc2, 0x86,
	0x72, 0x55, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SSHProtocolServiceClient is the client API for SSHProtocolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SSHProtocolServiceClient interface {
	PostUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHProtocolService_PostUploadPackClient, error)
	PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHProtocolService_PostReceivePackClient, error)
}

type sSHProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewSSHProtocolServiceClient(cc *grpc.ClientConn) SSHProtocolServiceClient {
	return &sSHProtocolServiceClient{cc}
}

func (c *sSHProtocolServiceClient) PostUploadPack(ctx context.Context, opts ...grpc.CallOption) (SSHProtocolService_PostUploadPackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SSHProtocolService_serviceDesc.Streams[0], "/ssh.SSHProtocolService/PostUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHProtocolServicePostUploadPackClient{stream}
	return x, nil
}

type SSHProtocolService_PostUploadPackClient interface {
	Send(*UploadPackRequest) error
	Recv() (*UploadPackResponse, error)
	grpc.ClientStream
}

type sSHProtocolServicePostUploadPackClient struct {
	grpc.ClientStream
}

func (x *sSHProtocolServicePostUploadPackClient) Send(m *UploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHProtocolServicePostUploadPackClient) Recv() (*UploadPackResponse, error) {
	m := new(UploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sSHProtocolServiceClient) PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SSHProtocolService_PostReceivePackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SSHProtocolService_serviceDesc.Streams[1], "/ssh.SSHProtocolService/PostReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHProtocolServicePostReceivePackClient{stream}
	return x, nil
}

type SSHProtocolService_PostReceivePackClient interface {
	Send(*ReceivePackRequest) error
	Recv() (*ReceivePackResponse, error)
	grpc.ClientStream
}

type sSHProtocolServicePostReceivePackClient struct {
	grpc.ClientStream
}

func (x *sSHProtocolServicePostReceivePackClient) Send(m *ReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHProtocolServicePostReceivePackClient) Recv() (*ReceivePackResponse, error) {
	m := new(ReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SSHProtocolServiceServer is the server API for SSHProtocolService service.
type SSHProtocolServiceServer interface {
	PostUploadPack(SSHProtocolService_PostUploadPackServer) error
	PostReceivePack(SSHProtocolService_PostReceivePackServer) error
}

// UnimplementedSSHProtocolServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSSHProtocolServiceServer struct {
}

func (*UnimplementedSSHProtocolServiceServer) PostUploadPack(srv SSHProtocolService_PostUploadPackServer) error {
	return status.Errorf(codes.Unimplemented, "method PostUploadPack not implemented")
}
func (*UnimplementedSSHProtocolServiceServer) PostReceivePack(srv SSHProtocolService_PostReceivePackServer) error {
	return status.Errorf(codes.Unimplemented, "method PostReceivePack not implemented")
}

func RegisterSSHProtocolServiceServer(s *grpc.Server, srv SSHProtocolServiceServer) {
	s.RegisterService(&_SSHProtocolService_serviceDesc, srv)
}

func _SSHProtocolService_PostUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHProtocolServiceServer).PostUploadPack(&sSHProtocolServicePostUploadPackServer{stream})
}

type SSHProtocolService_PostUploadPackServer interface {
	Send(*UploadPackResponse) error
	Recv() (*UploadPackRequest, error)
	grpc.ServerStream
}

type sSHProtocolServicePostUploadPackServer struct {
	grpc.ServerStream
}

func (x *sSHProtocolServicePostUploadPackServer) Send(m *UploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHProtocolServicePostUploadPackServer) Recv() (*UploadPackRequest, error) {
	m := new(UploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SSHProtocolService_PostReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHProtocolServiceServer).PostReceivePack(&sSHProtocolServicePostReceivePackServer{stream})
}

type SSHProtocolService_PostReceivePackServer interface {
	Send(*ReceivePackResponse) error
	Recv() (*ReceivePackRequest, error)
	grpc.ServerStream
}

type sSHProtocolServicePostReceivePackServer struct {
	grpc.ServerStream
}

func (x *sSHProtocolServicePostReceivePackServer) Send(m *ReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHProtocolServicePostReceivePackServer) Recv() (*ReceivePackRequest, error) {
	m := new(ReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SSHProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ssh.SSHProtocolService",
	HandlerType: (*SSHProtocolServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostUploadPack",
			Handler:       _SSHProtocolService_PostUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PostReceivePack",
			Handler:       _SSHProtocolService_PostReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ssh/ssh.proto",
}
