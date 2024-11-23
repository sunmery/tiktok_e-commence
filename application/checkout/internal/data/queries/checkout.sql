-- name: CreateCheckout :one
INSERT INTO checkout.checkout(user_id, firstname, lastname, email, address_id, credit_card_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
