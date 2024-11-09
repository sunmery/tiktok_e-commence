-- name: CreateUser :one
INSERT INTO users(email, password)
VALUES ($2, $3)
    WHERE id = $1
RETURNING *;
