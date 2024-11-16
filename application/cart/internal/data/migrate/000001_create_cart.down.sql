-- 设置 search_path 到 carts schema
SET search_path TO carts;

-- 移除 cart_items 表上的外键和约束
ALTER TABLE IF EXISTS cart_items DROP CONSTRAINT IF EXISTS unique_cart_product;
ALTER TABLE IF EXISTS cart_items DROP CONSTRAINT IF EXISTS cart_items_cart_id_fkey;
ALTER TABLE IF EXISTS cart_items DROP CONSTRAINT IF EXISTS cart_items_product_id_fkey;
ALTER TABLE IF EXISTS cart_items DROP CONSTRAINT IF EXISTS cart_items_user_id_fkey;

-- 移除 carts 表上的外键
ALTER TABLE IF EXISTS carts DROP CONSTRAINT IF EXISTS carts_user_id_fkey;

-- 删除索引
DROP INDEX IF EXISTS idx_carts_id;
DROP INDEX IF EXISTS idx_user_id;

-- 删除表
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;

-- 删除用户表上的唯一约束
SET search_path TO public;
ALTER TABLE public.user DROP CONSTRAINT IF EXISTS IDX_user_id;

-- 最后删除 carts schema
SET search_path TO public;
DROP SCHEMA IF EXISTS carts CASCADE;
