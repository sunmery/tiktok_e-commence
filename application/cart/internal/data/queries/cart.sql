-- name: CreateOrUpdateCartItem :one
WITH cart AS (
    INSERT INTO carts.carts (user_id)
        VALUES (@user_id)
        ON CONFLICT (user_id) DO NOTHING
        RETURNING id),
     existing_cart AS (SELECT id
                       FROM carts.carts
                       WHERE user_id = @user_id)
INSERT
INTO carts.cart_items (user_id, cart_id, product_id, quantity)
VALUES (@user_id,
        COALESCE((SELECT id FROM cart), (SELECT id FROM existing_cart)),
        @product_id,
        @quantity)
ON CONFLICT (cart_id, product_id) DO UPDATE
    SET quantity = carts.cart_items.quantity + EXCLUDED.quantity
RETURNING *;

-- name: DeleteCart :one
DELETE
FROM carts.carts
WHERE user_id = @user_id
RETURNING *;

-- name: DeleteCartItem :one
DELETE
FROM carts.cart_items
WHERE user_id = @user_id
  AND cart_id = @cart_id
  AND product_id = @product_id
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
FROM carts.carts c
         INNER JOIN
     carts.cart_items ci ON c.user_id = ci.user_id
         INNER JOIN
     products.products p ON ci.product_id = p.id
WHERE c.user_id = @user_id
ORDER BY ci.created_at;

