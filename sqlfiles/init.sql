

CREATE TABLE promotion (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR       NOT NULL,
    description VARCHAR       NOT NULL,
    status      SMALLINT      ,
    created_at  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

CREATE TABLE product (
    sku         VARCHAR       PRIMARY KEY,
    name        VARCHAR       NOT NULL,
    price       NUMERIC       NOT NULL,
    currency    VARCHAR       NOT NULL,
    quantity    BIGINT        ,
    status      SMALLINT      ,
    created_at  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);

CREATE TABLE product_promo_requirement (
    id          SERIAL        PRIMARY KEY,
    promo_id    BIGINT        NOT NULL,
    product_sku VARCHAR       NOT NULL,
    minimum_qty BIGINT        NOT NULL,
    created_at  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP,
    FOREIGN KEY (promo_id) REFERENCES promotion (id),
    FOREIGN KEY (product_sku) REFERENCES product (sku)
);

CREATE TABLE promo_outcome (
    id             SERIAL        PRIMARY KEY,
    promo_id       BIGINT        NOT NULL,
    product_sku    VARCHAR       NOT NULL,
    promotion_type VARCHAR       NOT NULL,
    amount         NUMERIC       NOT NULL,
    qty            BIGINT        NOT NULL,
    created_at     TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP,
    FOREIGN KEY (promo_id) REFERENCES promotion (id),
    FOREIGN KEY (product_sku) REFERENCES product (sku)
);

INSERT INTO product (sku, name, price, currency, quantity) VALUES ('120P90', 'Google Home', 49.99, 'USD', 10);
INSERT INTO product (sku, name, price, currency, quantity) VALUES ('43N23P', 'Macbook Pro', 5399.99, 'USD', 5);
INSERT INTO product (sku, name, price, currency, quantity) VALUES ('A304SD', 'Alexa Speaker', 109.50, 'USD', 10);
INSERT INTO product (sku, name, price, currency, quantity) VALUES ('234234', 'Raspberry Pi B', 30.00, 'USD', 2);

INSERT INTO promotion (id, name, description) VALUES (1, 'buy_macbook_free_raspberry', 'Each sale of a MacBook Pro comes with a free Raspberry Pi B');
INSERT INTO product_promo_requirement (promo_id, product_sku, minimum_qty) VALUES (1, '43N23P', 1);
INSERT INTO product_promo_requirement (promo_id, product_sku, minimum_qty) VALUES (1, '234234', 1);
INSERT INTO promo_outcome (promo_id, product_sku, promotion_type, amount, qty) VALUES (1, '234234', 'percentage', 100, 1);

INSERT INTO promotion (id, name, description) VALUES (2, 'buy3_google_home_pay2', 'Buy 3 Google Homes for the price of 2');
INSERT INTO product_promo_requirement (promo_id, product_sku, minimum_qty) VALUES (2, '120P90', 3);
INSERT INTO promo_outcome (promo_id, product_sku, promotion_type, amount, qty) VALUES (2, '120P90', 'percentage', 100, 1);

INSERT INTO promotion (id, name, description) VALUES (3, 'buy3_alexa_disc_10', 'Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers');
INSERT INTO product_promo_requirement (promo_id, product_sku, minimum_qty) VALUES (3, 'A304SD', 3);
INSERT INTO promo_outcome (promo_id, product_sku, promotion_type, amount, qty) VALUES (3, 'A304SD', 'percentage', 10, 999999);