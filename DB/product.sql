-- 建立 product 表（如果不存在）
CREATE TABLE IF NOT EXISTS product (
    -- 商品的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 商品名稱
    name VARCHAR(100) NOT NULL,
    -- 商品價格
    price DECIMAL(10, 2) NOT NULL,
    -- 商品庫存數量
    stock_quantity INTEGER NOT NULL,
    -- 商品分類的唯一標識符
    category_id UUID NOT NULL,
    -- 商品描述
    description TEXT,
    -- 圖片/影片連結 1
    media_link_1 VARCHAR(255),
    -- 圖片/影片連結 2
    media_link_2 VARCHAR(255),
    -- 圖片/影片連結 3
    media_link_3 VARCHAR(255),
    -- 圖片/影片連結 4
    media_link_4 VARCHAR(255),
    -- 圖片/影片連結 5
    media_link_5 VARCHAR(255),
    -- 圖片/影片連結 6
    media_link_6 VARCHAR(255)
);

-- 為表添加註釋
COMMENT ON TABLE product IS '用於存儲商品信息的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN product.id IS '商品的唯一標識符';
COMMENT ON COLUMN product.name IS '商品名稱';
COMMENT ON COLUMN product.price IS '商品價格';
COMMENT ON COLUMN product.stock_quantity IS '商品庫存數量';
COMMENT ON COLUMN product.category_id IS '商品分類的唯一標識符';
COMMENT ON COLUMN product.description IS '商品描述';
COMMENT ON COLUMN product.media_link_1 IS '圖片/影片連結 1';
COMMENT ON COLUMN product.media_link_2 IS '圖片/影片連結 2';
COMMENT ON COLUMN product.media_link_3 IS '圖片/影片連結 3';
COMMENT ON COLUMN product.media_link_4 IS '圖片/影片連結 4';
COMMENT ON COLUMN product.media_link_5 IS '圖片/影片連結 5';
COMMENT ON COLUMN product.media_link_6 IS '圖片/影片連結 6';