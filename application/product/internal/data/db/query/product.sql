-- name: CreateProduct :one
INSERT INTO products(name,description,picture,price,categories)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListProducts :many
SELECT *
FROM products
WHERE sqlc.arg(category_name) = ANY (categories)
ORDER BY id
OFFSET sqlc.arg(page)
LIMIT sqlc.arg(page_size);
