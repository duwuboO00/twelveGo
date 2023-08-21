# 電商網站 API 文件

## 登入

### 登入

- **URL:** `/login`
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

- **URL:** `/logout`
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

- **URL:** `/cart/remove`
- **Method:** POST
- **Description:** 從購物車移除商品
- **Request:**
  - Body:
    ```json
    {
        "product_id": "p001"
    }
    ```
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
        "message": "已從購物車移除"
    }
    ```
