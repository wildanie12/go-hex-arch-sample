package entities

import (
	"time"

	"gorm.io/gorm"
)

// ProductQuantity main struct, tagged with gorm and faker
type ProductQuantity struct {
	ID uint						`gorm:"primaryKey" faker:"-"`
	ProductableType	string		`gorm:"size:255" faker:"oneof: products, product_variants"`
	ProductableID uint			`faker:"-"`
	Barcode string				`gorm:"size:255" faker:"uuid_digit"`
	InitialStock int
	CurrentStock int			`faker:"-"`
	DateExpiry time.Time
	AlertExpiryDays int
	CreatedAt time.Time			`faker:"-"`
	UpdatedAt time.Time			`faker:"-"`
	DeletedAt gorm.DeletedAt 	`gorm:"index" faker:"-"`
}