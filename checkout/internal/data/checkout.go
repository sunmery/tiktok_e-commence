package data

import (
	creditCardsV1 "checkout/api/credit_cards/v1"
	OrderV1 "checkout/api/order/v1"
	paymentV1 "checkout/api/payment/v1"
	"checkout/internal/data/models"

	"context"
	"fmt"

	"checkout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type checkoutRepo struct {
	data *Data
	log  *log.Helper
}

func (c *checkoutRepo) Checkout(ctx context.Context, req *biz.CheckoutReq) (*biz.CheckoutResp, error) {
	checkout, err := c.data.db.CreateCheckout(ctx, models.CreateCheckoutParams{
		Owner:        req.Owner,
		Name:         req.Name,
		Email:        req.Email,
		AddressID:    req.AddressId,
		CreditCardID: req.CreditCardId,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("checkout:%+v\n", checkout)
	// orderItems := make([]*OrderV1.OrderItem)
	// for i, i2 := range collection {
	//
	// }
	order, err := c.data.orderClient.PlaceOrder(ctx, &OrderV1.PlaceOrderReq{
		Owner:     req.Owner,
		Name:      req.Name,
		Currency:  "CNY",
		AddressId: uint32(req.AddressId),
		Email:     req.Email,
		// OrderItems: []*OrderV1.OrderItem{},
		OrderItems: nil,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("order:%+v\n", order)

	cards, err := c.data.creditCardsClient.GetCreditCard(ctx, &creditCardsV1.GetCreditCardsRequest{
		Owner:            req.Owner,
		Name:             req.Name,
		CreditCardNumber: "6451-5730-2939-6899",
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("cards:%+v\n", cards)

	payment, err := c.data.paymentClient.Charge(ctx, &paymentV1.ChargeReq{
		Amount: 0,
		CreditCard: &paymentV1.CreditCardInfo{
			Number:          cards.CreditCards.CreditCardNumber,
			Cvv:             cards.CreditCards.CreditCardCvv,
			ExpirationYear:  cards.CreditCards.CreditCardExpirationYear,
			ExpirationMonth: cards.CreditCards.CreditCardExpirationMonth,
		},
		OrderId: order.Order.OrderId,
		Owner:   req.Owner,
		Name:    req.Name,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("payment:%+v\n", payment)

	return &biz.CheckoutResp{
		OrderId:       order.Order.OrderId,
		TransactionId: payment.TransactionId,
	}, nil
}

// NewCheckoutRepo .
func NewCheckoutRepo(data *Data, logger log.Logger) biz.CheckoutRepo {
	return &checkoutRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
