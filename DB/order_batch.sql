-- 進貨批次表
CREATE TABLE IF NOT EXISTS order_batch (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- pk
    batch_no VARCHAR(50) NOT NULL,                  -- 批次號
    name VARCHAR(255) NOT NULL,                     -- 名稱
    description TEXT                               -- 說明
);

-- 為表添加註釋
COMMENT ON TABLE order_batch IS '進貨批次表';

-- 為列添加註釋
COMMENT ON COLUMN order_batch.id IS '進貨批次ID';
COMMENT ON COLUMN order_batch.batch_no IS '批次號';
COMMENT ON COLUMN order_batch.name IS '名稱';
COMMENT ON COLUMN order_batch.description IS '說明';
