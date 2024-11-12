-- name: CreateProduct :one
INSERT INTO products(name,description,picture,price,categories)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
