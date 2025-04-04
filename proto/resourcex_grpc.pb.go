// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: resourcex.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Resourcex_ResourcexHealthCheck_FullMethodName = "/proto.Resourcex/ResourcexHealthCheck"
	Resourcex_ResourcexUploadFile_FullMethodName  = "/proto.Resourcex/ResourcexUploadFile"
)

// ResourcexClient is the client API for Resourcex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourcexClient interface {
	ResourcexHealthCheck(ctx context.Context, in *ResourcexHealthCheckRequest, opts ...grpc.CallOption) (*ResourcexHealthCheckResponse, error)
	// 客户端流式RPC方法，用于上传文件
	ResourcexUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ResourcexUploadFileChunk, ResourcexUploadFileResponse], error)
}

type resourcexClient struct {
	cc grpc.ClientConnInterface
}

func NewResourcexClient(cc grpc.ClientConnInterface) ResourcexClient {
	return &resourcexClient{cc}
}

func (c *resourcexClient) ResourcexHealthCheck(ctx context.Context, in *ResourcexHealthCheckRequest, opts ...grpc.CallOption) (*ResourcexHealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourcexHealthCheckResponse)
	err := c.cc.Invoke(ctx, Resourcex_ResourcexHealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcexClient) ResourcexUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ResourcexUploadFileChunk, ResourcexUploadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Resourcex_ServiceDesc.Streams[0], Resourcex_ResourcexUploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResourcexUploadFileChunk, ResourcexUploadFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Resourcex_ResourcexUploadFileClient = grpc.ClientStreamingClient[ResourcexUploadFileChunk, ResourcexUploadFileResponse]

// ResourcexServer is the server API for Resourcex service.
// All implementations must embed UnimplementedResourcexServer
// for forward compatibility.
type ResourcexServer interface {
	ResourcexHealthCheck(context.Context, *ResourcexHealthCheckRequest) (*ResourcexHealthCheckResponse, error)
	// 客户端流式RPC方法，用于上传文件
	ResourcexUploadFile(grpc.ClientStreamingServer[ResourcexUploadFileChunk, ResourcexUploadFileResponse]) error
	mustEmbedUnimplementedResourcexServer()
}

// UnimplementedResourcexServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResourcexServer struct{}

