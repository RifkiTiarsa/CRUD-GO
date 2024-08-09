CREATE TABLE mst_employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    address VARCHAR(100) NOT NULL,
    UNIQUE (name),
    UNIQUE (phone_number)
);