// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/casbin/casbin.proto

package casbin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for Casbin service

func NewCasbinEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Casbin service

type CasbinService interface {
	// 角色权限管理
	AddPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeletePermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdatePermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetPermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 用户角色管理
	AddRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 微服务内部调用
	// 验证用户权限
	Validate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type casbinService struct {
	c    client.Client
	name string
}

func NewCasbinService(name string, c client.Client) CasbinService {
	return &casbinService{
		c:    c,
		name: name,
	}
}

func (c *casbinService) AddPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.AddPermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) DeletePermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.DeletePermissions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) UpdatePermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.UpdatePermissions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) GetPermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.GetPermissions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) AddRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.AddRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) DeleteRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.DeleteRoles", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) UpdateRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.UpdateRoles", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) GetRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.GetRoles", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casbinService) Validate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Casbin.Validate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Casbin service

type CasbinHandler interface {
	// 角色权限管理
	AddPermission(context.Context, *Request, *Response) error
	DeletePermissions(context.Context, *Request, *Response) error
	UpdatePermissions(context.Context, *Request, *Response) error
	GetPermissions(context.Context, *Request, *Response) error
	// 用户角色管理
	AddRole(context.Context, *Request, *Response) error
	DeleteRoles(context.Context, *Request, *Response) error
	UpdateRoles(context.Context, *Request, *Response) error
	GetRoles(context.Context, *Request, *Response) error
	// 微服务内部调用
	// 验证用户权限
	Validate(context.Context, *Request, *Response) error
}

func RegisterCasbinHandler(s server.Server, hdlr CasbinHandler, opts ...server.HandlerOption) error {
	type casbin interface {
		AddPermission(ctx context.Context, in *Request, out *Response) error
		DeletePermissions(ctx context.Context, in *Request, out *Response) error
		UpdatePermissions(ctx context.Context, in *Request, out *Response) error
		GetPermissions(ctx context.Context, in *Request, out *Response) error
		AddRole(ctx context.Context, in *Request, out *Response) error
		DeleteRoles(ctx context.Context, in *Request, out *Response) error
		UpdateRoles(ctx context.Context, in *Request, out *Response) error
		GetRoles(ctx context.Context, in *Request, out *Response) error
		Validate(ctx context.Context, in *Request, out *Response) error
	}
	type Casbin struct {
		casbin
	}
	h := &casbinHandler{hdlr}
	return s.Handle(s.NewHandler(&Casbin{h}, opts...))
}

type casbinHandler struct {
	CasbinHandler
}

func (h *casbinHandler) AddPermission(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.AddPermission(ctx, in, out)
}

func (h *casbinHandler) DeletePermissions(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.DeletePermissions(ctx, in, out)
}

func (h *casbinHandler) UpdatePermissions(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.UpdatePermissions(ctx, in, out)
}

func (h *casbinHandler) GetPermissions(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.GetPermissions(ctx, in, out)
}

func (h *casbinHandler) AddRole(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.AddRole(ctx, in, out)
}

func (h *casbinHandler) DeleteRoles(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.DeleteRoles(ctx, in, out)
}

func (h *casbinHandler) UpdateRoles(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.UpdateRoles(ctx, in, out)
}

func (h *casbinHandler) GetRoles(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.GetRoles(ctx, in, out)
}

func (h *casbinHandler) Validate(ctx context.Context, in *Request, out *Response) error {
	return h.CasbinHandler.Validate(ctx, in, out)
}
