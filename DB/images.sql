-- 建立 images 表（如果不存在）
CREATE TABLE IF NOT EXISTS images (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- 圖片ID
    name        VARCHAR(255) NOT NULL,                       -- 圖片名稱
    extension   VARCHAR(10)  NOT NULL,                       -- 圖片副檔名
    image_data  BYTEA,                                       -- 圖片二進位數據
    path        VARCHAR(255),                                -- 圖片路徑
    tags        JSONB                                        -- 圖片標籤
);

-- 為表添加註釋
COMMENT ON TABLE images IS '圖片表';

-- 為每個欄位添加註釋
COMMENT ON COLUMN images.id         IS '圖片ID';
COMMENT ON COLUMN images.name       IS '圖片名稱';
COMMENT ON COLUMN images.extension  IS '圖片副檔名';
COMMENT ON COLUMN images.image_data IS '圖片二進位數據';
COMMENT ON COLUMN images.path       IS '圖片路徑';
COMMENT ON COLUMN images.tags       IS '圖片標籤';
