// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.1
// source: transport/internet/grpc/encoding/stream.proto

package encoding

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GRPCService_Tun_FullMethodName      = "/xray.transport.internet.grpc.encoding.GRPCService/Tun"
	GRPCService_TunMulti_FullMethodName = "/xray.transport.internet.grpc.encoding.GRPCService/TunMulti"
)

// GRPCServiceClient is the client API for GRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCServiceClient interface {
	Tun(ctx context.Context, opts ...grpc.CallOption) (GRPCService_TunClient, error)
	TunMulti(ctx context.Context, opts ...grpc.CallOption) (GRPCService_TunMultiClient, error)
}

type gRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCServiceClient(cc grpc.ClientConnInterface) GRPCServiceClient {
	return &gRPCServiceClient{cc}
}

func (c *gRPCServiceClient) Tun(ctx context.Context, opts ...grpc.CallOption) (GRPCService_TunClient, error) {
	stream, err := c.cc.NewStream(ctx, &GRPCService_ServiceDesc.Streams[0], GRPCService_Tun_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCServiceTunClient{stream}
	return x, nil
}

type GRPCService_TunClient interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ClientStream
}

type gRPCServiceTunClient struct {
	grpc.ClientStream
}

func (x *gRPCServiceTunClient) Send(m *Hunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCServiceTunClient) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCServiceClient) TunMulti(ctx context.Context, opts ...grpc.CallOption) (GRPCService_TunMultiClient, error) {
	stream, err := c.cc.NewStream(ctx, &GRPCService_ServiceDesc.Streams[1], GRPCService_TunMulti_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCServiceTunMultiClient{stream}
	return x, nil
}

type GRPCService_TunMultiClient interface {
	Send(*MultiHunk) error
	Recv() (*MultiHunk, error)
	grpc.ClientStream
}

type gRPCServiceTunMultiClient struct {
	grpc.ClientStream
}

func (x *gRPCServiceTunMultiClient) Send(m *MultiHunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gRPCServiceTunMultiClient) Recv() (*MultiHunk, error) {
	m := new(MultiHunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GRPCServiceServer is the server API for GRPCService service.
// All implementations must embed UnimplementedGRPCServiceServer
// for forward compatibility
type GRPCServiceServer interface {
	Tun(GRPCService_TunServer) error
	TunMulti(GRPCService_TunMultiServer) error
	mustEmbedUnimplementedGRPCServiceServer()
}

// UnimplementedGRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCServiceServer struct {
}

func (UnimplementedGRPCServiceServer) Tun(GRPCService_TunServer) error {
	return status.Errorf(codes.Unimplemented, "method Tun not implemented")
}
func (UnimplementedGRPCServiceServer) TunMulti(GRPCService_TunMultiServer) error {
	return status.Errorf(codes.Unimplemented, "method TunMulti not implemented")
}
func (UnimplementedGRPCServiceServer) mustEmbedUnimplementedGRPCServiceServer() {}

// UnsafeGRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCServiceServer will
// result in compilation errors.
type UnsafeGRPCServiceServer interface {
	mustEmbedUnimplementedGRPCServiceServer()
}

func RegisterGRPCServiceServer(s grpc.ServiceRegistrar, srv GRPCServiceServer) {
	s.RegisterService(&GRPCService_ServiceDesc, srv)
}

func _GRPCService_Tun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCServiceServer).Tun(&gRPCServiceTunServer{stream})
}

type GRPCService_TunServer interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ServerStream
}

type gRPCServiceTunServer struct {
	grpc.ServerStream
}

func (x *gRPCServiceTunServer) Send(m *Hunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCServiceTunServer) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GRPCService_TunMulti_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GRPCServiceServer).TunMulti(&gRPCServiceTunMultiServer{stream})
}

type GRPCService_TunMultiServer interface {
	Send(*MultiHunk) error
	Recv() (*MultiHunk, error)
	grpc.ServerStream
}

type gRPCServiceTunMultiServer struct {
	grpc.ServerStream
}

func (x *gRPCServiceTunMultiServer) Send(m *MultiHunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gRPCServiceTunMultiServer) Recv() (*MultiHunk, error) {
	m := new(MultiHunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GRPCService_ServiceDesc is the grpc.ServiceDesc for GRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xray.transport.internet.grpc.encoding.GRPCService",
	HandlerType: (*GRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tun",
			Handler:       _GRPCService_Tun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TunMulti",
			Handler:       _GRPCService_TunMulti_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "transport/internet/grpc/encoding/stream.proto",
}
