package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
}

type Config struct {
	Port string `mapstructure:"PORT"`
	DB   DBConfig
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName(".env.dev")
	viper.SetConfigType("env")

	viper.SetDefault("Port", ":8080")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config.DB); err != nil {
		return nil, err
	}

	return &config, nil
}
