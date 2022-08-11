package entities

import (
	"time"

	"gorm.io/gorm"
)

// Product model
type Product struct {
	ID uint 							`gorm:"primarykey" faker:"-"`
	Name string							`faker:"name"`
	Description string					`faker:"sentence"`
	Price int64	
	Category string						`faker:"word"`
	Seller string						`faker:"name"`
	Slug string							`faker:"word"`
    CreatedAt time.Time					`faker:"-"`
    UpdatedAt time.Time					`faker:"-"`
    DeletedAt gorm.DeletedAt 			`gorm:"index" faker:"-"`

	ProductVariants []ProductVariant	`gorm:"foregnKey:ProductID;references:ID" faker:"-"`
}