// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/userApi/userApi.proto

package userApi

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserApi service

func NewUserApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserApi service

type UserApiService interface {
	AddUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	FindUserByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	IsRight(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeletePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdatePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type userApiService struct {
	c    client.Client
	name string
}

func NewUserApiService(name string, c client.Client) UserApiService {
	return &userApiService{
		c:    c,
		name: name,
	}
}

func (c *userApiService) AddUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.AddUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) DeleteUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.DeleteUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) UpdateUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.UpdateUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) FindUserByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.FindUserByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) AddRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.AddRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) DeleteRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.DeleteRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) UpdateRole(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.UpdateRole", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) IsRight(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.IsRight", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) AddPermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.AddPermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) DeletePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.DeletePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiService) UpdatePermission(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserApi.UpdatePermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserApi service

type UserApiHandler interface {
	AddUser(context.Context, *Request, *Response) error
	DeleteUser(context.Context, *Request, *Response) error
	UpdateUser(context.Context, *Request, *Response) error
	FindUserByID(context.Context, *Request, *Response) error
	Call(context.Context, *Request, *Response) error
	AddRole(context.Context, *Request, *Response) error
	DeleteRole(context.Context, *Request, *Response) error
	UpdateRole(context.Context, *Request, *Response) error
	IsRight(context.Context, *Request, *Response) error
	AddPermission(context.Context, *Request, *Response) error
	DeletePermission(context.Context, *Request, *Response) error
	UpdatePermission(context.Context, *Request, *Response) error
}

func RegisterUserApiHandler(s server.Server, hdlr UserApiHandler, opts ...server.HandlerOption) error {
	type userApi interface {
		AddUser(ctx context.Context, in *Request, out *Response) error
		DeleteUser(ctx context.Context, in *Request, out *Response) error
		UpdateUser(ctx context.Context, in *Request, out *Response) error
		FindUserByID(ctx context.Context, in *Request, out *Response) error
		Call(ctx context.Context, in *Request, out *Response) error
		AddRole(ctx context.Context, in *Request, out *Response) error
		DeleteRole(ctx context.Context, in *Request, out *Response) error
		UpdateRole(ctx context.Context, in *Request, out *Response) error
		IsRight(ctx context.Context, in *Request, out *Response) error
		AddPermission(ctx context.Context, in *Request, out *Response) error
		DeletePermission(ctx context.Context, in *Request, out *Response) error
		UpdatePermission(ctx context.Context, in *Request, out *Response) error
	}
	type UserApi struct {
		userApi
	}
	h := &userApiHandler{hdlr}
	return s.Handle(s.NewHandler(&UserApi{h}, opts...))
}

type userApiHandler struct {
	UserApiHandler
}

func (h *userApiHandler) AddUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.AddUser(ctx, in, out)
}

func (h *userApiHandler) DeleteUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.DeleteUser(ctx, in, out)
}

func (h *userApiHandler) UpdateUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.UpdateUser(ctx, in, out)
}

func (h *userApiHandler) FindUserByID(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.FindUserByID(ctx, in, out)
}

func (h *userApiHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.Call(ctx, in, out)
}

func (h *userApiHandler) AddRole(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.AddRole(ctx, in, out)
}

func (h *userApiHandler) DeleteRole(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.DeleteRole(ctx, in, out)
}

func (h *userApiHandler) UpdateRole(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.UpdateRole(ctx, in, out)
}

func (h *userApiHandler) IsRight(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.IsRight(ctx, in, out)
}

func (h *userApiHandler) AddPermission(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.AddPermission(ctx, in, out)
}

func (h *userApiHandler) DeletePermission(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.DeletePermission(ctx, in, out)
}

func (h *userApiHandler) UpdatePermission(ctx context.Context, in *Request, out *Response) error {
	return h.UserApiHandler.UpdatePermission(ctx, in, out)
}