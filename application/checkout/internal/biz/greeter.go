package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
}

type CheckoutReq struct {
	UserId         string  `json:"user_id"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Email          string  `json:"email"`
	Address        Address `json:"address"`
	CreditCardInfo string  `json:"credit_card_info"`
}

type CheckoutResp struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Checkout(ctx context.Context, req *CheckoutReq) (*CheckoutResp, error)
}

// GreeterUsecase is a Checkout usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Checkout usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Checkout, and returns the new Checkout.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Checkout) (*Checkout, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
