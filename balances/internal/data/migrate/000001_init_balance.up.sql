create schema balances;
create table balances.balances
(
    owner        VARCHAR(100)                                       NOT NULL,
    name         VARCHAR(100)                                       NOT NULL,
    balance      DOUBLE PRECISION DEFAULT 0.0 check ( balance > 0 ) NOT NULL,
    currency     VARCHAR          DEFAULT 'CNY'                     NOT NULL,
    created_time TIMESTAMPTZ      DEFAULT now(),
    updated_time TIMESTAMPTZ      DEFAULT now()
);
