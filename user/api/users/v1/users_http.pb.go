// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.28.3
// source: users/v1/users.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceGetUserInfo = "/api.user.v1.UserService/GetUserInfo"
const OperationUserServiceSignin = "/api.user.v1.UserService/Signin"

type UserServiceHTTPServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	Signin(context.Context, *SigninRequest) (*SigninReply, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/users", _UserService_Signin0_HTTP_Handler(srv))
	r.GET("/v1/users/profile", _UserService_GetUserInfo0_HTTP_Handler(srv))
}

func _UserService_Signin0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SigninRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceSignin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Signin(ctx, req.(*SigninRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SigninReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserInfo0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoResponse, err error)
	Signin(ctx context.Context, req *SigninRequest, opts ...http.CallOption) (rsp *SigninReply, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoResponse, error) {
	var out GetUserInfoResponse
	pattern := "/v1/users/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) Signin(ctx context.Context, in *SigninRequest, opts ...http.CallOption) (*SigninReply, error) {
	var out SigninReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceSignin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
