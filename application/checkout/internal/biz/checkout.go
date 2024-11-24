package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

//	type Address struct {
//		StreetAddress string `json:"street_address"`
//		City          string `json:"city"`
//		State         string `json:"state"`
//		Country       string `json:"country"`
//		ZipCode       string `json:"zip_code"`
//	}
//
//	type CreditCardInfo struct {
//		CreditCardNumber          string `json:"credit_card_number"`
//		CreditCardCvv             int32  `json:"credit_card_cvv"`
//		CreditCardExpirationYear  int32  `json:"credit_card_expiration_year"`
//		CreditCardExpirationMonth int32  `json:"credit_card_expiration_month"`
//	}
type CheckoutReq struct {
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	AddressId    int32  `json:"address_id"`
	CreditCardId int32  `json:"credit_card_id"`
}

type CheckoutResp struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
}

// CheckoutRepo is a Greater repo.
type CheckoutRepo interface {
	Checkout(ctx context.Context, req *CheckoutReq) (*CheckoutResp, error)
}

// CheckoutUsecase is a Checkout usecase.
type CheckoutUsecase struct {
	repo CheckoutRepo
	log  *log.Helper
}

// NewCheckoutUsecase new a Checkout usecase.
func NewCheckoutUsecase(repo CheckoutRepo, logger log.Logger) *CheckoutUsecase {
	return &CheckoutUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCheckout creates a Checkout, and returns the new Checkout.
func (uc *CheckoutUsecase) CreateCheckout(ctx context.Context, req *CheckoutReq) (*CheckoutResp, error) {
	uc.log.WithContext(ctx).Infof("CreateCheckout: %v", req)
	return uc.repo.Checkout(ctx, req)
}
