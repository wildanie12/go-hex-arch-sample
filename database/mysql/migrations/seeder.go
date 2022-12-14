package migrations

import (
	"log"
	"math/rand"
	"strconv"

	"gorm.io/gorm"

	"github.com/bxcodec/faker/v3"
	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
)


func seed(db *gorm.DB) error {

	// Product
	for i := 0; i < 10; i++ {
		product := _entities.Product{}
		if err := faker.FakeData(&product); err != nil {
			log.Println(_color.ThisF("yellow", "[seeder] failed faking data: %v", err))
		}
		db.Save(&product)
		log.Println(_color.ThisF("green", "[seeder] ✅ success seeding product data: [id: %d, product: %s]", product.ID, product.Name))


		// Product Variants
		nVariants := rand.Intn(5)
		for i := 0; i < nVariants; i++ {
			variant := _entities.ProductVariant{}
			if err := faker.FakeData(&variant); err != nil {
				log.Println(_color.ThisF("yellow", "[seeder] failed faking data: %v", err))
			}
			variant.ProductID = product.ID
			variant.Sku = variant.Name + "-" + strconv.Itoa(rand.Intn(100))
			db.Save(&variant)
			log.Println(_color.ThisF("green", "[seeder] ✅ success seeding |- product variant: [id: %d, name: %s]", variant.ID, variant.Name))

			// Product Quantity variants
			nQuantities := rand.Intn(3)
			for j := 0; j < nQuantities; j++ {
				quantity := _entities.ProductQuantity{}
				if err := faker.FakeData(&quantity); err != nil {
					log.Println(_color.ThisF("yellow", "[seeder] failed faking data: %v", err))
				}
				quantity.ProductableType = "product_variants"
				quantity.ProductableID = variant.ID
				quantity.CurrentStock = quantity.InitialStock - rand.Intn(5)
				db.Save(&quantity)
				log.Println(_color.ThisF("green", "[seeder] ✅ success seeding    |- product quantity: [id: %d, init stock: %d]", quantity.ID, quantity.InitialStock))
			}

		}

		// Product Quantity (if variants non-existence)
		nQuantities := rand.Intn(3)
		for j := 0; j < nQuantities; j++ {
			if nVariants <= 0 {
				quantity := _entities.ProductQuantity{}
				if err := faker.FakeData(&quantity); err != nil {
					log.Println(_color.ThisF("yellow", "[seeder] failed faking data: %v", err))
				}
				quantity.ProductableType = "products"
				quantity.ProductableID = product.ID
				quantity.CurrentStock = quantity.InitialStock - rand.Intn(5)
				db.Save(&quantity)
				log.Println(_color.ThisF("green", "[seeder] ✅ success seeding |- product quantity: [id: %d, init stock: %d]", quantity.ID, quantity.InitialStock))
			}
		}
	}

	return nil
}
