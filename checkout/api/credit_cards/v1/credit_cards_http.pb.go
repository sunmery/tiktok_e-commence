// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.28.3
// source: credit_cards/v1/credit_cards.proto

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

const OperationCreditCardsServiceCreateCreditCard = "/api.credit_cards.v1.CreditCardsService/CreateCreditCard"
const OperationCreditCardsServiceDeleteCreditCard = "/api.credit_cards.v1.CreditCardsService/DeleteCreditCard"
const OperationCreditCardsServiceGetCreditCard = "/api.credit_cards.v1.CreditCardsService/GetCreditCard"
const OperationCreditCardsServiceListCreditCards = "/api.credit_cards.v1.CreditCardsService/ListCreditCards"
const OperationCreditCardsServiceUpdateCreditCard = "/api.credit_cards.v1.CreditCardsService/UpdateCreditCard"

type CreditCardsServiceHTTPServer interface {
	CreateCreditCard(context.Context, *CreditCards) (*CardsReply, error)
	DeleteCreditCard(context.Context, *DeleteCreditCardsRequest) (*CardsReply, error)
	GetCreditCard(context.Context, *GetCreditCardsRequest) (*GetCreditCardsReply, error)
	ListCreditCards(context.Context, *ListCreditCardsRequest) (*ListCreditCardsReply, error)
	UpdateCreditCard(context.Context, *CreditCards) (*CardsReply, error)
}

func RegisterCreditCardsServiceHTTPServer(s *http.Server, srv CreditCardsServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/credit_cards", _CreditCardsService_CreateCreditCard0_HTTP_Handler(srv))
	r.PATCH("/v1/credit_cards", _CreditCardsService_UpdateCreditCard0_HTTP_Handler(srv))
	r.DELETE("/v1/credit_cards/{id}", _CreditCardsService_DeleteCreditCard0_HTTP_Handler(srv))
	r.GET("/v1/credit_cards/{credit_card_number}", _CreditCardsService_GetCreditCard0_HTTP_Handler(srv))
	r.GET("/v1/credit_cards/all", _CreditCardsService_ListCreditCards0_HTTP_Handler(srv))
}

func _CreditCardsService_CreateCreditCard0_HTTP_Handler(srv CreditCardsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreditCards
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCreditCardsServiceCreateCreditCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCreditCard(ctx, req.(*CreditCards))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardsReply)
		return ctx.Result(200, reply)
	}
}

func _CreditCardsService_UpdateCreditCard0_HTTP_Handler(srv CreditCardsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreditCards
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCreditCardsServiceUpdateCreditCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCreditCard(ctx, req.(*CreditCards))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardsReply)
		return ctx.Result(200, reply)
	}
}

func _CreditCardsService_DeleteCreditCard0_HTTP_Handler(srv CreditCardsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCreditCardsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCreditCardsServiceDeleteCreditCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCreditCard(ctx, req.(*DeleteCreditCardsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardsReply)
		return ctx.Result(200, reply)
	}
}

func _CreditCardsService_GetCreditCard0_HTTP_Handler(srv CreditCardsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCreditCardsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCreditCardsServiceGetCreditCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCreditCard(ctx, req.(*GetCreditCardsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCreditCardsReply)
		return ctx.Result(200, reply)
	}
}

func _CreditCardsService_ListCreditCards0_HTTP_Handler(srv CreditCardsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCreditCardsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCreditCardsServiceListCreditCards)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCreditCards(ctx, req.(*ListCreditCardsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCreditCardsReply)
		return ctx.Result(200, reply)
	}
}

type CreditCardsServiceHTTPClient interface {
	CreateCreditCard(ctx context.Context, req *CreditCards, opts ...http.CallOption) (rsp *CardsReply, err error)
	DeleteCreditCard(ctx context.Context, req *DeleteCreditCardsRequest, opts ...http.CallOption) (rsp *CardsReply, err error)
	GetCreditCard(ctx context.Context, req *GetCreditCardsRequest, opts ...http.CallOption) (rsp *GetCreditCardsReply, err error)
	ListCreditCards(ctx context.Context, req *ListCreditCardsRequest, opts ...http.CallOption) (rsp *ListCreditCardsReply, err error)
	UpdateCreditCard(ctx context.Context, req *CreditCards, opts ...http.CallOption) (rsp *CardsReply, err error)
}

type CreditCardsServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCreditCardsServiceHTTPClient(client *http.Client) CreditCardsServiceHTTPClient {
	return &CreditCardsServiceHTTPClientImpl{client}
}

func (c *CreditCardsServiceHTTPClientImpl) CreateCreditCard(ctx context.Context, in *CreditCards, opts ...http.CallOption) (*CardsReply, error) {
	var out CardsReply
	pattern := "/v1/credit_cards"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCreditCardsServiceCreateCreditCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CreditCardsServiceHTTPClientImpl) DeleteCreditCard(ctx context.Context, in *DeleteCreditCardsRequest, opts ...http.CallOption) (*CardsReply, error) {
	var out CardsReply
	pattern := "/v1/credit_cards/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCreditCardsServiceDeleteCreditCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CreditCardsServiceHTTPClientImpl) GetCreditCard(ctx context.Context, in *GetCreditCardsRequest, opts ...http.CallOption) (*GetCreditCardsReply, error) {
	var out GetCreditCardsReply
	pattern := "/v1/credit_cards/{credit_card_number}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCreditCardsServiceGetCreditCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CreditCardsServiceHTTPClientImpl) ListCreditCards(ctx context.Context, in *ListCreditCardsRequest, opts ...http.CallOption) (*ListCreditCardsReply, error) {
	var out ListCreditCardsReply
	pattern := "/v1/credit_cards/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCreditCardsServiceListCreditCards))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CreditCardsServiceHTTPClientImpl) UpdateCreditCard(ctx context.Context, in *CreditCards, opts ...http.CallOption) (*CardsReply, error) {
	var out CardsReply
	pattern := "/v1/credit_cards"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCreditCardsServiceUpdateCreditCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}