-- name: CreateCartItem :exec
WITH cart AS (
    INSERT INTO carts (user_id)
        VALUES (sqlc.arg(user_id))
        ON CONFLICT (user_id) DO UPDATE SET user_id = carts.user_id
        RETURNING id),
     existing_cart AS (SELECT id
                       FROM carts
                       WHERE user_id = sqlc.arg(user_id))
INSERT
INTO cart_items (user_id, cart_id, product_id, quantity)
VALUES (sqlc.arg(user_id),
        (SELECT id FROM cart UNION ALL SELECT id FROM existing_cart LIMIT 1),
        sqlc.arg(product_id),
        sqlc.arg(quantity))
ON CONFLICT (cart_id, product_id) DO UPDATE
    SET quantity = cart_items.quantity + EXCLUDED.quantity;

-- name: DeleteCart :one
DELETE
FROM carts
WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: DeleteCartItem :one
DELETE
FROM cart_items
WHERE user_id = sqlc.arg(user_id)
  AND cart_id = sqlc.arg(cart_id)
  AND product_id = sqlc.arg(product_id)
RETURNING *;

-- name: GetCart :many
SELECT c.user_id,
       p.name,
       p.description,
       p.price,
       p.picture,
       p.categories,
       ci.product_id,
       ci.quantity
FROM carts. c
         INNER JOIN
     cart_items ci ON c.user_id = ci.user_id
         INNER JOIN
     products p ON ci.id = p.id
WHERE c.user_id = sqlc.arg(user_id)
ORDER BY ci.created_at;

