CREATE SCHEMA IF NOT EXISTS orders;

SET search_path TO orders;

CREATE SCHEMA IF NOT EXISTS orders;

CREATE TABLE orders.orders
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255)                  NOT NULL,
    owner      VARCHAR(100)                  NOT NULL,
    name       VARCHAR(100)                  NOT NULL,
    address_id INT                           NOT NULL, -- 地址id, 关联地址微服务
    currency   CHAR(3)                       NOT NULL, -- 货币类型, ISO 4217 标准
    status     VARCHAR(10) DEFAULT 'pending' NOT NULL, -- 支付状态, pending已创建订单但未支付、paid已支付、cancelled已取消订单
    created_at timestamptz                   NOT NULL,
    updated_at timestamptz DEFAULT now()     NOT NULL
);

CREATE TABLE orders.order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT NOT NULL,
    product_id INT NOT NULL,
    quantity   INT NOT NULL,
    cost       INT NOT NULL
);

-- 通过order_items的order_id可以查询是哪个用户创建的订单
ALTER TABLE orders.order_items
    ADD
        CONSTRAINT fk_order_items_order_id
            FOREIGN KEY (order_id) REFERENCES orders.orders (id)
                ON DELETE CASCADE;

CREATE INDEX idx_order_user ON orders.orders (owner, name);
CREATE INDEX idx_order_items_order_id ON orders.order_items (order_id);
