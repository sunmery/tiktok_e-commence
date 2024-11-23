CREATE SCHEMA credit_cards;

CREATE TABLE credit_cards.credit_cards
(
    id                           SERIAL PRIMARY KEY,
    owner                        VARCHAR(100) NOT NULL DEFAULT 'tiktok',
    username                     VARCHAR(100) NOT NULL,
    credit_card_number           VARCHAR      NOT NULL,
    credit_card_cvv              INT          NOT NULL,
    credit_card_expiration_year  INT          NOT NULL,
    credit_card_expiration_month INT          NOT NULL
);
