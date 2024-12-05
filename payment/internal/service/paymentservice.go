package service

import (
	"context"
	"errors"
	"fmt"
	pb "payment/api/payment/v1"
	"payment/internal/biz"
	"payment/internal/pkg/token"
)

type PaymentServiceService struct {
	pb.UnimplementedPaymentServiceServer
	pc *biz.PaymentUsecase
}

func NewPaymentServiceService(pc *biz.PaymentUsecase) *PaymentServiceService {
	return &PaymentServiceService{pc: pc}
}
func (s *PaymentServiceService) Charge(ctx context.Context, req *pb.ChargeReq) (*pb.ChargeResp, error) {
	fmt.Printf("Received request: %+v\n", req)

	// 检查 req 是否为 nil
	if req == nil {
		return nil, errors.New("request is nil")
	}

	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}

	if payload == nil {
		return nil, errors.New("invalid token: payload is nil")
	}

	fmt.Printf("payload: %+v\n", payload)

	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	charge, cErr := s.pc.Charge(ctx, &biz.ChargeReq{
		Amount: req.Amount,
		CreditCard: &biz.CreditCardInfo{
			Number:          req.CreditCard.Number,
			Cvv:             req.CreditCard.Cvv,
			ExpirationYear:  req.CreditCard.ExpirationYear,
			ExpirationMonth: req.CreditCard.ExpirationMonth,
		},

		OrderId: req.OrderId,
		Owner:   payload.Owner,
		Name:    payload.Name,
	})
	if cErr != nil {
		return nil, cErr
	}

	return &pb.ChargeResp{
		TransactionId: charge.TransactionId,
	}, nil
}
