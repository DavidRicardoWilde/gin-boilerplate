package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Env          string
	GlobalConfig *viper.Viper
)

func LoadConfigFile(path string) {
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigFile(path)
	if err := GlobalConfig.ReadInConfig(); err != nil {
		logrus.WithError(err).Error("viper read config file error")
	}
}
