package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

// Инициализирует структуру Config
// Заполняет структуру параметрами из файла или окружения
func InitConfig() (*Config, error) {
	viper.SetDefault("PORT", ":8080")

	viper.SetConfigName(".env.dev")
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
