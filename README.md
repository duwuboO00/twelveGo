# TwelveGo 電商後端服務

這是一個使用 Go 語言開發的電商網站後端系統。

## 功能特色

- 使用者認證與授權 (JWT + Session)
- 商品管理系統
- 購物車功能
- 訂單管理
- 會員系統
- 商品分類管理

## 技術棧

- 語言：Go 1.21+
- Web 框架：Gin
- 資料庫：PostgreSQL
- 快取：Redis (待實現)
- 會話管理：gorilla/sessions
- API 文件：OpenAPI/Swagger (待實現)

## 快速開始

### 環境需求

- Go 1.21+
- PostgreSQL
- Docker (選用)

### 安裝與執行

1. 複製專案

```bash
git clone [repository-url]
```

2. 安裝依賴

```bash
go mod download
```

3. 設定配置

```bash
cp config_example.yaml config.yaml
# 編輯 config.yaml 設定相關參數
```

4. 執行專案

```bash
./scripts/start.sh
```

5. Docker 部屬

```bash
docker build -f deployments/Dockerfile -t twelve-go .
docker run -p 8080:8080 twelve-go
```

## TODO

### backend

- [ ] DB update
  - [x] 完成
  - [ ] revicew 完成
- [ ] select one apis
  - [ ] 完成
  - [ ] revicew 完成
- [ ] create apis
  - [ ] 完成
  - [ ] revicew 完成
- [ ] update apis
  - [ ] 完成
  - [ ] revicew 完成
- [ ] delete apis
  - [ ] 完成
  - [ ] revicew 完成
- [ ] 庫存計算邏輯
  - [ ] 完成
  - [ ] revicew 完成
- [ ] oAuth API
  - [ ] Google
    - [ ] 完成
    - [ ] revicew 完成
  - [ ] Line
    - [ ] 完成
    - [ ] revicew 完成
  - [ ] Discord
    - [ ] 完成
    - [ ] revicew 完成
- [ ] session
  - [ ] 完成
  - [ ] revicew 完成
- [ ] JWT
  - [ ] 完成
  - [ ] revicew 完成
- [ ] Role & 權限認證
  - [ ] 完成
  - [ ] revicew 完成
- [ ] API 權限認證
  - [ ] 完成
  - [ ] revicew 完成
- [ ] 權限認證 filter / middleware
  - [ ] 完成
  - [ ] revicew 完成

### infra

#### docker compose

- [ ] db docker
  - [ ] 完成
  - [ ] revicew 完成
- [ ] docker sever
  - [ ] 完成
  - [ ] revicew 完成
