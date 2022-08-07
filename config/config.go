package config

// AppConfig main struct
type AppConfig struct {
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
		Name string
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
		initHTTP(appConfig)
		initGRPC(appConfig)
	}
	return appConfig
}

