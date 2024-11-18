-- name: CreatUserAddress :one
INSERT INTO addresses.addresses(user_id, street_address, city, state, country, zip_code)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserAddress :one
SELECT *
FROM addresses.addresses
WHERE user_id = @user_id
LIMIT 1;

-- name: ListUserAddresses :many
SELECT *
FROM addresses.addresses
WHERE user_id = @user_id;

-- name: UpdateUserAddress :one
UPDATE addresses.addresses
SET street_address = coalesce(sqlc.narg(street_address), street_address),
    city           = coalesce(sqlc.narg(city), city),
    state          = coalesce(sqlc.narg(state), state),
    country        = coalesce(sqlc.narg(country), country),
    zip_code       = coalesce(sqlc.narg(zip_code), zip_code)
WHERE user_id = @user_id
RETURNING *;

-- name: DeleteUserAddress :one
DELETE
FROM addresses.addresses
WHERE user_id = $1
  AND id = $2
RETURNING *;

-- name: DeleteUserAddresses :one
DELETE
FROM addresses.addresses
WHERE user_id = @user_id
RETURNING *;
