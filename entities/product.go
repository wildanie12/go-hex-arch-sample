package entities

import (
	"time"

	"gorm.io/gorm"
)

// Product model
type Product struct {
	ID uint 							`gorm:"primarykey" faker:"-"`
	Name string							`size:"255" faker:"name"`
	Description string					`faker:"sentence"`
	Price int64	
	Category string						`size:"255" faker:"word"`
	Seller string						`size:"255" faker:"name"`
	Slug string							`size:"255" faker:"word"`
    CreatedAt time.Time					`faker:"-"`
    UpdatedAt time.Time					`faker:"-"`
    DeletedAt gorm.DeletedAt 			`gorm:"index" faker:"-"`

	ProductVariants []ProductVariant	`gorm:"foregnKey:ProductID;references:ID" faker:"-"`
}