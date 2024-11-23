CREATE SCHEMA IF NOT EXISTS products;

CREATE TABLE IF NOT EXISTS products.products
(
    id          SERIAL PRIMARY KEY,
    owner       VARCHAR(100) NOT NULL DEFAULT 'tiktok',
    name        VARCHAR(100) NOT NULL,
    description TEXT         NOT NULL,
    picture     TEXT         NOT NULL,
    price       REAL         NOT NULL,
    categories  TEXT[]       NOT NULL
);
