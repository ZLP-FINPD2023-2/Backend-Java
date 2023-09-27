package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

// заполняет конфиг значениями из файла с параметрами окружения
func InitConfig() (*Config, error) {
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName(".env.dev")
	viper.SetConfigType("env")

	viper.SetDefault("Port", ":8080")

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
