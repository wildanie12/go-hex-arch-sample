package config

import (
	"log"
	"os"

	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
)

const gRPCDefaultPort = "9000"

func initGRPC(appConfig *AppConfig) {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		log.Println(_color.ThisF("yellow", "[config] gRPC Port is not set, using port %s instead", gRPCDefaultPort))
		port = gRPCDefaultPort
	}
	if host == "" {
		log.Println(_color.This("yellow", "[config] gRPC Host is not set, using 127.0.0.1 instead"))
		host = "127.0.0.1"
	}
	appConfig.GRPC.Host = host
	appConfig.GRPC.Port = port
}