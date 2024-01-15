package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// AppConfig main struct.
type AppConfig struct {
	EnableFrontendPresenter bool
	EnableInternalPresenter bool
	EnableEventPresenter    bool
	HTTP                    struct {
		Port string
	}
	MySQL struct {
		OrderURI string
	}
}

var appConfig *AppConfig

// Get initialized application config.
func Get() *AppConfig {
	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	// read from dotenv files if available
	if err := godotenv.Load(); err != nil {
		fmt.Println("Fail to load from dotenv, using any os env configuration...")
	}

	// build and assign config
	cfg := AppConfig{}

	// presenter toggle switch config
	if strings.ToUpper(getEnvWarn("ENABLE_FRONTEND_PRESENTER")) != "FALSE" {
		cfg.EnableFrontendPresenter = true
	}

	// http config
	cfg.HTTP.Port = getEnvWarn("HTTP_PORT")

	// mysql config
	cfg.MySQL.OrderURI = getEnvWarn("MYSQL_ORDER_URI")

	return &cfg
}

func getEnvWarn(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("(!) config warn:", key, "is not yet configured")
	}
	return value
}

func makeError(err error) error {
	return fmt.Errorf("(x) config error: %v", err)
}
