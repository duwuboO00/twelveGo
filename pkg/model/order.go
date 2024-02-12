package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID                    uuid.UUID `db:"id"`
	UserID                uuid.UUID `db:"user_id"`
	OrderDate             time.Time `db:"order_date"`
	TaxAmount             float64   `db:"tax_amount"`
	ShippingFee           float64   `db:"shipping_fee"`
	TotalAmount           float64   `db:"total_amount"`
	EstimatedDeliveryDate time.Time `db:"estimated_delivery_date"`
	Status                string    `db:"status"`
}

func (o *Order) Create() error {
	// 實現創建訂單的邏輯
	return nil
}

func (o *Order) Update() error {
	// 實現更新訂單的邏輯
	return nil
}

func (o *Order) Delete() error {
	// 實現刪除訂單的邏輯
	return nil
}
