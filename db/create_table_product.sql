CREATE TABLE mst_product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    unit VARCHAR(100) NOT NULL,
    UNIQUE (name)
);