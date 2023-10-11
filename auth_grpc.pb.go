// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: auth.proto

package mainflux

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthzService_Authorize_FullMethodName = "/mainflux.AuthzService/Authorize"
)

// AuthzServiceClient is the client API for AuthzService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzServiceClient interface {
	// Authorize checks if the subject is authorized to perform
	// the action on the object.
	Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeRes, error)
}

type authzServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzServiceClient(cc grpc.ClientConnInterface) AuthzServiceClient {
	return &authzServiceClient{cc}
}

func (c *authzServiceClient) Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeRes, error) {
	out := new(AuthorizeRes)
	err := c.cc.Invoke(ctx, AuthzService_Authorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzServiceServer is the server API for AuthzService service.
// All implementations must embed UnimplementedAuthzServiceServer
// for forward compatibility
type AuthzServiceServer interface {
	// Authorize checks if the subject is authorized to perform
	// the action on the object.
	Authorize(context.Context, *AuthorizeReq) (*AuthorizeRes, error)
	mustEmbedUnimplementedAuthzServiceServer()
}

// UnimplementedAuthzServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthzServiceServer struct {
}

func (UnimplementedAuthzServiceServer) Authorize(context.Context, *AuthorizeReq) (*AuthorizeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedAuthzServiceServer) mustEmbedUnimplementedAuthzServiceServer() {}

// UnsafeAuthzServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzServiceServer will
// result in compilation errors.
type UnsafeAuthzServiceServer interface {
	mustEmbedUnimplementedAuthzServiceServer()
}

func RegisterAuthzServiceServer(s grpc.ServiceRegistrar, srv AuthzServiceServer) {
	s.RegisterService(&AuthzService_ServiceDesc, srv)
}

func _AuthzService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthzService_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzServiceServer).Authorize(ctx, req.(*AuthorizeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthzService_ServiceDesc is the grpc.ServiceDesc for AuthzService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthzService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mainflux.AuthzService",
	HandlerType: (*AuthzServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _AuthzService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

const (
	AuthService_Issue_FullMethodName           = "/mainflux.AuthService/Issue"
	AuthService_Login_FullMethodName           = "/mainflux.AuthService/Login"
	AuthService_Refresh_FullMethodName         = "/mainflux.AuthService/Refresh"
	AuthService_Identify_FullMethodName        = "/mainflux.AuthService/Identify"
	AuthService_Authorize_FullMethodName       = "/mainflux.AuthService/Authorize"
	AuthService_AddPolicy_FullMethodName       = "/mainflux.AuthService/AddPolicy"
	AuthService_DeletePolicy_FullMethodName    = "/mainflux.AuthService/DeletePolicy"
	AuthService_ListObjects_FullMethodName     = "/mainflux.AuthService/ListObjects"
	AuthService_ListAllObjects_FullMethodName  = "/mainflux.AuthService/ListAllObjects"
	AuthService_CountObjects_FullMethodName    = "/mainflux.AuthService/CountObjects"
	AuthService_ListSubjects_FullMethodName    = "/mainflux.AuthService/ListSubjects"
	AuthService_ListAllSubjects_FullMethodName = "/mainflux.AuthService/ListAllSubjects"
	AuthService_CountSubjects_FullMethodName   = "/mainflux.AuthService/CountSubjects"
	AuthService_Assign_FullMethodName          = "/mainflux.AuthService/Assign"
	AuthService_Members_FullMethodName         = "/mainflux.AuthService/Members"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Issue(ctx context.Context, in *IssueReq, opts ...grpc.CallOption) (*Token, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Token, error)
	Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*Token, error)
	Identify(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityRes, error)
	Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeRes, error)
	AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*AddPolicyRes, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyReq, opts ...grpc.CallOption) (*DeletePolicyRes, error)
	ListObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error)
	ListAllObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error)
	CountObjects(ctx context.Context, in *CountObjectsReq, opts ...grpc.CallOption) (*CountObjectsRes, error)
	ListSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error)
	ListAllSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error)
	CountSubjects(ctx context.Context, in *CountSubjectsReq, opts ...grpc.CallOption) (*CountSubjectsRes, error)
	Assign(ctx context.Context, in *Assignment, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Members(ctx context.Context, in *MembersReq, opts ...grpc.CallOption) (*MembersRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Issue(ctx context.Context, in *IssueReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, AuthService_Issue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, AuthService_Refresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Identify(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityRes, error) {
	out := new(IdentityRes)
	err := c.cc.Invoke(ctx, AuthService_Identify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeRes, error) {
	out := new(AuthorizeRes)
	err := c.cc.Invoke(ctx, AuthService_Authorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*AddPolicyRes, error) {
	out := new(AddPolicyRes)
	err := c.cc.Invoke(ctx, AuthService_AddPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyReq, opts ...grpc.CallOption) (*DeletePolicyRes, error) {
	out := new(DeletePolicyRes)
	err := c.cc.Invoke(ctx, AuthService_DeletePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error) {
	out := new(ListObjectsRes)
	err := c.cc.Invoke(ctx, AuthService_ListObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListAllObjects(ctx context.Context, in *ListObjectsReq, opts ...grpc.CallOption) (*ListObjectsRes, error) {
	out := new(ListObjectsRes)
	err := c.cc.Invoke(ctx, AuthService_ListAllObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CountObjects(ctx context.Context, in *CountObjectsReq, opts ...grpc.CallOption) (*CountObjectsRes, error) {
	out := new(CountObjectsRes)
	err := c.cc.Invoke(ctx, AuthService_CountObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error) {
	out := new(ListSubjectsRes)
	err := c.cc.Invoke(ctx, AuthService_ListSubjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListAllSubjects(ctx context.Context, in *ListSubjectsReq, opts ...grpc.CallOption) (*ListSubjectsRes, error) {
	out := new(ListSubjectsRes)
	err := c.cc.Invoke(ctx, AuthService_ListAllSubjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CountSubjects(ctx context.Context, in *CountSubjectsReq, opts ...grpc.CallOption) (*CountSubjectsRes, error) {
	out := new(CountSubjectsRes)
	err := c.cc.Invoke(ctx, AuthService_CountSubjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Assign(ctx context.Context, in *Assignment, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_Assign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Members(ctx context.Context, in *MembersReq, opts ...grpc.CallOption) (*MembersRes, error) {
	out := new(MembersRes)
	err := c.cc.Invoke(ctx, AuthService_Members_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Issue(context.Context, *IssueReq) (*Token, error)
	Login(context.Context, *LoginReq) (*Token, error)
	Refresh(context.Context, *RefreshReq) (*Token, error)
	Identify(context.Context, *IdentityReq) (*IdentityRes, error)
	Authorize(context.Context, *AuthorizeReq) (*AuthorizeRes, error)
	AddPolicy(context.Context, *AddPolicyReq) (*AddPolicyRes, error)
	DeletePolicy(context.Context, *DeletePolicyReq) (*DeletePolicyRes, error)
	ListObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error)
	ListAllObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error)
	CountObjects(context.Context, *CountObjectsReq) (*CountObjectsRes, error)
	ListSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error)
	ListAllSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error)
	CountSubjects(context.Context, *CountSubjectsReq) (*CountSubjectsRes, error)
	Assign(context.Context, *Assignment) (*emptypb.Empty, error)
	Members(context.Context, *MembersReq) (*MembersRes, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Issue(context.Context, *IssueReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Refresh(context.Context, *RefreshReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthServiceServer) Identify(context.Context, *IdentityReq) (*IdentityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Identify not implemented")
}
func (UnimplementedAuthServiceServer) Authorize(context.Context, *AuthorizeReq) (*AuthorizeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedAuthServiceServer) AddPolicy(context.Context, *AddPolicyReq) (*AddPolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPolicy not implemented")
}
func (UnimplementedAuthServiceServer) DeletePolicy(context.Context, *DeletePolicyReq) (*DeletePolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedAuthServiceServer) ListObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (UnimplementedAuthServiceServer) ListAllObjects(context.Context, *ListObjectsReq) (*ListObjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllObjects not implemented")
}
func (UnimplementedAuthServiceServer) CountObjects(context.Context, *CountObjectsReq) (*CountObjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountObjects not implemented")
}
func (UnimplementedAuthServiceServer) ListSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubjects not implemented")
}
func (UnimplementedAuthServiceServer) ListAllSubjects(context.Context, *ListSubjectsReq) (*ListSubjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllSubjects not implemented")
}
func (UnimplementedAuthServiceServer) CountSubjects(context.Context, *CountSubjectsReq) (*CountSubjectsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountSubjects not implemented")
}
func (UnimplementedAuthServiceServer) Assign(context.Context, *Assignment) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (UnimplementedAuthServiceServer) Members(context.Context, *MembersReq) (*MembersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Members not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Issue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Issue(ctx, req.(*IssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Refresh(ctx, req.(*RefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Identify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Identify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Identify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Identify(ctx, req.(*IdentityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authorize(ctx, req.(*AuthorizeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AddPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddPolicy(ctx, req.(*AddPolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeletePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeletePolicy(ctx, req.(*DeletePolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListObjects(ctx, req.(*ListObjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListAllObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListAllObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListAllObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListAllObjects(ctx, req.(*ListObjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CountObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountObjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CountObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CountObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CountObjects(ctx, req.(*CountObjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListSubjects(ctx, req.(*ListSubjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListAllSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListAllSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListAllSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListAllSubjects(ctx, req.(*ListSubjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CountSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountSubjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CountSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CountSubjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CountSubjects(ctx, req.(*CountSubjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Assignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Assign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Assign(ctx, req.(*Assignment))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Members_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Members(ctx, req.(*MembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mainflux.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issue",
			Handler:    _AuthService_Issue_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AuthService_Refresh_Handler,
		},
		{
			MethodName: "Identify",
			Handler:    _AuthService_Identify_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _AuthService_Authorize_Handler,
		},
		{
			MethodName: "AddPolicy",
			Handler:    _AuthService_AddPolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _AuthService_DeletePolicy_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _AuthService_ListObjects_Handler,
		},
		{
			MethodName: "ListAllObjects",
			Handler:    _AuthService_ListAllObjects_Handler,
		},
		{
			MethodName: "CountObjects",
			Handler:    _AuthService_CountObjects_Handler,
		},
		{
			MethodName: "ListSubjects",
			Handler:    _AuthService_ListSubjects_Handler,
		},
		{
			MethodName: "ListAllSubjects",
			Handler:    _AuthService_ListAllSubjects_Handler,
		},
		{
			MethodName: "CountSubjects",
			Handler:    _AuthService_CountSubjects_Handler,
		},
		{
			MethodName: "Assign",
			Handler:    _AuthService_Assign_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _AuthService_Members_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
