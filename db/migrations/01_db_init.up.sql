CREATE SCHEMA IF NOT EXISTS cyclo;

--  brands
CREATE TABLE IF NOT EXISTS cyclo.brands (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- products
CREATE TABLE IF NOT EXISTS cyclo.products (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    price FLOAT NOT NULL,
    brand_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CONSTRAINT product_brand_fk FOREIGN KEY (brand_id) REFERENCES cyclo.brands(id)
);



-- -- product brand
-- CREATE TABLE IF NOT EXISTS cyclo.product_brands (
--     id serial PRIMARY KEY,
--     product_id int NOT NULL,
--     brand_id int NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMP,
--     
--     CONSTRAINT product_fk FOREIGN KEY (product_id) REFERENCES cyclo.products(id),
--     CONSTRAINT brand_fk FOREIGN KEY (brand_id) REFERENCES cyclo.brands(id)
-- );
