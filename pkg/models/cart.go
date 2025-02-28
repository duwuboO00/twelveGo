package models

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type ShoppingCart struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ProductID uuid.UUID `db:"product_id"`
	Quantity  int       `db:"quantity"` // 新增欄位
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// AddItem 新增項目至購物車
func (c *ShoppingCart) AddItem() error {
	if c.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	// 實現添加項目到購物車的邏輯
	return nil
}

// UpdateItem 更新購物車項目
func (c *ShoppingCart) UpdateItem() error {
	if c.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	// 實現更新購物車項目的邏輯
	return nil
}

// RemoveItem 從購物車移除項目
func (c *ShoppingCart) RemoveItem() error {
	// 實現從購物車移除項目的邏輯
	return nil
}
