-- name: CreatePayment :one
INSERT INTO payment.payments(snowflake_id, owner, name, amount, order_id, credit_card_number, credit_card_cvv,
                             credit_card_expiration_year, credit_card_expiration_month, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

