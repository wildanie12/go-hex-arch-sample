package config

import (
	"log"
	"os"

	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
)

const httpDefaultPort = "8000"

func initHTTP(appConfig *AppConfig) {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		log.Println(_color.ThisF("yellow", "[config] HTTP Port is not set, using port %s instead", httpDefaultPort))
		port = httpDefaultPort
	}
	if host == "" {
		log.Println(_color.This("yellow", "[config] HTTP Host is not set, using 127.0.0.1 instead"))
		host = "127.0.0.1"
	}
	appConfig.HTTP.Host = host
	appConfig.HTTP.Port = port
}