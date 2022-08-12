package config

import (
	"log"
	"os"

	_color "github.com/wildanie12/go-hex-arch-sample/utils/color"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


func initApp(appConfig *AppConfig) {
	env := os.Getenv("APP_ENV")

	switch cases.Lower(language.English).String(env) {

	case "production":
		appConfig.App.Env = "production"
		log.Println(_color.This("green", "✅ using production environment"))
		return
	case "development":
		appConfig.App.Env = "development"
		log.Println(_color.This("blue", "✅ using development environment"))
		return
	default:
		appConfig.App.Env = "development"
		log.Println(_color.This("blue", "✅ env is not specified, using development environment instead"))
		return
	}
}