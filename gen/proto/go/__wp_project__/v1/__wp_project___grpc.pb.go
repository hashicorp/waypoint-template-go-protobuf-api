// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: %%wp_project%%/v1/%%wp_project%%.proto

package %%wp_project%%v1

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

// %%Wp_project%%ServiceClient is the client API for %%Wp_project%%Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type %%Wp_project%%ServiceClient interface {
	// HelloWorld is the default rpc. Feel free to delete.
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type %%wp_project%%ServiceClient struct {
	cc grpc.ClientConnInterface
}

func New%%Wp_project%%ServiceClient(cc grpc.ClientConnInterface) %%Wp_project%%ServiceClient {
	return &%%wp_project%%ServiceClient{cc}
}

func (c *%%wp_project%%ServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/%%wp_project%%.v1.%%Wp_project%%Service/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// %%Wp_project%%ServiceServer is the server API for %%Wp_project%%Service service.
// All implementations must embed Unimplemented%%Wp_project%%ServiceServer
// for forward compatibility
type %%Wp_project%%ServiceServer interface {
	// HelloWorld is the default rpc. Feel free to delete.
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	mustEmbedUnimplemented%%Wp_project%%ServiceServer()
}

// Unimplemented%%Wp_project%%ServiceServer must be embedded to have forward compatible implementations.
type Unimplemented%%Wp_project%%ServiceServer struct {
}

func (Unimplemented%%Wp_project%%ServiceServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (Unimplemented%%Wp_project%%ServiceServer) mustEmbedUnimplemented%%Wp_project%%ServiceServer() {
}

// Unsafe%%Wp_project%%ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to %%Wp_project%%ServiceServer will
// result in compilation errors.
type Unsafe%%Wp_project%%ServiceServer interface {
	mustEmbedUnimplemented%%Wp_project%%ServiceServer()
}

func Register%%Wp_project%%ServiceServer(s grpc.ServiceRegistrar, srv %%Wp_project%%ServiceServer) {
	s.RegisterService(&%%Wp_project%%Service_ServiceDesc, srv)
}

func _%%Wp_project%%Service_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(%%Wp_project%%ServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/%%wp_project%%.v1.%%Wp_project%%Service/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(%%Wp_project%%ServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// %%Wp_project%%Service_ServiceDesc is the grpc.ServiceDesc for %%Wp_project%%Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var %%Wp_project%%Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "%%wp_project%%.v1.%%Wp_project%%Service",
	HandlerType: (*%%Wp_project%%ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _%%Wp_project%%Service_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "%%wp_project%%/v1/%%wp_project%%.proto",
}
