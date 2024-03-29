// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: api/v1/knowledge_base/knowledge_base.proto

package knowledge_base

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
	KnowledgeBaseService_CreateKnowledgeBase_FullMethodName     = "/knowdateapp.knowledge.v1.knowledge_base.KnowledgeBaseService/CreateKnowledgeBase"
	KnowledgeBaseService_GetKnowledgeBaseById_FullMethodName    = "/knowdateapp.knowledge.v1.knowledge_base.KnowledgeBaseService/GetKnowledgeBaseById"
	KnowledgeBaseService_DeleteKnowledgeBaseById_FullMethodName = "/knowdateapp.knowledge.v1.knowledge_base.KnowledgeBaseService/DeleteKnowledgeBaseById"
)

// KnowledgeBaseServiceClient is the client API for KnowledgeBaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KnowledgeBaseServiceClient interface {
	CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*CreateKnowledgeBaseResponse, error)
	GetKnowledgeBaseById(ctx context.Context, in *GetKnowledgeBaseByIdRequest, opts ...grpc.CallOption) (*GetKnowledgeBaseByIdResponse, error)
	DeleteKnowledgeBaseById(ctx context.Context, in *DeleteKnowledgeBaseByIdRequest, opts ...grpc.CallOption) (*DeleteKnowledgeBaseByIdResponse, error)
}

type knowledgeBaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKnowledgeBaseServiceClient(cc grpc.ClientConnInterface) KnowledgeBaseServiceClient {
	return &knowledgeBaseServiceClient{cc}
}

func (c *knowledgeBaseServiceClient) CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*CreateKnowledgeBaseResponse, error) {
	out := new(CreateKnowledgeBaseResponse)
	err := c.cc.Invoke(ctx, KnowledgeBaseService_CreateKnowledgeBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBaseServiceClient) GetKnowledgeBaseById(ctx context.Context, in *GetKnowledgeBaseByIdRequest, opts ...grpc.CallOption) (*GetKnowledgeBaseByIdResponse, error) {
	out := new(GetKnowledgeBaseByIdResponse)
	err := c.cc.Invoke(ctx, KnowledgeBaseService_GetKnowledgeBaseById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBaseServiceClient) DeleteKnowledgeBaseById(ctx context.Context, in *DeleteKnowledgeBaseByIdRequest, opts ...grpc.CallOption) (*DeleteKnowledgeBaseByIdResponse, error) {
	out := new(DeleteKnowledgeBaseByIdResponse)
	err := c.cc.Invoke(ctx, KnowledgeBaseService_DeleteKnowledgeBaseById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnowledgeBaseServiceServer is the server API for KnowledgeBaseService service.
// All implementations must embed UnimplementedKnowledgeBaseServiceServer
// for forward compatibility
type KnowledgeBaseServiceServer interface {
	CreateKnowledgeBase(context.Context, *CreateKnowledgeBaseRequest) (*CreateKnowledgeBaseResponse, error)
	GetKnowledgeBaseById(context.Context, *GetKnowledgeBaseByIdRequest) (*GetKnowledgeBaseByIdResponse, error)
	DeleteKnowledgeBaseById(context.Context, *DeleteKnowledgeBaseByIdRequest) (*DeleteKnowledgeBaseByIdResponse, error)
	mustEmbedUnimplementedKnowledgeBaseServiceServer()
}

// UnimplementedKnowledgeBaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKnowledgeBaseServiceServer struct {
}

func (UnimplementedKnowledgeBaseServiceServer) CreateKnowledgeBase(context.Context, *CreateKnowledgeBaseRequest) (*CreateKnowledgeBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKnowledgeBase not implemented")
}
func (UnimplementedKnowledgeBaseServiceServer) GetKnowledgeBaseById(context.Context, *GetKnowledgeBaseByIdRequest) (*GetKnowledgeBaseByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnowledgeBaseById not implemented")
}
func (UnimplementedKnowledgeBaseServiceServer) DeleteKnowledgeBaseById(context.Context, *DeleteKnowledgeBaseByIdRequest) (*DeleteKnowledgeBaseByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKnowledgeBaseById not implemented")
}
func (UnimplementedKnowledgeBaseServiceServer) mustEmbedUnimplementedKnowledgeBaseServiceServer() {}

// UnsafeKnowledgeBaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KnowledgeBaseServiceServer will
// result in compilation errors.
type UnsafeKnowledgeBaseServiceServer interface {
	mustEmbedUnimplementedKnowledgeBaseServiceServer()
}

func RegisterKnowledgeBaseServiceServer(s grpc.ServiceRegistrar, srv KnowledgeBaseServiceServer) {
	s.RegisterService(&KnowledgeBaseService_ServiceDesc, srv)
}

func _KnowledgeBaseService_CreateKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBaseServiceServer).CreateKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnowledgeBaseService_CreateKnowledgeBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBaseServiceServer).CreateKnowledgeBase(ctx, req.(*CreateKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBaseService_GetKnowledgeBaseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKnowledgeBaseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBaseServiceServer).GetKnowledgeBaseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnowledgeBaseService_GetKnowledgeBaseById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBaseServiceServer).GetKnowledgeBaseById(ctx, req.(*GetKnowledgeBaseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBaseService_DeleteKnowledgeBaseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKnowledgeBaseByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBaseServiceServer).DeleteKnowledgeBaseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnowledgeBaseService_DeleteKnowledgeBaseById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBaseServiceServer).DeleteKnowledgeBaseById(ctx, req.(*DeleteKnowledgeBaseByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KnowledgeBaseService_ServiceDesc is the grpc.ServiceDesc for KnowledgeBaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KnowledgeBaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "knowdateapp.knowledge.v1.knowledge_base.KnowledgeBaseService",
	HandlerType: (*KnowledgeBaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKnowledgeBase",
			Handler:    _KnowledgeBaseService_CreateKnowledgeBase_Handler,
		},
		{
			MethodName: "GetKnowledgeBaseById",
			Handler:    _KnowledgeBaseService_GetKnowledgeBaseById_Handler,
		},
		{
			MethodName: "DeleteKnowledgeBaseById",
			Handler:    _KnowledgeBaseService_DeleteKnowledgeBaseById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/knowledge_base/knowledge_base.proto",
}
