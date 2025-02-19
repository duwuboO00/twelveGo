# 電商網站資料表清單

下表依據你在 `/DB` 資料夾中的實際檔案列出，並可視專案需求填寫對應的用途說明。

| 中文名稱                       | 英文名稱                             | 資料表名稱                                                      | 資料表用途                |
|-------------------------------|--------------------------------------|-----------------------------------------------------------------|---------------------------|
| 地址                          | Address                              | [address.sql](./address.sql)                                    | 儲存地址相關資訊         |
| 廠商                          | Company                              | [company.sql](./company.sql)                                    | 儲存廠商資訊             |
| 最愛商品                      | Favorite                             | [favorite.sql](./favorite.sql)                                  | 儲存使用者最愛的商品資訊 |
| 圖片                          | Images                               | [images.sql](./images.sql)                                      | 儲存商品/其他圖片相關資訊|
| 訂單                          | Order                                | [order.sql](./order.sql)                                        | 儲存訂單資訊             |
| 進貨批次                      | Order Batch                          | [order_batch.sql](./order_batch.sql)                            | 儲存進貨批次資訊         |
| 訂單商品項目                  | Order Item                           | [order_item.sql](./order_item.sql)                              | 儲存訂單內商品的資訊     |
| 付款方式                      | Payment Method                       | [payment_method.sql](./payment_method.sql)                      | 儲存可用的付款方式       |
| 商品                          | Product                              | [product.sql](./product.sql)                                    | 儲存商品資訊             |
| 商品分類                      | Product Category                     | [product_category.sql](./product_category.sql)                  | 儲存商品分類資訊         |
| 商品與分類關聯               | Product-Product Category Relation    | [product_product_category_relation.sql](./product_product_category_relation.sql) | 儲存商品與分類的關聯 |
| 購物車                        | Shopping Cart                        | [shopping_cart.sql](./shopping_cart.sql)                        | 儲存購物車資訊           |
| 縮網址                          | URLs                                 | [urls.sql](./urls.sql)                                          | 檔案映射網址用         |
| 使用者                        | User                                 | [users.sql](./users.sql)                                          | 儲存使用者資訊           |
| 使用者權限                    | User Permission                      | [user_permission.sql](./user_permission.sql)                    | 儲存使用者權限定義       |
| 使用者權限關聯                | User Permission Relation             | [user_permission_relation.sql](./user_permission_relation.sql)  | 儲存使用者與權限關聯     |
| 使用者 Sessions               | User Sessions                        | [user_sessions .sql](./user_sessions.sql)                      | 儲存使用者 Session 資訊  |

## 注意事項
- 實際用途及字段定義可依需求在對應的 `.sql` 檔案內編修並保持同步。  
- 若有新增或刪除表，請在此表中更新以確保檔案與文件相符。
