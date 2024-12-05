CREATE SCHEMA IF NOT EXISTS orders;

CREATE TABLE orders.orders
(
    id         SERIAL PRIMARY KEY,
    owner      VARCHAR(100)                  NOT NULL,
    name       VARCHAR(100)                  NOT NULL,
    email      VARCHAR(100)                  NOT NULL,
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
