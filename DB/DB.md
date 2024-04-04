# 電商網站資料表清單

| 中文名稱         | 英文名稱        | 資料表名稱        | 資料表用途                                      |
|------------------|----------------|-------------------|-----------------------------------------------|
| 使用者           | User           | [CREATE_USER.sql](./DB/CREATE_USER.sql)           | 儲存使用者資訊                                |
| 訂單             | Order          | [CREATE_ORDER.sql](./DB/CREATE_ORDER.sql)         | 儲存訂單資訊                                  |
| 訂單商品項目     | Order Item     | [CREATE_ORDER_ITEM.sql](./DB/CREATE_ORDER_ITEM.sql) | 儲存訂單內商品的資訊                         |
| 商品             | Product        | [CREATE_PRODUCT.sql](./DB/CREATE_PRODUCT.sql)     | 儲存商品資訊                                  |
| 商品分類         | Product Category | [CREATE_PRODUCT_CATEGORY.sql](./DB/CREATE_PRODUCT_CATEGORY.sql) | 儲存商品分類資訊 |
| 購物車           | Shopping Cart  | [CREATE_SHOPPING_CART.sql](./DB/CREATE_SHOPPING_CART.sql) | 儲存購物車資訊                             |
| 付款方式         | Payment Method | [CREATE_PAYMENT_METHOD.sql](./DB/CREATE_PAYMENT_METHOD.sql) | 儲存付款方式資訊                             |
| 權限             | Permission     | [CREATE_PERMISSION.sql](./DB/CREATE_PERMISSION.sql) | 儲存使用者權限資訊                           |
| 最愛商品         | Favorite       | [CREATE_FAVORITE.sql](./DB/CREATE_FAVORITE.sql)   | 儲存使用者最愛的商品資訊                     |
| 商品與分類關聯   | Product Category Relation | [CREATE_PRODUCT_CATEGORY_RELATION.sql](./DB/CREATE_PRODUCT_CATEGORY_RELATION.sql) | 儲存商品與分類之間的關聯 |
| 使用者與權限關聯 | User Permission Relation | [CREATE_USER_PERMISSION_RELATION.sql](./DB/CREATE_USER_PERMISSION_RELATION.sql) | 儲存使用者與權限之間的關聯 |

## 注意事項
- 本資料表清單僅供參考，實際應用中可能需要根據需求進行調整和擴展。
- 每個資料表的結構和欄位定義可以在相對應的 `.sql` 文件中查看。
