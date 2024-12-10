-- name: CreateCreditCard :one
INSERT INTO credit_cards.credit_cards(owner,
                                      name,
                                      number,
                                      cvv,
                                      expiration_year,
                                      expiration_month)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateCreditCard :one
UPDATE credit_cards.credit_cards
SET number           = coalesce(sqlc.narg(number), number),
    cvv              = coalesce(sqlc.narg(cvv), cvv),
    expiration_year  = coalesce(sqlc.narg(expiration_year), expiration_year),
    expiration_month = coalesce(sqlc.narg(expiration_month), expiration_month)
WHERE owner = @owner
  AND name = @name
RETURNING *;

-- name: DeleteCreditCard :one
DELETE
FROM credit_cards.credit_cards
WHERE owner = @owner
  AND name = @name
  AND id = @id
RETURNING *;

-- name: GetCreditCards :one
SELECT *
FROM credit_cards.credit_cards
WHERE owner = @owner
  AND name = @name
  AND number = @number;

-- name: SearchCreditCards :many
SELECT *
FROM credit_cards.credit_cards
WHERE owner = @owner
  AND name = @name
  AND number ILIKE '%' || @number;

-- name: ListCreditCards :many
SELECT *
FROM credit_cards.credit_cards
WHERE owner = @owner
  AND name = @name;
