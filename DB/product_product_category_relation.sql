-- 建立 product_product_category_relation 表（如果不存在）
CREATE TABLE IF NOT EXISTS product_product_category_relation (
    -- 關聯的唯一標識符
    id UUID PRIMARY KEY,
    -- 商品的唯一標識符
    product_id UUID NOT NULL,
    -- 商品分類的唯一標識符
    category_id UUID NOT NULL,
    -- 建立時間
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 為表添加註釋
COMMENT ON TABLE product_product_category_relation IS '用於存儲商品和商品分類之間關聯的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN product_product_category_relation.id IS '關聯的唯一標識符';
COMMENT ON COLUMN product_product_category_relation.product_id IS '商品的唯一標識符';
COMMENT ON COLUMN product_product_category_relation.category_id IS '商品分類的唯一標識符';
COMMENT ON COLUMN product_product_category_relation.created_at IS '建立時間';
