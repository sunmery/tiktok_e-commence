// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: balance/v1/balance.proto

package v1

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
	Balance_CreateBalance_FullMethodName = "/api.balance.v1.Balance/CreateBalance"
	Balance_UpdateBalance_FullMethodName = "/api.balance.v1.Balance/UpdateBalance"
	Balance_GetBalance_FullMethodName    = "/api.balance.v1.Balance/GetBalance"
)

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceClient interface {
	CreateBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error)
	UpdateBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error)
}

type balanceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceClient(cc grpc.ClientConnInterface) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) CreateBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceReply)
	err := c.cc.Invoke(ctx, Balance_CreateBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceClient) UpdateBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceReply)
	err := c.cc.Invoke(ctx, Balance_UpdateBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*BalanceReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BalanceReply)
	err := c.cc.Invoke(ctx, Balance_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
// All implementations must embed UnimplementedBalanceServer
// for forward compatibility.
type BalanceServer interface {
	CreateBalance(context.Context, *BalanceRequest) (*BalanceReply, error)
	UpdateBalance(context.Context, *BalanceRequest) (*BalanceReply, error)
	GetBalance(context.Context, *GetBalanceRequest) (*BalanceReply, error)
	mustEmbedUnimplementedBalanceServer()
}

// UnimplementedBalanceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBalanceServer struct{}

func (UnimplementedBalanceServer) CreateBalance(context.Context, *BalanceRequest) (*BalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBalance not implemented")
}
func (UnimplementedBalanceServer) UpdateBalance(context.Context, *BalanceRequest) (*BalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedBalanceServer) GetBalance(context.Context, *GetBalanceRequest) (*BalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBalanceServer) mustEmbedUnimplementedBalanceServer() {}
func (UnimplementedBalanceServer) testEmbeddedByValue()                 {}

// UnsafeBalanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceServer will
// result in compilation errors.
type UnsafeBalanceServer interface {
	mustEmbedUnimplementedBalanceServer()
}

func RegisterBalanceServer(s grpc.ServiceRegistrar, srv BalanceServer) {
	// If the following call pancis, it indicates UnimplementedBalanceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Balance_ServiceDesc, srv)
}

func _Balance_CreateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).CreateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Balance_CreateBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).CreateBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balance_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Balance_UpdateBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).UpdateBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balance_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Balance_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Balance_ServiceDesc is the grpc.ServiceDesc for Balance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Balance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.balance.v1.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBalance",
			Handler:    _Balance_CreateBalance_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _Balance_UpdateBalance_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Balance_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "balance/v1/balance.proto",
}