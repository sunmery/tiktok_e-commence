-- 记录支付交易信息

CREATE SCHEMA payment;

CREATE TABLE payment.payments
(
    id                           SERIAL PRIMARY KEY,
    snowflake_id                 BIGINT UNIQUE             NOT NULL,
    owner                        VARCHAR(100)              NOT NULL,
    name                         VARCHAR(100)              NOT NULL,
    amount                       DECIMAL(10, 2)            NOT NULL,
    order_id                     INT                       NOT NULL,
    credit_card_number           VARCHAR(20)               NOT NULL, -- 卡号
    credit_card_cvv              VARCHAR(4)                NOT NULL, -- 安全码
    credit_card_expiration_year  CHAR(4)                   NOT NULL, -- 过期年份
    credit_card_expiration_month CHAR(2)                   NOT NULL, -- 过期月份
    status                       VARCHAR(20)               NOT NULL, -- 支付状态: PENDING, SUCCESS, FAILED
    created_at                   timestamptz DEFAULT now() NOT NULL
);