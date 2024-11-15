SET search_path TO carts;

-- 先删除外键约束
ALTER TABLE cart_items DROP CONSTRAINT IF EXISTS cart_items_cart_id_fkey;
ALTER TABLE cart_items DROP CONSTRAINT IF EXISTS cart_items_product_id_fkey;
ALTER TABLE carts DROP CONSTRAINT IF EXISTS carts_user_id_fkey;

-- 删除表
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS cart_items;

-- 删除索引
DROP INDEX IF EXISTS idx_user_id;
-- 删除schema
DROP SCHEMA IF EXISTS carts;
