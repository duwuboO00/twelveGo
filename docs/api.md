# 電商網站 API 文件

## 登入

### 登入

- **URL:** `/auth/login`
- **Method:** POST
- **Description:** 使用者登入
- **Request:**
  - Body:
    ```json
    {
        "username": "user123",
        "password": "password123"
    }
    ```
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "message": "登入成功",
        "user_id": "user123"
    }
    ```

### 登出

- **URL:** `/auth/logout`
- **Method:** POST
- **Description:** 使用者登出
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "message": "登出成功"
    }
    ```

## 商品

### 搜尋商品

- **URL:** `/products/search`
- **Method:** GET
- **Description:** 搜尋商品
- **Parameters:**
  - `keyword` (optional): 關鍵字
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "results": [
            {
                "product_id": "p001",
                "name": "商品一",
                "price": 100
            },
            {
                "product_id": "p002",
                "name": "商品二",
                "price": 200
            }
        ]
    }
    ```

### 商品分類

- **URL:** `/products/categories`
- **Method:** GET
- **Description:** 取得商品分類列表
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "categories": ["電子產品", "服飾", "食品"]
    }
    ```

## 購物車

### 列出購物車物品

- **URL:** `/cart`
- **Method:** GET
- **Description:** 列出購物車內的商品
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "items": [
            {
                "product_id": "p001",
                "name": "商品一",
                "quantity": 2
            }
        ]
    }
    ```

### 加入購物車

- **URL:** `/cart/add`
- **Method:** POST
- **Description:** 加入商品至購物車
- **Request:**
  - Body:
    ```json
    {
        "product_id": "p002",
        "quantity": 1
    }
    ```
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "message": "已加入購物車"
    }
    ```

### 移出購物車

- **URL:** `/cart/remove/{id}`
- **Method:** DELETE
- **Description:** 從購物車移除商品
- **Parameters:**
  - `id` (path, required): 商品 ID
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "message": "已從購物車移除"
    }
    ```

## 其他 API

### 伺服器健康檢查

- **URL:** `/heartbeat`
- **Method:** GET
- **Description:** 確認伺服器是否運行

### Google OAuth 登入

- **URL:** `/auth/google`
- **Method:** GET
- **Description:** 進入 Google OAuth 登入頁面

### Google OAuth 回調

- **URL:** `/auth/google/callback`
- **Method:** GET
- **Parameters:**
  - `code` (query, required): Google OAuth 授權碼
- **Description:** Google OAuth 回調處理

### 更新用戶名稱

- **URL:** `/auth/updateName`
- **Method:** PUT
- **Description:** 更新用戶名稱
- **Request:**
  - Body:
    ```json
    {
        "name": "Alice"
    }
    ```

### 取得當前用戶資訊

- **URL:** `/user/profile`
- **Method:** GET
- **Description:** 返回目前登入用戶資訊

### 上傳檔案

- **URL:** `/upload`
- **Method:** POST
- **Description:** 上傳單個或多個檔案

### 列出所有上傳檔案

- **URL:** `/upload`
- **Method:** GET
- **Description:** 取得所有上傳的檔案

### 取得指定檔案

- **URL:** `/upload/{id}`
- **Method:** GET
- **Parameters:**
  - `id` (path, required): 檔案 ID
- **Description:** 取得指定檔案

### 會員 API 健康檢查

- **URL:** `/member/heartbeat`
- **Method:** GET
- **Description:** 會員 API 健康檢查
