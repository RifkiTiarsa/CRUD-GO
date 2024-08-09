CREATE TABLE tx_bill_detail (
    id SERIAL PRIMARY KEY,
    bill_id BIGINT UNSIGNED NOT NULL,
    product_id BIGINT UNSIGNED NOT NULL,
    product_price INT NOT NULL,
    qty INT NOT NULL,
    FOREIGN KEY (bill_id) REFERENCES tx_transaction(id),
    FOREIGN KEY (product_id) REFERENCES mst_product(id)
);