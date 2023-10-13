// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ory/keto/relation_tuples/v1alpha2/expand_service.proto

package rts

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

// ExpandServiceClient is the client API for ExpandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpandServiceClient interface {
	// Expands the subject set into a tree of subjects.
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error)
}

type expandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpandServiceClient(cc grpc.ClientConnInterface) ExpandServiceClient {
	return &expandServiceClient{cc}
}

func (c *expandServiceClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error) {
	out := new(ExpandResponse)
	err := c.cc.Invoke(ctx, "/ory.keto.relation_tuples.v1alpha2.ExpandService/Expand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpandServiceServer is the server API for ExpandService service.
// All implementations should embed UnimplementedExpandServiceServer
// for forward compatibility
type ExpandServiceServer interface {
	// Expands the subject set into a tree of subjects.
	Expand(context.Context, *ExpandRequest) (*ExpandResponse, error)
}

// UnimplementedExpandServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExpandServiceServer struct {
}

func (UnimplementedExpandServiceServer) Expand(context.Context, *ExpandRequest) (*ExpandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}

// UnsafeExpandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpandServiceServer will
// result in compilation errors.
type UnsafeExpandServiceServer interface {
	mustEmbedUnimplementedExpandServiceServer()
}

func RegisterExpandServiceServer(s grpc.ServiceRegistrar, srv ExpandServiceServer) {
	s.RegisterService(&ExpandService_ServiceDesc, srv)
}

func _ExpandService_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpandServiceServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ory.keto.relation_tuples.v1alpha2.ExpandService/Expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpandServiceServer).Expand(ctx, req.(*ExpandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpandService_ServiceDesc is the grpc.ServiceDesc for ExpandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ory.keto.relation_tuples.v1alpha2.ExpandService",
	HandlerType: (*ExpandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Expand",
			Handler:    _ExpandService_Expand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ory/keto/relation_tuples/v1alpha2/expand_service.proto",
}
