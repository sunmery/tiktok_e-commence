CREATE SCHEMA IF NOT EXISTS carts;

CREATE TABLE carts.cart_items
(
    id         SERIAL PRIMARY KEY,
    owner      VARCHAR(100)                NOT NULL,
    name       VARCHAR(100)                NOT NULL,
    product_id INT                         NOT NULL,
    quantity   INT                         NOT NULL,
    created_at timestamptz DEFAULT (now()) NOT NULL,
    updated_at timestamptz DEFAULT (now()) NOT NULL
);
