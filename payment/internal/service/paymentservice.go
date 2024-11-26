package service

import (
	"context"
	"payment/internal/biz"

	pb "payment/api/payment/v1"
)

type PaymentServiceService struct {
	pb.UnimplementedPaymentServiceServer
	pc *biz.PaymentUsecase
}

func NewPaymentServiceService(pc *biz.PaymentUsecase) *PaymentServiceService {
	return &PaymentServiceService{pc: pc}
}

func (s *PaymentServiceService) Charge(ctx context.Context, req *pb.ChargeReq) (*pb.ChargeResp, error) {
	charge, err := s.pc.Charge(ctx, &biz.ChargeReq{
		Amount: req.Amount,
		CreditCardInfo: biz.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: req.OrderId,
		Owner:   req.Owner,
		Name:    req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ChargeResp{
		TransactionId: charge.TransactionId,
	}, nil
}
