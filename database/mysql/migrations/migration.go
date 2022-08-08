package main

import (
	"log"

	"github.com/joho/godotenv"
	_config "github.com/wildanie12/go-hex-arch-sample/config"
	_mysql "github.com/wildanie12/go-hex-arch-sample/database/mysql"
	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
)

//go:generate ./migrate.sh

func main() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println(_color.This("red", "[config] cannot load .env file configuration"))
	}

	config := _config.New()

	_ = _mysql.New(config)

}