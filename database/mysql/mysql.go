package mysql

import (
	"fmt"
	"log"

	_config "github.com/wildanie12/go-hex-arch-sample/config"
	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New MySQL Connection, created using gorm dialector
func New(appConfig *_config.AppConfig) *gorm.DB {
	if appConfig == nil {
		log.Println(_color.This("red", "[Database - MySQL] app config is not yet initialized, cancelling mysql connection..."))
		return nil
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		appConfig.MySQL.Username,
		appConfig.MySQL.Password,
		appConfig.MySQL.Host,
		appConfig.MySQL.Port,
		appConfig.MySQL.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(_color.ThisF("red", "[Database - MySQL] failed to open connection: %v", err))
		log.Println(_color.ThisF("blue", "[Database - MySQL] dsn: %s", dsn))
	}
	return db
}