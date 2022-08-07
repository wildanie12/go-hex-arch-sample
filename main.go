package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	_config "github.com/wildanie12/go-hex-arch-sample/config"
	_jsonUtil "github.com/wildanie12/go-hex-arch-sample/utils/json"
)

func main() {
	config := _config.New()

	fmt.Println(_jsonUtil.Encode(config))
}