-- 建立最愛表（如果不存在）
CREATE TABLE IF NOT EXISTS favorite (
    -- 最愛的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 關聯的使用者ID
    user_id UUID REFERENCES user(id) ON DELETE CASCADE,
    -- 關聯的商品ID
    product_id UUID REFERENCES product(id) ON DELETE CASCADE,
    -- 最愛的建立時間
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 為表添加註釋
COMMENT ON TABLE favorite IS '用於存儲使用者最愛商品資訊的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN favorite.id IS '最愛的唯一標識符';
COMMENT ON COLUMN favorite.user_id IS '關聯的使用者ID';
COMMENT ON COLUMN favorite.product_id IS '關聯的商品ID';
COMMENT ON COLUMN favorite.created_at IS '最愛的建立時間';
