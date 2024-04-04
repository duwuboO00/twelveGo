
-- 建立 product_category 表（如果不存在）
CREATE TABLE IF NOT EXISTS product_category (
    -- 商品分類的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 商品分類名稱
    name TEXT NOT NULL,
    -- 商品分類的描述
    description TEXT
);


-- 為表添加註釋
COMMENT ON TABLE product_category IS '用於存儲商品分類信息的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN product_category.id IS '商品分類的唯一標識符';
COMMENT ON COLUMN product_category.name IS '商品分類名稱';
COMMENT ON COLUMN product_category.description IS '商品分類的描述';
