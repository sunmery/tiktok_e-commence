package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// 支付状态
const (
	PENDING = "pending"
	SUCCESS = "success"
	CANCEL  = "cancel"
)

type CreditCardInfo struct {
	Number          string `json:"credit_card_number"`
	Cvv             string `json:"credit_card_cvv"`
	ExpirationYear  string `json:"credit_card_expiration_year"`
	ExpirationMonth string `json:"credit_card_expiration_month"`
}

type ChargeReq struct {
	Amount     float64         `json:"amount"`
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
