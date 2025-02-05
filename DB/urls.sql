-- 建立 urls 表（如果不存在）
CREATE TABLE IF NOT EXISTS urls (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- 短網址記錄唯一ID
    key         TEXT UNIQUE,                                 -- 短網址 key，對應 target_url
    secret_key  TEXT UNIQUE,                                 -- 管理短網址的 secret key
    target_url  TEXT,                                        -- 原始連結
    is_active   BOOLEAN DEFAULT TRUE,                        -- 短網址是否有效
    clicks      INT DEFAULT 0                                -- 短網址被點擊次數
);

-- 為表添加註釋
COMMENT ON TABLE urls IS '用於存儲短網址資訊的表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN urls.id         IS '短網址記錄唯一ID';
COMMENT ON COLUMN urls.key        IS '短網址 key，對應 target_url';
COMMENT ON COLUMN urls.secret_key IS '管理短網址的 secret key';
COMMENT ON COLUMN urls.target_url IS '原始連結';
COMMENT ON COLUMN urls.is_active  IS '短網址是否有效';
COMMENT ON COLUMN urls.clicks     IS '短網址被點擊次數';
