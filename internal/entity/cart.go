package entity

import "time"

// Cart entity.
type Cart struct {
	ID          uint      `json:"id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"owner_id"`
	ProductName string    `json:"name"`
	ProductSKU  string    `json:"sku"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
