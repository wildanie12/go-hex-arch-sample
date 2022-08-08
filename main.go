package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	_grpc "github.com/wildanie12/go-hex-arch-sample/app/grpc"
	_config "github.com/wildanie12/go-hex-arch-sample/config"
	_migrate "github.com/wildanie12/go-hex-arch-sample/database/mysql/migrations"
)

func main() {
	cli()
	config := _config.New()

	appGRPC := _grpc.New(config.GRPC.Host, config.GRPC.Port)
	appGRPC.Start()
}


func cli() {
	args := os.Args
	if len(args) <= 1 {
		return
	}

	if args[1] == "-m" || args[1] == "--migrate" {
		_migrate.Start()
	}
	os.Exit(0)
}