package config

import (
	"log"
	"os"

	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
)


func initMySQL(appConfig *AppConfig) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	db := os.Getenv("MYSQL_DB")

	if host == "" {
		host = "localhost"
		log.Println(_color.ThisF("yellow", "[config] MySQL host is not set, using host `%s` instead", host))
	}

	if port == "" {
		port = "3306"
		log.Println(_color.ThisF("yellow", "[config] MySQL port is not set, using port `%s` instead", port))
	}

	if user == "" {
		user = "root"
		log.Println(_color.ThisF("yellow", "[config] MySQL user is not set, using username `%s` instead", user))
	}

	if pass == "" {
		pass = "root"
		log.Println(_color.ThisF("yellow", "[config] MySQL pass is not set, using password `%s` instead", pass))
	}

	if db == "" {
		db = "go_hex_arch"
		log.Println(_color.ThisF("yellow", "[config] MySQL db is not set, using database `%s` instead", db))
	}

	appConfig.MySQL.Host = host
	appConfig.MySQL.Port = port
	appConfig.MySQL.Username = user
	appConfig.MySQL.Password = pass
	appConfig.MySQL.Database = db
}