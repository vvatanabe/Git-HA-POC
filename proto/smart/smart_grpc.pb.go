// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package smart

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SmartProtocolServiceClient is the client API for SmartProtocolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartProtocolServiceClient interface {
	PostUploadPack(ctx context.Context, in *UploadPackRequest, opts ...grpc.CallOption) (SmartProtocolService_PostUploadPackClient, error)
	PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SmartProtocolService_PostReceivePackClient, error)
	GetInfoRefs(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (*InfoRefsResponse, error)
}

type smartProtocolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartProtocolServiceClient(cc grpc.ClientConnInterface) SmartProtocolServiceClient {
	return &smartProtocolServiceClient{cc}
}

func (c *smartProtocolServiceClient) PostUploadPack(ctx context.Context, in *UploadPackRequest, opts ...grpc.CallOption) (SmartProtocolService_PostUploadPackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SmartProtocolService_serviceDesc.Streams[0], "/smart.SmartProtocolService/PostUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartProtocolServicePostUploadPackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartProtocolService_PostUploadPackClient interface {
	Recv() (*UploadPackResponse, error)
	grpc.ClientStream
}

type smartProtocolServicePostUploadPackClient struct {
	grpc.ClientStream
}

func (x *smartProtocolServicePostUploadPackClient) Recv() (*UploadPackResponse, error) {
	m := new(UploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartProtocolServiceClient) PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SmartProtocolService_PostReceivePackClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SmartProtocolService_serviceDesc.Streams[1], "/smart.SmartProtocolService/PostReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartProtocolServicePostReceivePackClient{stream}
	return x, nil
}

type SmartProtocolService_PostReceivePackClient interface {
	Send(*ReceivePackRequest) error
	CloseAndRecv() (*ReceivePackResponse, error)
	grpc.ClientStream
}

type smartProtocolServicePostReceivePackClient struct {
	grpc.ClientStream
}

func (x *smartProtocolServicePostReceivePackClient) Send(m *ReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *smartProtocolServicePostReceivePackClient) CloseAndRecv() (*ReceivePackResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartProtocolServiceClient) GetInfoRefs(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (*InfoRefsResponse, error) {
	out := new(InfoRefsResponse)
	err := c.cc.Invoke(ctx, "/smart.SmartProtocolService/GetInfoRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartProtocolServiceServer is the server API for SmartProtocolService service.
// All implementations must embed UnimplementedSmartProtocolServiceServer
// for forward compatibility
type SmartProtocolServiceServer interface {
	PostUploadPack(*UploadPackRequest, SmartProtocolService_PostUploadPackServer) error
	PostReceivePack(SmartProtocolService_PostReceivePackServer) error
	GetInfoRefs(context.Context, *InfoRefsRequest) (*InfoRefsResponse, error)
	mustEmbedUnimplementedSmartProtocolServiceServer()
}

// UnimplementedSmartProtocolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmartProtocolServiceServer struct {
}

func (UnimplementedSmartProtocolServiceServer) PostUploadPack(*UploadPackRequest, SmartProtocolService_PostUploadPackServer) error {
	return status.Errorf(codes.Unimplemented, "method PostUploadPack not implemented")
}
func (UnimplementedSmartProtocolServiceServer) PostReceivePack(SmartProtocolService_PostReceivePackServer) error {
	return status.Errorf(codes.Unimplemented, "method PostReceivePack not implemented")
}
func (UnimplementedSmartProtocolServiceServer) GetInfoRefs(context.Context, *InfoRefsRequest) (*InfoRefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoRefs not implemented")
}
func (UnimplementedSmartProtocolServiceServer) mustEmbedUnimplementedSmartProtocolServiceServer() {}

// UnsafeSmartProtocolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartProtocolServiceServer will
// result in compilation errors.
type UnsafeSmartProtocolServiceServer interface {
	mustEmbedUnimplementedSmartProtocolServiceServer()
}

func RegisterSmartProtocolServiceServer(s grpc.ServiceRegistrar, srv SmartProtocolServiceServer) {
	s.RegisterService(&_SmartProtocolService_serviceDesc, srv)
}

func _SmartProtocolService_PostUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UploadPackRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartProtocolServiceServer).PostUploadPack(m, &smartProtocolServicePostUploadPackServer{stream})
}

type SmartProtocolService_PostUploadPackServer interface {
	Send(*UploadPackResponse) error
	grpc.ServerStream
}

type smartProtocolServicePostUploadPackServer struct {
	grpc.ServerStream
}

func (x *smartProtocolServicePostUploadPackServer) Send(m *UploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SmartProtocolService_PostReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SmartProtocolServiceServer).PostReceivePack(&smartProtocolServicePostReceivePackServer{stream})
}

type SmartProtocolService_PostReceivePackServer interface {
	SendAndClose(*ReceivePackResponse) error
	Recv() (*ReceivePackRequest, error)
	grpc.ServerStream
}

type smartProtocolServicePostReceivePackServer struct {
	grpc.ServerStream
}

func (x *smartProtocolServicePostReceivePackServer) SendAndClose(m *ReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *smartProtocolServicePostReceivePackServer) Recv() (*ReceivePackRequest, error) {
	m := new(ReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SmartProtocolService_GetInfoRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartProtocolServiceServer).GetInfoRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smart.SmartProtocolService/GetInfoRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartProtocolServiceServer).GetInfoRefs(ctx, req.(*InfoRefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmartProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smart.SmartProtocolService",
	HandlerType: (*SmartProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfoRefs",
			Handler:    _SmartProtocolService_GetInfoRefs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostUploadPack",
			Handler:       _SmartProtocolService_PostUploadPack_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostReceivePack",
			Handler:       _SmartProtocolService_PostReceivePack_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "smart/smart.proto",
}