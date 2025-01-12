-- name: CreatBalance :one
INSERT INTO balances.balances(owner, name, balance)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBalance :one
SELECT *
FROM balances.balances
WHERE owner = @owner
  AND name = @name;

-- name: UpdateBalance :one
UPDATE balances.balances
SET balance = coalesce(sqlc.narg(balance), balance)
WHERE owner = @owner
  AND name = @name
RETURNING *;

-- name: DeleteBalance :one
DELETE
FROM balances.balances
WHERE id = @id
  AND owner = @owner
  AND name = @name
RETURNING *;