func (UnimplementedResourcexServer) ResourcexHealthCheck(context.Context, *ResourcexHealthCheckRequest) (*ResourcexHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourcexHealthCheck not implemented")
}
func (UnimplementedResourcexServer) ResourcexUploadFile(grpc.ClientStreamingServer[ResourcexUploadFileChunk, ResourcexUploadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ResourcexUploadFile not implemented")
}
func (UnimplementedResourcexServer) mustEmbedUnimplementedResourcexServer() {}
func (UnimplementedResourcexServer) testEmbeddedByValue()                   {}

// UnsafeResourcexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourcexServer will
// result in compilation errors.
type UnsafeResourcexServer interface {
	mustEmbedUnimplementedResourcexServer()
}

func RegisterResourcexServer(s grpc.ServiceRegistrar, srv ResourcexServer) {
	// If the following call pancis, it indicates UnimplementedResourcexServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Resourcex_ServiceDesc, srv)
}

func _Resourcex_ResourcexHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourcexHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcexServer).ResourcexHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resourcex_ResourcexHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcexServer).ResourcexHealthCheck(ctx, req.(*ResourcexHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resourcex_ResourcexUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResourcexServer).ResourcexUploadFile(&grpc.GenericServerStream[ResourcexUploadFileChunk, ResourcexUploadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Resourcex_ResourcexUploadFileServer = grpc.ClientStreamingServer[ResourcexUploadFileChunk, ResourcexUploadFileResponse]

// Resourcex_ServiceDesc is the grpc.ServiceDesc for Resourcex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resourcex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Resourcex",
	HandlerType: (*ResourcexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResourcexHealthCheck",
			Handler:    _Resourcex_ResourcexHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ResourcexUploadFile",
			Handler:       _Resourcex_ResourcexUploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "resourcex.proto",
}

const (
	StaticWebsite_StaticWebsiteHealthCheck_FullMethodName = "/proto.StaticWebsite/StaticWebsiteHealthCheck"
	StaticWebsite_StaticWebsiteUploadFile_FullMethodName  = "/proto.StaticWebsite/StaticWebsiteUploadFile"
	StaticWebsite_StaticWebsiteDeleteFile_FullMethodName  = "/proto.StaticWebsite/StaticWebsiteDeleteFile"
	StaticWebsite_StaticWebsitePublish_FullMethodName     = "/proto.StaticWebsite/StaticWebsitePublish"
)

// StaticWebsiteClient is the client API for StaticWebsite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 静态站点服务
type StaticWebsiteClient interface {
	StaticWebsiteHealthCheck(ctx context.Context, in *StaticWebsiteHealthCheckRequest, opts ...grpc.CallOption) (*StaticWebsiteHealthCheckResponse, error)
	// 用于上传静态站点文件
	StaticWebsiteUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse], error)
	// 删除静态站点文件
	StaticWebsiteDeleteFile(ctx context.Context, in *StaticWebsiteDeleteFileRequest, opts ...grpc.CallOption) (*StaticWebsiteDeleteFileResponse, error)
	// 发布站点
	StaticWebsitePublish(ctx context.Context, in *StaticWebsitePublishRequest, opts ...grpc.CallOption) (*StaticWebsitePublishResponse, error)
}

type staticWebsiteClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticWebsiteClient(cc grpc.ClientConnInterface) StaticWebsiteClient {
	return &staticWebsiteClient{cc}
}

func (c *staticWebsiteClient) StaticWebsiteHealthCheck(ctx context.Context, in *StaticWebsiteHealthCheckRequest, opts ...grpc.CallOption) (*StaticWebsiteHealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticWebsiteHealthCheckResponse)
	err := c.cc.Invoke(ctx, StaticWebsite_StaticWebsiteHealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticWebsiteClient) StaticWebsiteUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StaticWebsite_ServiceDesc.Streams[0], StaticWebsite_StaticWebsiteUploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StaticWebsite_StaticWebsiteUploadFileClient = grpc.ClientStreamingClient[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse]

func (c *staticWebsiteClient) StaticWebsiteDeleteFile(ctx context.Context, in *StaticWebsiteDeleteFileRequest, opts ...grpc.CallOption) (*StaticWebsiteDeleteFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticWebsiteDeleteFileResponse)
	err := c.cc.Invoke(ctx, StaticWebsite_StaticWebsiteDeleteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticWebsiteClient) StaticWebsitePublish(ctx context.Context, in *StaticWebsitePublishRequest, opts ...grpc.CallOption) (*StaticWebsitePublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticWebsitePublishResponse)
	err := c.cc.Invoke(ctx, StaticWebsite_StaticWebsitePublish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticWebsiteServer is the server API for StaticWebsite service.
// All implementations must embed UnimplementedStaticWebsiteServer
// for forward compatibility.
//
// 静态站点服务
type StaticWebsiteServer interface {
	StaticWebsiteHealthCheck(context.Context, *StaticWebsiteHealthCheckRequest) (*StaticWebsiteHealthCheckResponse, error)
	// 用于上传静态站点文件
	StaticWebsiteUploadFile(grpc.ClientStreamingServer[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse]) error
	// 删除静态站点文件
	StaticWebsiteDeleteFile(context.Context, *StaticWebsiteDeleteFileRequest) (*StaticWebsiteDeleteFileResponse, error)
	// 发布站点
	StaticWebsitePublish(context.Context, *StaticWebsitePublishRequest) (*StaticWebsitePublishResponse, error)
	mustEmbedUnimplementedStaticWebsiteServer()
}

// UnimplementedStaticWebsiteServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStaticWebsiteServer struct{}

func (UnimplementedStaticWebsiteServer) StaticWebsiteHealthCheck(context.Context, *StaticWebsiteHealthCheckRequest) (*StaticWebsiteHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaticWebsiteHealthCheck not implemented")
}
func (UnimplementedStaticWebsiteServer) StaticWebsiteUploadFile(grpc.ClientStreamingServer[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StaticWebsiteUploadFile not implemented")
}
func (UnimplementedStaticWebsiteServer) StaticWebsiteDeleteFile(context.Context, *StaticWebsiteDeleteFileRequest) (*StaticWebsiteDeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaticWebsiteDeleteFile not implemented")
}
func (UnimplementedStaticWebsiteServer) StaticWebsitePublish(context.Context, *StaticWebsitePublishRequest) (*StaticWebsitePublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaticWebsitePublish not implemented")
}
func (UnimplementedStaticWebsiteServer) mustEmbedUnimplementedStaticWebsiteServer() {}
func (UnimplementedStaticWebsiteServer) testEmbeddedByValue()                       {}

// UnsafeStaticWebsiteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticWebsiteServer will
// result in compilation errors.
type UnsafeStaticWebsiteServer interface {
	mustEmbedUnimplementedStaticWebsiteServer()
}

func RegisterStaticWebsiteServer(s grpc.ServiceRegistrar, srv StaticWebsiteServer) {
	// If the following call pancis, it indicates UnimplementedStaticWebsiteServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StaticWebsite_ServiceDesc, srv)
}

func _StaticWebsite_StaticWebsiteHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaticWebsiteHealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticWebsiteServer).StaticWebsiteHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticWebsite_StaticWebsiteHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticWebsiteServer).StaticWebsiteHealthCheck(ctx, req.(*StaticWebsiteHealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticWebsite_StaticWebsiteUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StaticWebsiteServer).StaticWebsiteUploadFile(&grpc.GenericServerStream[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StaticWebsite_StaticWebsiteUploadFileServer = grpc.ClientStreamingServer[StaticWebsiteUploadFileChunk, StaticWebsiteUploadFileResponse]

func _StaticWebsite_StaticWebsiteDeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaticWebsiteDeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticWebsiteServer).StaticWebsiteDeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticWebsite_StaticWebsiteDeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticWebsiteServer).StaticWebsiteDeleteFile(ctx, req.(*StaticWebsiteDeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticWebsite_StaticWebsitePublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaticWebsitePublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticWebsiteServer).StaticWebsitePublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticWebsite_StaticWebsitePublish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticWebsiteServer).StaticWebsitePublish(ctx, req.(*StaticWebsitePublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaticWebsite_ServiceDesc is the grpc.ServiceDesc for StaticWebsite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaticWebsite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StaticWebsite",
	HandlerType: (*StaticWebsiteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StaticWebsiteHealthCheck",
			Handler:    _StaticWebsite_StaticWebsiteHealthCheck_Handler,
		},
		{
			MethodName: "StaticWebsiteDeleteFile",
			Handler:    _StaticWebsite_StaticWebsiteDeleteFile_Handler,
		},
		{
			MethodName: "StaticWebsitePublish",
			Handler:    _StaticWebsite_StaticWebsitePublish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StaticWebsiteUploadFile",
			Handler:       _StaticWebsite_StaticWebsiteUploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "resourcex.proto",
}
