package service

import (
	"context"

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
	resp, err := s.cc.CreateCheckout(ctx, &biz.CheckoutReq{
		UserId:       req.UserId,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
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
