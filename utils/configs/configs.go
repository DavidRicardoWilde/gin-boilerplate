package configs

import "gin-boilerplate/configs"

var (
	globalAppConfig       *configs.AppConfig
	globalApiServerConfig *configs.ApiServerConfig
	globalDbConfig        *configs.DbConfig
)

func AllConfigInit() {
	// gin engine web server config
	globalAppConfig = initAppConfig()
	globalApiServerConfig = initApiServerConfig()
	globalDbConfig = initDbConfig()

	// database config
	// redis config
}

func GetGlobalAppConfig() *configs.AppConfig {
	return globalAppConfig
}

func GetGlobalAppServerCfg() *configs.ApiServerConfig {
	return globalApiServerConfig
}

func GetGlobalDbConfig() *configs.DbConfig {
	return globalDbConfig
}

func initAppConfig() *configs.AppConfig {
	var appCfg *configs.AppConfig
	var cfg map[string]string

	if configs.Local {
		cfg = configs.LocalConfig.GetStringMapString("app")
	} else {
		cfg = configs.GlobalConfig.GetStringMapString("app")
	}

	appCfg = &configs.AppConfig{
		Name:        cfg["name"],
		Version:     cfg["version"],
		Description: cfg["description"],
		Environment: cfg["environment"],
		LogLevel:    cfg["log-level"],
	}

	return appCfg
}

func initApiServerConfig() *configs.ApiServerConfig {
	var apiServerCfg *configs.ApiServerConfig
	var cfg map[string]interface{}

	if configs.Local {
		cfg = configs.LocalConfig.GetStringMap("api-server")
	} else {
		cfg = configs.GlobalConfig.GetStringMap("api-server")
	}

	apiServerCfg = &configs.ApiServerConfig{
		BasePath:   cfg["base-path"].(string),
		ServerPort: cfg["server-port"].(string),
		Cors:       cfg["cors"].(bool),
	}

	return apiServerCfg
}

func initDbConfig() *configs.DbConfig {
	var dbCfg *configs.DbConfig
	var cfg map[string]interface{}

	if configs.Local {
		cfg = configs.LocalConfig.GetStringMap("database-config")
	} else {
		cfg = configs.GlobalConfig.GetStringMap("database-config")
	}

	dbCfg = &configs.DbConfig{
		Driver: cfg["driver"].(string),
		Usr:    cfg["usr"].(string),
		Pwd:    cfg["pwd"].(string),
		Host:   cfg["host"].(string),
		Port:   cfg["port"].(string),
		DbName: cfg["db-name"].(string),
	}

	return dbCfg
}

func GetBoolByKey(key string) bool {
	if configs.Local {
		return configs.LocalConfig.GetBool(key)
	} else {
		return configs.GlobalConfig.GetBool(key)
	}
}
