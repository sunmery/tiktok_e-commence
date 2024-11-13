CREATE TABLE carts
(
    id         SERIAL PRIMARY KEY,
    user_id    INT                         NOT NULL,
    created_at timestamptz DEFAULT (now()) NOT NULL,
    updated_at timestamptz DEFAULT (now()) NOT NULL
);

CREATE TABLE cart_items
(
    id         SERIAL PRIMARY KEY,
    cart_id    INT                         NOT NULL, -- 添加 cart_id 外键关联到购物车
    product_id INT                         NOT NULL,
    quantity   INT                         NOT NULL,
    created_at timestamptz DEFAULT (now()) NOT NULL,
    updated_at timestamptz DEFAULT (now()) NOT NULL
);

-- 优化通过用户查询购物车的性能
CREATE INDEX idx_user_id ON carts (id);

-- 关联用户表的id
ALTER TABLE carts
    ADD
        FOREIGN KEY (user_id) REFERENCES users (id);

-- 级联删除
ALTER TABLE cart_items
    ADD
        FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE CASCADE;

-- 商品删除时删除购物车项
ALTER TABLE cart_items
    ADD
        FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE;


