// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.4
// source: user/user.proto

package user

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

const OperationUserLogin = "/user.User/Login"
const OperationUserRegister = "/user.User/Register"

type UserHTTPServer interface {
	Login(context.Context, *UserRegisterLoginRequest) (*UserRegisterLoginResponse, error)
	// Register Sends a greeting
	Register(context.Context, *UserRegisterLoginRequest) (*UserRegisterLoginResponse, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/douyin/user/register", _User_Register0_HTTP_Handler(srv))
	r.POST("/douyin/user/login", _User_Login0_HTTP_Handler(srv))
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*UserRegisterLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRegisterLoginResponse)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRegisterLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*UserRegisterLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserRegisterLoginResponse)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	Login(ctx context.Context, req *UserRegisterLoginRequest, opts ...http.CallOption) (rsp *UserRegisterLoginResponse, err error)
	Register(ctx context.Context, req *UserRegisterLoginRequest, opts ...http.CallOption) (rsp *UserRegisterLoginResponse, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *UserRegisterLoginRequest, opts ...http.CallOption) (*UserRegisterLoginResponse, error) {
	var out UserRegisterLoginResponse
	pattern := "/douyin/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *UserRegisterLoginRequest, opts ...http.CallOption) (*UserRegisterLoginResponse, error) {
	var out UserRegisterLoginResponse
	pattern := "/douyin/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
