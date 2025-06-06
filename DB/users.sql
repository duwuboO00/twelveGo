-- 建立 user 表（如果不存在）
CREATE TABLE IF NOT EXISTS "users" (
    -- 用戶的唯一標識符
    id UUID PRIMARY KEY,
    -- 用戶的名稱
    user_name VARCHAR(100) NOT NULL,
    -- 用戶的電子郵件地址
    email VARCHAR(255) UNIQUE,
    -- Line ID
    line_id VARCHAR(30) UNIQUE,
    -- Google ID
    google_id VARCHAR(30) UNIQUE,
    -- Apple ID
    apple_id VARCHAR(30) UNIQUE,
    -- 用戶的出生日期
    date_of_birth DATE,
    -- 用戶的性別（M 表示男性，F 表示女性，U 未知）
    gender CHAR(1) CHECK (gender IN ('M', 'F', 'U')),
    -- 用戶的地址
    address TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- 為表添加註釋
COMMENT ON TABLE "users" IS '用於存儲用戶信息的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN "users".id IS '用戶的唯一標識符';
COMMENT ON COLUMN "users".user_name IS '用戶的名稱';
COMMENT ON COLUMN "users".email IS '用戶的電子郵件地址';
COMMENT ON COLUMN "users".line_id IS 'Line ID';
COMMENT ON COLUMN "users".google_id IS 'Google ID';
COMMENT ON COLUMN "users".apple_id IS 'Apple ID';
COMMENT ON COLUMN "users".date_of_birth IS '用戶的出生日期';
COMMENT ON COLUMN "users".gender IS '用戶的性別（M 表示男性，F 表示女性，U 未知）';
COMMENT ON COLUMN "users".address IS '用戶的地址';
