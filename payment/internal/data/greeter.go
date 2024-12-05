package data

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"payment/internal/biz"
	"payment/internal/data/models"

	"github.com/go-kratos/kratos/v2/log"
)

type paymentRepo struct {
	data *Data
	log  *log.Helper
}

func (p *paymentRepo) Charge(ctx context.Context, req *biz.ChargeReq) (*biz.ChargeResp, error) {
	payment, err := p.data.db.CreatePayment(ctx, models.CreatePaymentParams{
		SnowflakeID:          0,
		Owner:                req.Owner,
		Name:                 req.Name,
		Amount:               pgtype.Numeric{},
		OrderID:              int32(req.OrderId),
		CreditCardNumber:     "",
		CreditCardExpiration: pgtype.Date{},
		Status:               "",
	})
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
