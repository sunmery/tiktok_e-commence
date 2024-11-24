CREATE SCHEMA checkout;

CREATE TABLE checkout.checkout
(
    id             SERIAL PRIMARY KEY,
    owner          VARCHAR(100) NOT NULL,
    name           VARCHAR(100) NOT NULL,
    firstname      VARCHAR(50)  NOT NULL,
    lastname       VARCHAR(50)  NOT NULL,
    email          VARCHAR(255) NOT NULL,
    address_id     INT          NOT NULL, -- 引用地址表的字段
    credit_card_id INT          NOT NULL  -- 引用银行卡表的字段
);
