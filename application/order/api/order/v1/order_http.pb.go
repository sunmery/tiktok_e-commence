// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: order/v1/order.proto

package orderv1

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

const OperationOrderServiceListOrder = "/api.order.v1.OrderService/ListOrder"
const OperationOrderServiceMarkOrderPaid = "/api.order.v1.OrderService/MarkOrderPaid"
const OperationOrderServicePlaceOrder = "/api.order.v1.OrderService/PlaceOrder"

type OrderServiceHTTPServer interface {
	// ListOrder 列出订单
	ListOrder(context.Context, *ListOrderReq) (*ListOrderResp, error)
	// MarkOrderPaid 支付订单
	MarkOrderPaid(context.Context, *MarkOrderPaidReq) (*MarkOrderPaidResp, error)
	// PlaceOrder 创建订单
	PlaceOrder(context.Context, *PlaceOrderReq) (*PlaceOrderResp, error)
}

func RegisterOrderServiceHTTPServer(s *http.Server, srv OrderServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/order", _OrderService_PlaceOrder0_HTTP_Handler(srv))
	r.GET("/v1/order", _OrderService_ListOrder0_HTTP_Handler(srv))
	r.POST("/v1/order/mark", _OrderService_MarkOrderPaid0_HTTP_Handler(srv))
}

func _OrderService_PlaceOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlaceOrderReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServicePlaceOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PlaceOrder(ctx, req.(*PlaceOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlaceOrderResp)
		return ctx.Result(200, reply)
	}
}

func _OrderService_ListOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceListOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderResp)
		return ctx.Result(200, reply)
	}
}

func _OrderService_MarkOrderPaid0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MarkOrderPaidReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceMarkOrderPaid)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MarkOrderPaid(ctx, req.(*MarkOrderPaidReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MarkOrderPaidResp)
		return ctx.Result(200, reply)
	}
}

type OrderServiceHTTPClient interface {
	ListOrder(ctx context.Context, req *ListOrderReq, opts ...http.CallOption) (rsp *ListOrderResp, err error)
	MarkOrderPaid(ctx context.Context, req *MarkOrderPaidReq, opts ...http.CallOption) (rsp *MarkOrderPaidResp, err error)
	PlaceOrder(ctx context.Context, req *PlaceOrderReq, opts ...http.CallOption) (rsp *PlaceOrderResp, err error)
}

type OrderServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderServiceHTTPClient(client *http.Client) OrderServiceHTTPClient {
	return &OrderServiceHTTPClientImpl{client}
}

func (c *OrderServiceHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderReq, opts ...http.CallOption) (*ListOrderResp, error) {
	var out ListOrderResp
	pattern := "/v1/order"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderServiceListOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrderServiceHTTPClientImpl) MarkOrderPaid(ctx context.Context, in *MarkOrderPaidReq, opts ...http.CallOption) (*MarkOrderPaidResp, error) {
	var out MarkOrderPaidResp
	pattern := "/v1/order/mark"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderServiceMarkOrderPaid))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrderServiceHTTPClientImpl) PlaceOrder(ctx context.Context, in *PlaceOrderReq, opts ...http.CallOption) (*PlaceOrderResp, error) {
	var out PlaceOrderResp
	pattern := "/v1/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderServicePlaceOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}