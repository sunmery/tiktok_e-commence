// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"context"
)

type Querier interface {
	//CreatePayment
	//
	//  INSERT INTO payment.payments(snowflake_id, owner, name, amount, order_id, credit_card_number, credit_card_cvv,
	//                               credit_card_expiration_year, credit_card_expiration_month, status)
	//  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	//  RETURNING id, snowflake_id, owner, name, amount, order_id, credit_card_number, credit_card_cvv, credit_card_expiration_year, credit_card_expiration_month, status, created_at
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (PaymentPayments, error)
}

var _ Querier = (*Queries)(nil)