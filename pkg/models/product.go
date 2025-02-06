package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Price        float64   `db:"price"`
	StockQuantity int      `db:"stock_quantity"`
	CategoryID    uuid.UUID `db:"category_id"`
	Description  string    `db:"description"`
	MediaLink1   string    `db:"media_link_1"`
	MediaLink2   string    `db:"media_link_2"`
	MediaLink3   string    `db:"media_link_3"`
	MediaLink4   string    `db:"media_link_4"`
	MediaLink5   string    `db:"media_link_5"`
	MediaLink6   string    `db:"media_link_6"`
}

type ProductCategory struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

func (p *Product) Create() error {
	// 實現創建商品的邏輯
	return nil
}

func (p *Product) Update() error {
	// 實現更新商品的邏輯
	return nil
}

func (p *Product) Delete() error {
	// 實現刪除商品的邏輯
	return nil
}
