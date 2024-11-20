-- name: CreateProduct :one
INSERT INTO products.products(name, description, picture, price, categories)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListProducts :many
SELECT *
FROM products.products
ORDER BY id
OFFSET sqlc.arg(page) LIMIT sqlc.arg(page_size);

-- name: GetProduct :one
SELECT *
FROM products.products
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: SearchProducts :many
SELECT *
FROM products.products
WHERE name = sqlc.arg(name);
