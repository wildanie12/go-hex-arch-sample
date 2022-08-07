package main

import (
	_ "github.com/joho/godotenv/autoload"
	_grpc "github.com/wildanie12/go-hex-arch-sample/app/grpc"
	_config "github.com/wildanie12/go-hex-arch-sample/config"
)

func main() {
	config := _config.New()

	appGRPC := _grpc.New(config.GRPC.Host, config.GRPC.Port)
	appGRPC.Start()
}