CREATE SCHEMA IF NOT EXISTS orders;

SET search_path TO orders;

CREATE TABLE orders
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255)              NOT NULL,
    user_id       VARCHAR(100)              NOT NULL,
    address_id    INT                       NOT NULL,
    user_currency char(3)                   NOT NULL, -- 货币类型, ISO 4217 标准
    created_at    timestamptz DEFAULT now() NOT NULL,
    updated_at    timestamptz DEFAULT now() NOT NULL
);

-- 关联订单到用户
ALTER TABLE orders
    ADD
        FOREIGN KEY (user_id) REFERENCES public.user (id);

-- 关联订单到地址
ALTER TABLE orders
    ADD
        FOREIGN KEY (address_id) REFERENCES addresses.addresses (id);
