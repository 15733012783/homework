package config

import "github.com/spf13/viper"

type Config struct {
}

func ReadConfig(filePath string) {
	viper.SetConfigFile("./config/goods.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
