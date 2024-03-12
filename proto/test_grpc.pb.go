// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: test.proto

package proto

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
	StreamGreeter_Greet_FullMethodName = "/stream.StreamGreeter/greet"
)

// StreamGreeterClient is the config API for StreamGreeter service.yaml.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamGreeterClient interface {
	Greet(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (*StreamResp, error)
}

type streamGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamGreeterClient(cc grpc.ClientConnInterface) StreamGreeterClient {
	return &streamGreeterClient{cc}
}

func (c *streamGreeterClient) Greet(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (*StreamResp, error) {
	out := new(StreamResp)
	err := c.cc.Invoke(ctx, StreamGreeter_Greet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamGreeterServer is the server API for StreamGreeter service.yaml.
// All implementations must embed UnimplementedStreamGreeterServer
// for forward compatibility
type StreamGreeterServer interface {
	Greet(context.Context, *StreamReq) (*StreamResp, error)
	mustEmbedUnimplementedStreamGreeterServer()
}

// UnimplementedStreamGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedStreamGreeterServer struct {
}

func (UnimplementedStreamGreeterServer) Greet(context.Context, *StreamReq) (*StreamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedStreamGreeterServer) mustEmbedUnimplementedStreamGreeterServer() {}

// UnsafeStreamGreeterServer may be embedded to opt out of forward compatibility for this service.yaml.
// Use of this interface is not recommended, as added methods to StreamGreeterServer will
// result in compilation errors.
type UnsafeStreamGreeterServer interface {
	mustEmbedUnimplementedStreamGreeterServer()
}

func RegisterStreamGreeterServer(s grpc.ServiceRegistrar, srv StreamGreeterServer) {
	s.RegisterService(&StreamGreeter_ServiceDesc, srv)
}

func _StreamGreeter_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_Greet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).Greet(ctx, req.(*StreamReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamGreeter_ServiceDesc is the grpc.ServiceDesc for StreamGreeter service.yaml.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamGreeter",
	HandlerType: (*StreamGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "greet",
			Handler:    _StreamGreeter_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
