-- 建立 user_permission_relation 表（如果不存在）
CREATE TABLE IF NOT EXISTS user_permission_relation (
    -- 關聯的唯一標識符
    id UUID PRIMARY KEY,
    -- 使用者的唯一標識符
    user_id UUID NOT NULL,
    -- 權限的唯一標識符
    permission_id UUID NOT NULL,
    -- 建立時間
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 為表添加註釋
COMMENT ON TABLE user_permission_relation IS '用於存儲使用者和權限之間關聯的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN user_permission_relation.id IS '關聯的唯一標識符';
COMMENT ON COLUMN user_permission_relation.user_id IS '使用者的唯一標識符';
COMMENT ON COLUMN user_permission_relation.permission_id IS '權限的唯一標識符';
COMMENT ON COLUMN user_permission_relation.created_at IS '建立時間';
