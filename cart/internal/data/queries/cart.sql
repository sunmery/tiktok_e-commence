-- name: CreateOrUpdateCartItem :one
WITH existing_item AS (SELECT id
                       FROM carts.cart_items
                       WHERE owner = @owner
                         AND name = @name
                         AND product_id = 1)
INSERT
INTO carts.cart_items (owner, name, product_id, quantity)
VALUES (@owner, @name, @product_id, @quantity)
ON CONFLICT (owner, name, product_id) DO UPDATE
    SET quantity   = carts.cart_items.quantity + EXCLUDED.quantity,
        updated_at = now()
RETURNING *;

-- name: DeleteCart :one
DELETE
FROM carts.cart_items
WHERE owner = @owner
  AND name = @name
RETURNING *;

-- name: DeleteCartItem :one
DELETE
FROM carts.cart_items
WHERE owner = @owner
  AND name = @name
  AND product_id = @product_id
RETURNING *;

-- name: GetCart :many
SELECT p.description,
       p.price,
       p.picture,
       p.categories,
       ci.product_id,
       ci.quantity
FROM carts.cart_items ci
         INNER JOIN
     products.products p
     ON ci.product_id = p.id
WHERE ci.owner = @owner
  AND ci.name = @name
ORDER BY ci.created_at;

