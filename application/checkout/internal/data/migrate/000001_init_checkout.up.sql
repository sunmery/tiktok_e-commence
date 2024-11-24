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
    credit_card_id INT          NOT NULL  -- -- 引用银行卡表的字段
);

-- 关联用户
-- ALTER TABLE checkout.checkout
--     ADD
--         FOREIGN KEY (name) REFERENCES public.user (id);

-- 关联地址
ALTER TABLE checkout.checkout
    ADD
        FOREIGN KEY (address_id) REFERENCES addresses.addresses (id);

-- 关联银行卡
ALTER TABLE checkout.checkout
    ADD
        FOREIGN KEY (credit_card_id) REFERENCES credit_cards.credit_cards (id);
-- 索引
CREATE INDEX idx_checkout_id ON checkout.checkout (id);
CREATE INDEX idx_checkout_username ON checkout.checkout (name);
CREATE INDEX idx_checkout_address_id ON checkout.checkout (address_id);
CREATE INDEX idx_checkout_credit_card_id ON checkout.checkout (credit_card_id);
