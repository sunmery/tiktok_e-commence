-- name: CreateOrder :one
INSERT INTO orders.orders (email, user_id, address_id, user_currency)
VALUES (@email, @user_id, @address_id, @user_currency)
RETURNING *;
