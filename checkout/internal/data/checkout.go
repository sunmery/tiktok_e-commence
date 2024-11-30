package data

import (
	"checkout/internal/data/modules"
	"checkout/internal/pkg/token"
	"context"
	"errors"
	"fmt"

	"checkout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type checkoutRepo struct {
	data *Data
	log  *log.Helper
}

func (c *checkoutRepo) Checkout(ctx context.Context, req *biz.CheckoutReq) (*biz.CheckoutResp, error) {
	payload, err := token.ExtractPayload(ctx)
	if err != nil {
		return nil, err
	}
	if req.Owner != payload.Owner || req.Name != payload.Name {
		return nil, errors.New("invalid token")
	}

	checkout, err := c.data.db.CreateCheckout(ctx, models.CreateCheckoutParams{
		Owner:        payload.Owner,
		Name:         payload.Name,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Email:        req.Email,
		AddressID:    req.AddressId,
		CreditCardID: req.CreditCardId,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("checkout:%+v\n", checkout)

	return &biz.CheckoutResp{
		OrderId:       "",
		TransactionId: "",
	}, err
}

// NewCheckoutRepo .
func NewCheckoutRepo(data *Data, logger log.Logger) biz.CheckoutRepo {
	return &checkoutRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
