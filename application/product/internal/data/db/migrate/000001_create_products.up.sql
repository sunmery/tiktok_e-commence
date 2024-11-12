CREATE TABLE IF NOT EXISTS products(
    id  SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL ,
    picture TEXT NOT NULL,
    price REAL NOT NULL,
    categories TEXT[]
);

-- 加快通过名称查询的商品
CREATE INDEX idx_products_name ON products (name);

-- GIN 索引会显著加速数组查询
CREATE INDEX idx_products_categories ON products USING GIN (categories);

-- 查询某类别的商品
-- SELECT * FROM products WHERE categories @> ARRAY['electronics'];
