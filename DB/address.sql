-- 建立 address 表（如果不存在）
CREATE TABLE IF NOT EXISTS address (
    -- 地址的唯一標識符
    id UUID PRIMARY KEY,
    -- 使用者的唯一標識符
    user_id UUID REFERENCES "user"(id) ON DELETE CASCADE,
    -- 地址的詳細描述
    description TEXT,
    -- 郵遞區號
    postal_code VARCHAR(10),
    -- 地址的國家/地區
    country VARCHAR(100),
    -- 地址的城市
    city VARCHAR(100),
    -- 地址的區域
    state VARCHAR(100),
    -- 地址的街道
    street_address TEXT
);

-- 為表添加註釋
COMMENT ON TABLE address IS '用於存儲使用者地址的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN address.id IS '地址的唯一標識符';
COMMENT ON COLUMN address.user_id IS '使用者的唯一標識符';
COMMENT ON COLUMN address.description IS '地址的詳細描述';
COMMENT ON COLUMN address.postal_code IS '郵遞區號';
COMMENT ON COLUMN address.country IS '地址的國家/地區';
COMMENT ON COLUMN address.city IS '地址的城市';
COMMENT ON COLUMN address.state IS '地址的區域';
COMMENT ON COLUMN address.street_address IS '地址的街道';
