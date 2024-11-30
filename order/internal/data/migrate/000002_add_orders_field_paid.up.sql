ALTER TABLE orders.orders
    ADD
        COLUMN paid bool DEFAULT false NOT NULL;
