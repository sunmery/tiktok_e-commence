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

-- 关联用户表的id
ALTER TABLE carts.cart_items
    ADD
        FOREIGN KEY (owner, name) REFERENCES public.user (owner, name);

-- 每个购物车中只能有一个特定的商品（即一个 cart_id 对应多个商品，但每个商品只能出现一次
ALTER TABLE carts.cart_items
    ADD
        CONSTRAINT unique_cart_product UNIQUE (owner, name, product_id);

-- 商品删除时删除购物车项
ALTER TABLE carts.cart_items
    ADD
        FOREIGN KEY (product_id) REFERENCES products.products (id) ON DELETE CASCADE;
