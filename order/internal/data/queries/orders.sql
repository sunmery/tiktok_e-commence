-- name: CreateOrder :one
INSERT INTO orders.orders (email, user_id, address_id, user_currency)
VALUES (@email, @user_id, @address_id, @user_currency)
RETURNING *;

-- name: ListOrders :many
SELECT *
FROM orders.orders
WHERE user_id = @user_id
ORDER BY created_at;

-- name: MarkOrderPaid :many
SELECT *
FROM orders.orders
WHERE user_id = @user_id
  AND paid = true
ORDER BY created_at;
