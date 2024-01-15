package order

import (
	"time"
)

// Order entity.
type Order struct {
	ID                  uint   `gorm:"primaryKey" json:"id"`
	UserID              uint   `json:"user_id"`
	CustomerName        string `gorm:"size:255" json:"customer_name"`
	CustomerAddress     string `json:"customer_address"`
	CustomerPhoneNumber string `gorm:"size:255" json:"customer_phone_number"`
	Subtotal            uint64 `json:"subtotal"`
	ShipmentFee         uint64 `json:"shipping_fee"`
	Total               uint64 `json:"total"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	OrderItemlines      *[]Itemline `gorm:"foreignKey:OrderID;references:ID" json:"order_item_lines"`
}

// Itemline belongs to order entity.
type Itemline struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	OrderID     uint      `json:"order_id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	ProductName string    `gorm:"size:255" json:"product_name"`
	ProductSKU  string    `gorm:"size:255" json:"product_sku"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Order       *Order    `gorm:"foreignKey:OrderID;references:ID"  json:"order"`
}

// TableName return table name of order itemline.
func (Itemline) TableName() string {
	return "order_itemlines"
}

// Interface contains order repository contracts.
type Interface interface {
	// FindAll orders.
	FindAll() ([]Order, error)

	// FindAllByUser orders.
	FindAllByUser(userID int) ([]Order, error)

	// Find single order by ID.
	Find(id uint) (Order, error)

	// Create single order.
	Create(Order) (Order, error)

	// Update single order.
	Update(Order) (Order, error)

	// Delete single order.
	Delete(Order) (Order, error)

	// AddItemline into order.
	AddItemline(oil Itemline, orderID uint) (Itemline, error)

	// RemoveItemline into order.
	RemoveItemline(oil Itemline) (Itemline, error)
}
