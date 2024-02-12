package models

import (
	"time"
	"github.com/google/uuid"
)

type ShoppingCart struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ProductID uuid.UUID `db:"product_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c *ShoppingCart) AddItem() error {
	// 實現添加項目到購物車的邏輯
	return nil
}

func (c *ShoppingCart) UpdateItem() error {
	// 實現更新購物車項目的邏輯
	return nil
}

func (c *ShoppingCart) RemoveItem() error {
	// 實現從購物車移除項目的邏輯
	return nil
}
