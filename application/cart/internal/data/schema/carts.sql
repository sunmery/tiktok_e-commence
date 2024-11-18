CREATE SCHEMA IF NOT EXISTS carts;

CREATE TABLE carts.carts
(
    id         SERIAL PRIMARY KEY,
    user_id    VARCHAR(100) UNIQUE         NOT NULL,
    created_at timestamptz DEFAULT (now()) NOT NULL,
    updated_at timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE carts.cart_items
(
    id         SERIAL PRIMARY KEY,
    user_id    VARCHAR(100)                NOT NULL,
    cart_id    INT                         NOT NULL, -- 添加 cart_id 外键关联到购物车
    product_id INT                         NOT NULL,
    quantity   INT                         NOT NULL,
    created_at timestamptz DEFAULT (now()) NOT NULL,
    updated_at timestamptz DEFAULT (now()) NOT NULL
);
