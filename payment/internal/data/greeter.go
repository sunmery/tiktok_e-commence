package data

import (
	"context"
	"payment/internal/biz"
	"payment/internal/data/models"
	"payment/internal/pkg"

	"github.com/go-kratos/kratos/v2/log"
)

type paymentRepo struct {
	data *Data
	log  *log.Helper
}

func (p *paymentRepo) Charge(ctx context.Context, req *biz.ChargeReq) (*biz.ChargeResp, error) {
	params := models.CreatePaymentParams{
		SnowflakeID:               pkg.SnowflakeID().Int64(),
		Owner:                     req.Owner,
		Name:                      req.Name,
		Amount:                    req.Amount,
		OrderID:                   int32(req.OrderId),
		CreditCardNumber:          req.CreditCard.Number,
		CreditCardCvv:             req.CreditCard.Cvv,
		CreditCardExpirationYear:  req.CreditCard.ExpirationYear,
		CreditCardExpirationMonth: req.CreditCard.ExpirationMonth,
		Status:                    biz.PENDING,
	}
	
	payment, err := p.data.db.CreatePayment(ctx, params)
	if err != nil {
		return nil, err
	}

	return &biz.ChargeResp{TransactionId: uint32(payment.SnowflakeID)}, nil
}

func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &paymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
