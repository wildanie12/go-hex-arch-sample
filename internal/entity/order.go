package entity

import "time"

// Order entity.
type Order struct {
	ID                  uint   `json:"id"`
	UserID              uint   `json:"user_id"`
	CustomerName        string `json:"customer_name"`
	CustomerAddress     string `json:"customer_address"`
	CustomerPhoneNumber string `json:"customer_phone_number"`
	Subtotal            uint64 `json:"subtotal"`
	ShipmentFee         uint64 `json:"shipping_fee"`
	Total               uint64 `json:"total"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	OrderItemlines      *[]OrderItemline `json:"order_item_lines"`
}

// OrderItemline belongs to order entity.
type OrderItemline struct {
	ID          uint      `json:"id"`
	OrderID     uint      `json:"order_id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	ProductName string    `json:"product_name"`
	ProductSKU  string    `json:"product_sku"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
