// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/account.proto

package go_micro_api_console_account

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro-in-cn/starter-kit/console/account/genproto/srv"
	proto1 "github.com/micro/go-micro/v3/api/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Account service

func NewAccountEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Account service

type AccountService interface {
	// 登录接口
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*Response, error)
	// 登出接口
	Logout(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	// Info接口
	Info(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Account.Login", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Logout(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Account.Logout", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Info(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Account.Info", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	// 登录接口
	Login(context.Context, *LoginRequest, *Response) error
	// 登出接口
	Logout(context.Context, *proto1.Request, *proto1.Response) error
	// Info接口
	Info(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		Login(ctx context.Context, in *LoginRequest, out *Response) error
		Logout(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		Info(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	return s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) Login(ctx context.Context, in *LoginRequest, out *Response) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *accountHandler) Logout(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.AccountHandler.Logout(ctx, in, out)
}

func (h *accountHandler) Info(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.AccountHandler.Info(ctx, in, out)
}
