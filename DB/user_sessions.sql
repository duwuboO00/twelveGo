-- 建立 user_sessions 表（如果不存在）
CREATE TABLE IF NOT EXISTS user_sessions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- 會話紀錄的唯一標識符
    user_id     UUID REFERENCES users(id),                   -- 對應的使用者 ID（外鍵，參考 users 表）
    session_id  VARCHAR(255) UNIQUE NOT NULL,                -- 會話 ID（唯一）
    ip_address  VARCHAR(255) NOT NULL,                       -- 連線 IP 地址
    user_agent  TEXT,                                        -- 用戶端的 User-Agent
    status      BOOLEAN,                                     -- 會話狀態（有效或已失效）
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- 建立時間
    deaded_at   TIMESTAMP WITH TIME ZONE,                    -- 會話終止時間
    UNIQUE (user_id, session_id)                             -- 同一 user_id 與 session_id 不可重複
);

-- 為表添加註釋
COMMENT ON TABLE user_sessions IS '用於存儲使用者會話資訊的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN user_sessions.id         IS '會話紀錄的唯一標識符';
COMMENT ON COLUMN user_sessions.user_id    IS '對應的使用者 ID（外鍵，參考 users 表）';
COMMENT ON COLUMN user_sessions.session_id IS '會話 ID（唯一）';
COMMENT ON COLUMN user_sessions.ip_address IS '連線 IP 地址';
COMMENT ON COLUMN user_sessions.user_agent IS '用戶端的 User-Agent';
COMMENT ON COLUMN user_sessions.status     IS '會話狀態（有效或已失效）';
COMMENT ON COLUMN user_sessions.created_at IS '會話建立時間，預設為當下時間';
COMMENT ON COLUMN user_sessions.deaded_at  IS '會話終止或失效的時間';
