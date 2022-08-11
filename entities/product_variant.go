package entities

type ProductVariant struct {
	ID uint 			`gorm:"primaryKey"`
	ProductID uint
	Sku string			`gorm:"size:255" faker:"uuid_hyphenated"`
	Name string			`gorm:"size:255" faker:"word"`
	Price int64
	Description string	`faker:"paragraph"`
	Product Product		`gorm:"foreignKey:ProductID;references:ID" faker:"-"`
}