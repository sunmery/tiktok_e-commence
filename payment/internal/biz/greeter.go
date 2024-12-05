package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CreditCardInfo struct {
	Number          string `json:"credit_card_number"`
	Cvv             int32  `json:"credit_card_cvv"`
	ExpirationYear  int32  `json:"credit_card_expiration_year"`
	ExpirationMonth int32  `json:"credit_card_expiration_month"`
}

type ChargeReq struct {
	Amount     float32         `json:"amount"`
	CreditCard *CreditCardInfo `json:"credit_card"`
	OrderId    uint32          `json:"order_id"`
	Owner      string          `json:"owner"`
	Name       string          `json:"name"`
}

type ChargeResp struct {
	TransactionId uint32 `json:"transaction_id"`
}

// PaymentRepo is a Greater repo.
type PaymentRepo interface {
	Charge(ctx context.Context, req *ChargeReq) (*ChargeResp, error)
}

// PaymentUsecase is a Payment usecase.
type PaymentUsecase struct {
	repo PaymentRepo
	log  *log.Helper
}

// NewPaymentUsecase new a Payment usecase.
func NewPaymentUsecase(repo PaymentRepo, logger log.Logger) *PaymentUsecase {
	return &PaymentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PaymentUsecase) Charge(ctx context.Context, req *ChargeReq) (*ChargeResp, error) {
	uc.log.WithContext(ctx).Infof("Charge: %v", req)
	return uc.repo.Charge(ctx, req)
}
