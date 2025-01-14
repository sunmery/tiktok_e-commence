// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: checkout.sql

package models

import (
	"context"
)

const CreateCheckout = `-- name: CreateCheckout :one
INSERT INTO checkout.checkout(owner, name, firstname, lastname, email, address_id, credit_card_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, owner, name, firstname, lastname, email, address_id, credit_card_id
`

type CreateCheckoutParams struct {
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	AddressID    int32  `json:"addressID"`
	CreditCardID int32  `json:"creditCardID"`
}

// CreateCheckout
//
//	INSERT INTO checkout.checkout(owner, name, firstname, lastname, email, address_id, credit_card_id)
//	VALUES ($1, $2, $3, $4, $5, $6, $7)
//	RETURNING id, owner, name, firstname, lastname, email, address_id, credit_card_id
func (q *Queries) CreateCheckout(ctx context.Context, arg CreateCheckoutParams) (CheckoutCheckout, error) {
	row := q.db.QueryRow(ctx, CreateCheckout,
		arg.Owner,
		arg.Name,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.AddressID,
		arg.CreditCardID,
	)
	var i CheckoutCheckout
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.AddressID,
		&i.CreditCardID,
	)
	return i, err
}
