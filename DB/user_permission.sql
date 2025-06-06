-- 建立 permission 表（如果不存在）
CREATE TABLE IF NOT EXISTS user_permission (
    -- 權限的唯一標識符
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 權限名稱
    name TEXT NOT NULL,
    -- 權限描述
    description TEXT
);

-- 為表添加註釋
COMMENT ON TABLE user_permission IS '用於存儲權限信息的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN user_permission.id IS '權限的唯一標識符';
COMMENT ON COLUMN user_permission.name IS '權限名稱';
COMMENT ON COLUMN user_permission.description IS '權限描述';
