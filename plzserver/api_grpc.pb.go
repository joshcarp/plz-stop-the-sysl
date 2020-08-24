// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plzserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MyserverClient is the client API for Myserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyserverClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type myserverClient struct {
	cc grpc.ClientConnInterface
}

func NewMyserverClient(cc grpc.ClientConnInterface) MyserverClient {
	return &myserverClient{cc}
}

func (c *myserverClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/plzserver.myserver/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyserverServer is the server API for Myserver service.
// All implementations must embed UnimplementedMyserverServer
// for forward compatibility
type MyserverServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedMyserverServer()
}

// UnimplementedMyserverServer must be embedded to have forward compatible implementations.
type UnimplementedMyserverServer struct {
}

func (*UnimplementedMyserverServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedMyserverServer) mustEmbedUnimplementedMyserverServer() {}

func RegisterMyserverServer(s *grpc.Server, srv MyserverServer) {
	s.RegisterService(&_Myserver_serviceDesc, srv)
}

func _Myserver_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyserverServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plzserver.myserver/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyserverServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Myserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plzserver.myserver",
	HandlerType: (*MyserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Myserver_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}