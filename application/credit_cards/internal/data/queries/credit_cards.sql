-- name: CreateCreditCard :one
INSERT INTO credit_cards.credit_cards(owner,
                                      username,
                                      credit_card_number,
                                      credit_card_cvv,
                                      credit_card_expiration_year,
                                      credit_card_expiration_month)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateCreditCard :one
UPDATE credit_cards.credit_cards
SET credit_card_number           = coalesce(sqlc.narg(credit_card_number), credit_card_number),
    credit_card_cvv              = coalesce(sqlc.narg(credit_card_cvv), credit_card_cvv),
    credit_card_expiration_year  = coalesce(sqlc.narg(credit_card_expiration_year), credit_card_expiration_year),
    credit_card_expiration_month = coalesce(sqlc.narg(credit_card_expiration_month), credit_card_expiration_month)
WHERE username = @username
RETURNING *;

-- name: DeleteCreditCard :one
DELETE
FROM credit_cards.credit_cards
WHERE username = @username
  AND id = @id
RETURNING *;

-- name: GetCreditCards :many
SELECT *
FROM credit_cards.credit_cards
WHERE credit_card_number ILIKE '%' || $1 || '%';

-- name: ListCreditCards :many
SELECT *
FROM credit_cards.credit_cards
WHERE username = @username;
