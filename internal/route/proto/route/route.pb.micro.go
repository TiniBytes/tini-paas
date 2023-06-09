// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/route/route.proto

package route

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

// Api Endpoints for Route service

func NewRouteEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Route service

type RouteService interface {
	AddRoute(ctx context.Context, in *RouteInfo, opts ...client.CallOption) (*Response, error)
	DeleteRoute(ctx context.Context, in *RouteID, opts ...client.CallOption) (*Response, error)
	UpdateRoute(ctx context.Context, in *RouteInfo, opts ...client.CallOption) (*Response, error)
	FindRouteByID(ctx context.Context, in *RouteID, opts ...client.CallOption) (*RouteInfo, error)
	FindAllRoute(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllRoute, error)
}

type routeService struct {
	c    client.Client
	name string
}

func NewRouteService(name string, c client.Client) RouteService {
	return &routeService{
		c:    c,
		name: name,
	}
}

func (c *routeService) AddRoute(ctx context.Context, in *RouteInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Route.AddRoute", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeService) DeleteRoute(ctx context.Context, in *RouteID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Route.DeleteRoute", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeService) UpdateRoute(ctx context.Context, in *RouteInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Route.UpdateRoute", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeService) FindRouteByID(ctx context.Context, in *RouteID, opts ...client.CallOption) (*RouteInfo, error) {
	req := c.c.NewRequest(c.name, "Route.FindRouteByID", in)
	out := new(RouteInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeService) FindAllRoute(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllRoute, error) {
	req := c.c.NewRequest(c.name, "Route.FindAllRoute", in)
	out := new(AllRoute)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Route service

type RouteHandler interface {
	AddRoute(context.Context, *RouteInfo, *Response) error
	DeleteRoute(context.Context, *RouteID, *Response) error
	UpdateRoute(context.Context, *RouteInfo, *Response) error
	FindRouteByID(context.Context, *RouteID, *RouteInfo) error
	FindAllRoute(context.Context, *FindAll, *AllRoute) error
}

func RegisterRouteHandler(s server.Server, hdlr RouteHandler, opts ...server.HandlerOption) error {
	type route interface {
		AddRoute(ctx context.Context, in *RouteInfo, out *Response) error
		DeleteRoute(ctx context.Context, in *RouteID, out *Response) error
		UpdateRoute(ctx context.Context, in *RouteInfo, out *Response) error
		FindRouteByID(ctx context.Context, in *RouteID, out *RouteInfo) error
		FindAllRoute(ctx context.Context, in *FindAll, out *AllRoute) error
	}
	type Route struct {
		route
	}
	h := &routeHandler{hdlr}
	return s.Handle(s.NewHandler(&Route{h}, opts...))
}

type routeHandler struct {
	RouteHandler
}

func (h *routeHandler) AddRoute(ctx context.Context, in *RouteInfo, out *Response) error {
	return h.RouteHandler.AddRoute(ctx, in, out)
}

func (h *routeHandler) DeleteRoute(ctx context.Context, in *RouteID, out *Response) error {
	return h.RouteHandler.DeleteRoute(ctx, in, out)
}

func (h *routeHandler) UpdateRoute(ctx context.Context, in *RouteInfo, out *Response) error {
	return h.RouteHandler.UpdateRoute(ctx, in, out)
}

func (h *routeHandler) FindRouteByID(ctx context.Context, in *RouteID, out *RouteInfo) error {
	return h.RouteHandler.FindRouteByID(ctx, in, out)
}

func (h *routeHandler) FindAllRoute(ctx context.Context, in *FindAll, out *AllRoute) error {
	return h.RouteHandler.FindAllRoute(ctx, in, out)
}
