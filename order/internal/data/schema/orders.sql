CREATE SCHEMA IF NOT EXISTS orders;

CREATE TABLE orders.orders
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255)              NOT NULL,
    user_id       VARCHAR(100)              NOT NULL,
    address_id    INT                       NOT NULL,
    user_currency char(3)                   NOT NULL, -- 货币类型, ISO 4217 标准
    paid          bool        DEFAULT false NOT NULL, -- 是否已支付
    created_at    timestamptz DEFAULT now() NOT NULL,
    updated_at    timestamptz DEFAULT now() NOT NULL
);
