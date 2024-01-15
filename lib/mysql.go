package lib

import (
	"log"

	"github.com/wildanie12/go-hex-arch-sample/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectOrderDB creates database connection to order mysql db.
func (di *DatabaseInstance) ConnectOrderDB() {
	var err error
	di.OrderDB, err = gorm.Open(mysql.Open(config.Get().MySQL.OrderURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// CloseOrderDB gracefully shutdown order db instance.
func (di *DatabaseInstance) CloseOrderDB() {
	if di.OrderDB == nil {
		return
	}
	db, err := di.OrderDB.DB()
	if err != nil {
		log.Println(err)
		return
	}
	db.Close()
}
