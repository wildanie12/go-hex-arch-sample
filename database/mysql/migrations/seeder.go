package migrations

import (
	"log"

	"gorm.io/gorm"

	"github.com/bxcodec/faker/v3"
	_entities "github.com/wildanie12/go-hex-arch-sample/entities"
	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
)


func seed(db *gorm.DB) error {

	// Product
	for i := 0; i < 10; i++ {
		data := _entities.Product{}
		err := faker.FakeData(&data)
		if err != nil {
			log.Println(_color.ThisF("yellow", "[seeder] failed faking data: %v", err))
		}
		db.Save(&data)
		log.Println(_color.ThisF("green", "[seeder] ✅ success seeding product data: [id: %d, product: %s]", data.ID, data.Name))
	}
	return nil
}
