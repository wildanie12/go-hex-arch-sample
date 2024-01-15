package entity

import "time"

// Product entity.
type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	SKU         string    `json:"sku"`
	OwnerID     uint      `json:"owner_id"`
	Description string    `json:"description"`
	Price       uint64    `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
