// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dimutex

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

// DiMutexClient is the client API for DiMutex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiMutexClient interface {
	RequestAccess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	AnswerRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	HoldAndRelease(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error)
	Grant(ctx context.Context, in *Reply, opts ...grpc.CallOption) (*Empty, error)
	Hello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type diMutexClient struct {
	cc grpc.ClientConnInterface
}

func NewDiMutexClient(cc grpc.ClientConnInterface) DiMutexClient {
	return &diMutexClient{cc}
}

func (c *diMutexClient) RequestAccess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dimutex.DiMutex/RequestAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diMutexClient) AnswerRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/dimutex.DiMutex/AnswerRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diMutexClient) HoldAndRelease(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/dimutex.DiMutex/HoldAndRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diMutexClient) Grant(ctx context.Context, in *Reply, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dimutex.DiMutex/Grant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diMutexClient) Hello(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dimutex.DiMutex/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiMutexServer is the server API for DiMutex service.
// All implementations must embed UnimplementedDiMutexServer
// for forward compatibility
type DiMutexServer interface {
	RequestAccess(context.Context, *Request) (*Empty, error)
	AnswerRequest(context.Context, *Request) (*Reply, error)
	HoldAndRelease(context.Context, *Empty) (*Reply, error)
	Grant(context.Context, *Reply) (*Empty, error)
	Hello(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedDiMutexServer()
}

// UnimplementedDiMutexServer must be embedded to have forward compatible implementations.
type UnimplementedDiMutexServer struct {
}

func (UnimplementedDiMutexServer) RequestAccess(context.Context, *Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAccess not implemented")
}
func (UnimplementedDiMutexServer) AnswerRequest(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerRequest not implemented")
}
func (UnimplementedDiMutexServer) HoldAndRelease(context.Context, *Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HoldAndRelease not implemented")
}
func (UnimplementedDiMutexServer) Grant(context.Context, *Reply) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grant not implemented")
}
func (UnimplementedDiMutexServer) Hello(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedDiMutexServer) mustEmbedUnimplementedDiMutexServer() {}

// UnsafeDiMutexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiMutexServer will
// result in compilation errors.
type UnsafeDiMutexServer interface {
	mustEmbedUnimplementedDiMutexServer()
}

func RegisterDiMutexServer(s grpc.ServiceRegistrar, srv DiMutexServer) {
	s.RegisterService(&DiMutex_ServiceDesc, srv)
}

func _DiMutex_RequestAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiMutexServer).RequestAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dimutex.DiMutex/RequestAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiMutexServer).RequestAccess(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiMutex_AnswerRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiMutexServer).AnswerRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dimutex.DiMutex/AnswerRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiMutexServer).AnswerRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiMutex_HoldAndRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiMutexServer).HoldAndRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dimutex.DiMutex/HoldAndRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiMutexServer).HoldAndRelease(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiMutex_Grant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiMutexServer).Grant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dimutex.DiMutex/Grant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiMutexServer).Grant(ctx, req.(*Reply))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiMutex_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiMutexServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dimutex.DiMutex/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiMutexServer).Hello(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DiMutex_ServiceDesc is the grpc.ServiceDesc for DiMutex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiMutex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dimutex.DiMutex",
	HandlerType: (*DiMutexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAccess",
			Handler:    _DiMutex_RequestAccess_Handler,
		},
		{
			MethodName: "AnswerRequest",
			Handler:    _DiMutex_AnswerRequest_Handler,
		},
		{
			MethodName: "HoldAndRelease",
			Handler:    _DiMutex_HoldAndRelease_Handler,
		},
		{
			MethodName: "Grant",
			Handler:    _DiMutex_Grant_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _DiMutex_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "DiMutex.proto",
}
