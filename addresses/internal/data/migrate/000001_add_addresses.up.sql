CREATE SCHEMA IF NOT EXISTS addresses;

SET search_path TO addresses;

CREATE TABLE addresses
(
    id             SERIAL PRIMARY KEY,
    owner          VARCHAR(100) NOT NULL, -- 用户组织
    name           VARCHAR(100) NOT NULL, -- 用户地址
    street_address TEXT         NOT NULL, -- 街道地址
    city           VARCHAR(255) NOT NULL, -- 城市
    state          VARCHAR(20)  NOT NULL, -- 状态
    country        VARCHAR(100) NOT NULL, -- 国家
    zip_code       VARCHAR(20)  NOT NULL  -- 邮政编码
);

CREATE INDEX idx_addresses_name ON addresses (name);
