package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID            uuid.UUID `db:"id"`
	OrderID       uuid.UUID `db:"order_id"`
	ProductID     uuid.UUID `db:"product_id"`
	ProductLink   string    `db:"product_link"`
	ProductName   string    `db:"product_name"`
	ProductSubtitle string  `db:"product_subtitle"`
	UnitPrice     float64   `db:"unit_price"`
	Quantity      int       `db:"quantity"`
}

func (oi *OrderItem) Create() error {
	// 實現創建訂單明細的邏輯
	return nil
}

func (oi *OrderItem) Update() error {
	// 實現更新訂單明細的邏輯
	return nil
}

func (oi *OrderItem) Delete() error {
	// 實現刪除訂單明細的邏輯
	return nil
}
