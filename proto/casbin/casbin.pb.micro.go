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

// Api Endpoints for CasBin service

func NewCasBinEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CasBin service

type CasBinService interface {
	AddPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeletePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdatePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Validate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type casBinService struct {
	c    client.Client
	name string
}

func NewCasBinService(name string, c client.Client) CasBinService {
	return &casBinService{
		c:    c,
		name: name,
	}
}

func (c *casBinService) AddPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.AddPermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) DeletePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.DeletePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) UpdatePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.UpdatePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) GetPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.GetPermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) AddRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.AddRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) DeleteRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.DeleteRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) UpdateRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.UpdateRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) GetRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.GetRoles", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casBinService) Validate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "CasBin.Validate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CasBin service

type CasBinHandler interface {
	AddPermission(context.Context, *Request, *Response) error
	DeletePermission(context.Context, *Request, *Response) error
	UpdatePermission(context.Context, *Request, *Response) error
	GetPermission(context.Context, *Request, *Response) error
	AddRole(context.Context, *Request, *Response) error
	DeleteRole(context.Context, *Request, *Response) error
	UpdateRole(context.Context, *Request, *Response) error
	GetRoles(context.Context, *Request, *Response) error
	Validate(context.Context, *Request, *Response) error
}

func RegisterCasBinHandler(s server.Server, hdlr CasBinHandler, opts ...server.HandlerOption) error {
	type casBin interface {
		AddPermission(ctx context.Context, in *Request, out *Response) error
		DeletePermission(ctx context.Context, in *Request, out *Response) error
		UpdatePermission(ctx context.Context, in *Request, out *Response) error
		GetPermission(ctx context.Context, in *Request, out *Response) error
		AddRole(ctx context.Context, in *Request, out *Response) error
		DeleteRole(ctx context.Context, in *Request, out *Response) error
		UpdateRole(ctx context.Context, in *Request, out *Response) error
		GetRoles(ctx context.Context, in *Request, out *Response) error
		Validate(ctx context.Context, in *Request, out *Response) error
	}
	type CasBin struct {
		casBin
	}
	h := &casBinHandler{hdlr}
	return s.Handle(s.NewHandler(&CasBin{h}, opts...))
}

type casBinHandler struct {
	CasBinHandler
}

func (h *casBinHandler) AddPermission(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.AddPermission(ctx, in, out)
}

func (h *casBinHandler) DeletePermission(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.DeletePermission(ctx, in, out)
}

func (h *casBinHandler) UpdatePermission(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.UpdatePermission(ctx, in, out)
}

func (h *casBinHandler) GetPermission(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.GetPermission(ctx, in, out)
}

func (h *casBinHandler) AddRole(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.AddRole(ctx, in, out)
}

func (h *casBinHandler) DeleteRole(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.DeleteRole(ctx, in, out)
}

func (h *casBinHandler) UpdateRole(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.UpdateRole(ctx, in, out)
}

func (h *casBinHandler) GetRoles(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.GetRoles(ctx, in, out)
}

func (h *casBinHandler) Validate(ctx context.Context, in *Request, out *Response) error {
	return h.CasBinHandler.Validate(ctx, in, out)
}