-- name: CreateOrUpdateCartItem :one
WITH cart AS (
    INSERT INTO carts.carts (name)
        VALUES (@name)
        ON CONFLICT (name) DO NOTHING
        RETURNING id),
     existing_cart AS (SELECT id
                       FROM carts.carts
                       WHERE owner = @owner
                         AND name = @name)
INSERT
INTO carts.cart_items (owner, name, cart_id, product_id, quantity)
VALUES (@owner,
        @name,
        COALESCE((SELECT id FROM cart), (SELECT id FROM existing_cart)),
        @product_id,
        @quantity)
ON CONFLICT (cart_id, product_id) DO UPDATE
    SET quantity = carts.cart_items.quantity + EXCLUDED.quantity
RETURNING *;

-- name: DeleteCart :one
DELETE
FROM carts.carts
WHERE owner = @owner
  AND name = @name
RETURNING *;

-- name: DeleteCartItem :one
DELETE
FROM carts.cart_items
WHERE owner = @owner
  AND name = @name
  AND cart_id = @cart_id
  AND product_id = @product_id
RETURNING *;

-- name: GetCart :many
SELECT c.owner,
       c.name,
       p.name,
       p.description,
       p.price,
       p.picture,
       p.categories,
       ci.product_id,
       ci.quantity
FROM carts.carts c
         INNER JOIN
     carts.cart_items ci ON c.owner = ci.owner AND c.name = ci.name
         INNER JOIN
     products.products p ON ci.product_id = p.id
WHERE c.owner = @owner
  AND c.name = @name
ORDER BY ci.created_at;

