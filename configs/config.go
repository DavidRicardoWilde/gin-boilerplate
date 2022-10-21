package configs

import "github.com/spf13/viper"

type AppConfig struct {
	Name        string `yaml:"name" toml:"name" json:"name"`
	Version     string `yaml:"version" toml:"version" json:"version"`
	Description string `yaml:"description" toml:"description" json:"description"`
	Environment string `yaml:"environment" toml:"environment" json:"environment"`
	LogLevel    string `yaml:"log-level" toml:"log-level" json:"log-level"`
}

type ApiServerConfig struct {
	BasePath   string `yaml:"base-path" toml:"base-path" json:"basePath"`
	ServerPort string `yaml:"server-port" toml:"base-path" json:"serverPort"`
	Cors       bool   `yaml:"cors" toml:"cors" json:"cors"`
}

type DbConfig struct {
	Driver string
	Usr    string
	Pwd    string
	Host   string
	Port   int
	DbName string
}

var (
	Local        = false
	LocalConfig  *viper.Viper
	GlobalConfig *viper.Viper
	//PrvConfig    *viper.Viper
)

func init() {
	localConfigInit()
	globalConfigInit()
	// load private config from online
	//prvConfigInit()
}

// prvConfigInit loads private config from online
//func prvConfigInit() {
//	PrvConfig.AddSecureRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.json", "/etc/secrets/s.gpg")
//	PrvConfig.SetConfigType("json") // because there is no file extension in a stream of bytes,  supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
//	err := PrvConfig.ReadInConfig()
//	if err != nil {
//		logs.Log.Error("viper read private config online error, error: ", err)
//	}
//
//	if GlobalConfig.Get("app") == nil {
//		logs.Log.Error("Load private config failed")
//	}
//	logs.Log.Info("Load private config success")
//}

// globalConfigInit loads global config from local which is used for all environments
func globalConfigInit() {
	GlobalConfig = viper.New()
	GlobalConfig.AddConfigPath("./configs")
	GlobalConfig.SetConfigName("config")
	GlobalConfig.SetConfigType("toml")
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		//logs.Log.Error("viper read config.toml error, error: ", err)
	}

	if GlobalConfig.Get("app") == nil {
		//logs.Log.Error("Load global config failed")
	}
	//logs.Log.Info("Load global config success")
}

// localConfigInit loads local config from local which is used for private development environment
func localConfigInit() {
	LocalConfig = viper.New()
	LocalConfig.AddConfigPath("./configs")
	LocalConfig.SetConfigName("local")
	LocalConfig.SetConfigType("yaml")

	err := LocalConfig.ReadInConfig()
	if err != nil {
		//logs.Log.Error("viper read local.yaml error, error: ", err)
	}

	if LocalConfig.Get("env") == "dev" {
		Local = true
		//logs.Log.Error("Load local config failed")
	}
	//logs.Log.Info("Load local config success")
}
