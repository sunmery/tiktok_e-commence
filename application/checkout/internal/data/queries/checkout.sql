-- name: CreateCheckout :one
INSERT INTO checkout.checkout(owner, name, firstname, lastname, email, address_id, credit_card_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
