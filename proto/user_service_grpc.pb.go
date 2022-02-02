// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user_service.proto

package user_service

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

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorClient interface {
	GetUserByName(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	GetMockUserData(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
}

type orchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) GetUserByName(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.Orchestrator/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetMockUserData(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.Orchestrator/GetMockUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.Orchestrator/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
// All implementations must embed UnimplementedOrchestratorServer
// for forward compatibility
type OrchestratorServer interface {
	GetUserByName(context.Context, *UserRequest) (*User, error)
	GetMockUserData(context.Context, *UserRequest) (*User, error)
	GetUser(context.Context, *UserRequest) (*User, error)
	mustEmbedUnimplementedOrchestratorServer()
}

// UnimplementedOrchestratorServer must be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServer struct {
}

func (UnimplementedOrchestratorServer) GetUserByName(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedOrchestratorServer) GetMockUserData(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMockUserData not implemented")
}
func (UnimplementedOrchestratorServer) GetUser(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}

// UnsafeOrchestratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServer will
// result in compilation errors.
type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}

func _Orchestrator_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.Orchestrator/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetUserByName(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetMockUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetMockUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.Orchestrator/GetMockUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetMockUserData(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.Orchestrator/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orchestrator_ServiceDesc is the grpc.ServiceDesc for Orchestrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByName",
			Handler:    _Orchestrator_GetUserByName_Handler,
		},
		{
			MethodName: "GetMockUserData",
			Handler:    _Orchestrator_GetMockUserData_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Orchestrator_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
