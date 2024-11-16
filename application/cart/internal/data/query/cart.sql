-- name: CreateOrUpdateCartItem :one
WITH cart AS (
    INSERT INTO carts (user_id)
        VALUES (sqlc.arg(user_id))
        ON CONFLICT (user_id) DO NOTHING
        RETURNING id
),
     existing_cart AS (
         SELECT id FROM carts WHERE user_id = sqlc.arg(user_id)
     )
INSERT INTO cart_items (user_id, cart_id, product_id, quantity)
VALUES (
           sqlc.arg(user_id),
           COALESCE((SELECT id FROM cart), (SELECT id FROM existing_cart)),
           sqlc.arg(product_id),
           sqlc.arg(quantity)
       )
ON CONFLICT (cart_id, product_id) DO UPDATE
    SET quantity = cart_items.quantity + EXCLUDED.quantity
RETURNING *;

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
FROM carts c
         INNER JOIN
     cart_items ci ON c.user_id = ci.user_id
         INNER JOIN
     products.products p ON ci.product_id = p.id
WHERE c.user_id = sqlc.arg(user_id)
ORDER BY ci.created_at;

