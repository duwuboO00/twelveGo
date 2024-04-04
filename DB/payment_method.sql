-- 建立 payment_method 表（如果不存在）
CREATE TABLE IF NOT EXISTS payment_method (
    -- 支付方式的唯一標識符
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    -- 支付方式名稱
    name TEXT NOT NULL,
    -- 支付方式描述
    description TEXT
);

-- 為表添加註釋
COMMENT ON TABLE payment_method IS '用於存儲支付方式信息的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN payment_method.id IS '支付方式的唯一標識符';
COMMENT ON COLUMN payment_method.name IS '支付方式名稱';
COMMENT ON COLUMN payment_method.description IS '支付方式描述';
