package config

// AppConfig main struct
type AppConfig struct {
	App struct {
		Env string
	}
	HTTP struct {
		Host string
		Port string
	}
	GRPC struct {
		Host string
		Port string
	}
	MySQL struct {
		Host string
		Port string
		Database string
		Username string
		Password string
	}
}

var appConfig *AppConfig

// New config instance object
func New() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{}	

		// config modules initializations
		initApp(appConfig)
		initHTTP(appConfig)
		initGRPC(appConfig)
		initMySQL(appConfig)
	}
	return appConfig
}

