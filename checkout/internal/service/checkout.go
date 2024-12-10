package service

import (
	"checkout/internal/pkg/token"
	"context"
	"errors"
	"fmt"

	v1 "checkout/api/checkout/v1"
	"checkout/internal/biz"
)

type CheckoutService struct {
	v1.UnimplementedCheckoutServiceServer

	cc *biz.CheckoutUsecase
}

// NewCheckoutService new a Checkout service.
func NewCheckoutService(cc *biz.CheckoutUsecase) *CheckoutService {
	return &CheckoutService{cc: cc}
}

func (s *CheckoutService) Checkout(ctx context.Context, req *v1.CheckoutReq) (*v1.CheckoutResp, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("payload:", payload)
	fmt.Println("req:", req)
	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	resp, err := s.cc.CreateCheckout(ctx, &biz.CheckoutReq{
		Owner:        payload.Owner,
		Name:         payload.Name,
		Email:        req.Email,
		AddressId:    req.AddressId,
		CreditCardId: req.CreditCardId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CheckoutResp{
		OrderId:       resp.OrderId,
		TransactionId: resp.TransactionId,
	}, nil
}
