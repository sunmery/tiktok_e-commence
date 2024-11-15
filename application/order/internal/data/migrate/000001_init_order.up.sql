CREATE SCHEMA IF NOT EXISTS orders;

SET search_path TO orders;

CREATE TABLE orders.orders
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255)              NOT NULL,
    user_id       INTEGER                   NOT NULL,
    user_currency char(3)                   NOT NULL, -- 货币类型, ISO 4217 标准
    created_at    timestamptz DEFAULT now() NOT NULL,
    updated_at    timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE orders.addresses
(
    street_address TEXT         NOT NULL, -- 街道地址
    city           VARCHAR(255) NOT NULL, -- 城市
    state          VARCHAR(20)  NOT NULL, -- 状态
    country        VARCHAR(100) NOT NULL, -- 国家
    zip_code       int4         NOT NULL -- 邮政编码
);

ALTER TABLE orders.orders
    ADD
        FOREIGN KEY (user_id) REFERENCES public.users (id);

CREATE INDEX idx_order_email ON orders.orders (email);
