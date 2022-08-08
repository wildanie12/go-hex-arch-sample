package migrations

import (
	"log"

	_config "github.com/wildanie12/go-hex-arch-sample/config"
	_mysql "github.com/wildanie12/go-hex-arch-sample/database/mysql"
	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
	"gorm.io/gorm"
)

var config *_config.AppConfig
var db *gorm.DB

func init() {
	config = _config.New()
	db = _mysql.New(config)
}

// Start a migration seeder process
func Start() {
	err := migrate(db)
	if err != nil {
		log.Println(_color.ThisF("red", "[migration] failed running migration: %v", err))
		return
	}
	log.Println(_color.ThisF("green", "[migration] ✅ done running migration"))
	
	err = seed(db)
	if err != nil {
		log.Println(_color.ThisF("red", "[migration] failed seeding data: %v", err))
		return
	}
	log.Println(_color.ThisF("green", "[migration] ✅ done running data seeder "))
}