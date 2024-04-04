-- 建立 order_item 表（如果不存在）
CREATE TABLE IF NOT EXISTS order_item (
    -- 訂單明細的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 所屬訂單的唯一標識符
    order_id UUID NOT NULL,
    -- 商品的唯一標識符
    product_id UUID NOT NULL,
    -- 商品連結
    product_link VARCHAR(255) NOT NULL,
    -- 商品名稱
    product_name VARCHAR(100) NOT NULL,
    -- 商品副標/簡易說明
    product_subtitle VARCHAR(255),
    -- 商品單價
    unit_price DECIMAL(10, 2) NOT NULL,
    -- 購買數量
    quantity INTEGER NOT NULL
);

-- 為表添加註釋
COMMENT ON TABLE order_item IS '用於存儲訂單明細的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN order_item.id IS '訂單明細的唯一標識符';
COMMENT ON COLUMN order_item.order_id IS '所屬訂單的唯一標識符';
COMMENT ON COLUMN order_item.product_id IS '商品的唯一標識符';
COMMENT ON COLUMN order_item.product_link IS '商品連結';
COMMENT ON COLUMN order_item.product_name IS '商品名稱';
COMMENT ON COLUMN order_item.product_subtitle IS '商品副標/簡易說明';
COMMENT ON COLUMN order_item.unit_price IS '商品單價';
COMMENT ON COLUMN order_item.quantity IS '購買數量';
