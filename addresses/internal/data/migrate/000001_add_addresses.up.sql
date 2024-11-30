CREATE SCHEMA IF NOT EXISTS addresses;

SET search_path TO addresses;

CREATE TABLE addresses
(
    id             SERIAL PRIMARY KEY,
    user_id        VARCHAR(100) NOT NULL, -- 用户地址
    street_address TEXT         NOT NULL, -- 街道地址
    city           VARCHAR(255) NOT NULL, -- 城市
    state          VARCHAR(20)  NOT NULL, -- 状态
    country        VARCHAR(100) NOT NULL, -- 国家
    zip_code       int4         NOT NULL  -- 邮政编码
);

CREATE INDEX idx_addresses_user_id ON addresses (user_id);
