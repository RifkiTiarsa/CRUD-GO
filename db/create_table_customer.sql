CREATE TABLE mst_customer (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        phone_number VARCHAR(15) NOT NULL,
        address VARCHAR(100) NOT NULL,
		UNIQUE (name),
		UNIQUE (phone_number)
);