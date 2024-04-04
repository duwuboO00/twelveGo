-- 建立 user 表（如果不存在）
CREATE TABLE IF NOT EXISTS "user" (
    -- 用戶的唯一標識符
    id UUID PRIMARY KEY,
    -- 用戶的名稱
    user_name VARCHAR(100) NOT NULL,
    -- 用戶的電子郵件地址
    email VARCHAR(255) NOT NULL UNIQUE,
    -- Line ID
    line_id VARCHAR(30) UNIQUE,
    -- Google ID
    google_id VARCHAR(30) UNIQUE,
    -- Apple ID
    apple_id VARCHAR(30) UNIQUE,
    -- 用戶的出生日期
    date_of_birth DATE NOT NULL,
    -- 用戶的性別（M 表示男性，F 表示女性，U 未知）
    gender CHAR(1) CHECK (gender IN ('M', 'F', 'U')),
    -- 用戶的地址
    address TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- 為表添加註釋
COMMENT ON TABLE "user" IS '用於存儲用戶信息的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN "user".id IS '用戶的唯一標識符';
COMMENT ON COLUMN "user".user_name IS '用戶的名稱';
COMMENT ON COLUMN "user".email IS '用戶的電子郵件地址';
COMMENT ON COLUMN "user".line_id IS 'Line ID';
COMMENT ON COLUMN "user".google_id IS 'Google ID';
COMMENT ON COLUMN "user".apple_id IS 'Apple ID';
COMMENT ON COLUMN "user".date_of_birth IS '用戶的出生日期';
COMMENT ON COLUMN "user".gender IS '用戶的性別（M 表示男性，F 表示女性，U 未知）';
COMMENT ON COLUMN "user".address IS '用戶的地址';
