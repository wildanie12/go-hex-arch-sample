package entities

import (
	"time"
)

// ProductVariant main struct, tagged with gorm and faker
type ProductVariant struct {
	ID uint 							`gorm:"primaryKey"`
	ProductID uint
	Sku string							`gorm:"size:255" faker:"word"`
	Name string							`gorm:"size:255" faker:"word"`
	Price int64
	Description string					`faker:"paragraph"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product						`gorm:"foreignKey:ProductID;references:ID" faker:"-"`
	ProductQuantity []ProductQuantity	`gorm:"polymorphic:Productable"`
}