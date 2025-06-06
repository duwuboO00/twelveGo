-- 廠商表
CREATE TABLE IF NOT EXISTS company (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- pk
    name VARCHAR(255) NOT NULL,                      -- 廠商名稱
    description TEXT,                                -- 說明
    url1 TEXT,                                       -- 網址1
    url2 TEXT,                                       -- 網址2
    url3 TEXT,                                       -- 網址3
    alias1 TEXT,                                     -- 別稱1
    alias2 TEXT,                                     -- 別稱2
    alias3 TEXT,                                     -- 別稱3
    address TEXT,                                    -- 地址
    phone TEXT                                       -- 電話
);

-- 為表添加註釋
COMMENT ON TABLE company IS '廠商表';
COMMENT ON COLUMN company.id IS '廠商ID';
COMMENT ON COLUMN company.name IS '廠商名稱';
COMMENT ON COLUMN company.description IS '廠商說明';
COMMENT ON COLUMN company.url1 IS '廠商網址1';
COMMENT ON COLUMN company.url2 IS '廠商網址2';
COMMENT ON COLUMN company.url3 IS '廠商網址3';
COMMENT ON COLUMN company.alias1 IS '廠商別稱1';
COMMENT ON COLUMN company.alias2 IS '廠商別稱2';
COMMENT ON COLUMN company.alias3 IS '廠商別稱3';
COMMENT ON COLUMN company.address IS '廠商地址';
COMMENT ON COLUMN company.phone IS '廠商電話';
