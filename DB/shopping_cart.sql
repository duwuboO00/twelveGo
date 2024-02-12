-- 建立購物車表（如果不存在）
CREATE TABLE IF NOT EXISTS shopping_cart (
    -- 購物車的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 關聯的使用者ID
    user_id UUID REFERENCES user(id) ON DELETE CASCADE,
    -- 關聯的商品ID
    product_id UUID REFERENCES product(id) ON DELETE CASCADE,
    -- 購物車的建立時間
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    -- 購物車的更新時間
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);