SET search_path TO orders;

DROP TABLE IF EXISTS orders.addresses;
DROP TABLE IF EXISTS orders.orders;

DROP INDEX IF EXISTS orders.idx_order_email;

DROP SCHEMA IF EXISTS orders CASCADE;
