package migrations

import (
	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(

		// Entity Migration, write here...
		&_entities.Product{},
		&_entities.ProductVariant{},
		&_entities.ProductQuantity{},
	)

	return err
}