CREATE TABLE tx_transaction (
    id SERIAL PRIMARY KEY,
    bill_date VARCHAR(100) NOT NULL,
    entry_date VARCHAR(100) NOT NULL,
    finish_date VARCHAR(100) NOT NULL,
    employee_id BIGINT UNSIGNED NOT NULL,
    customer_id BIGINT UNSIGNED NOT NULL,
    total INT,
    FOREIGN KEY (employee_id) REFERENCES mst_employee(id),
    FOREIGN KEY (customer_id) REFERENCES mst_customer(id)
);