package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Name     string `mapstructure:"POSTGRES_DB"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
}

type Config struct {
	Port string `mapstructure:"PORT"`
	DB   DBConfig
}

// Инициализирует структуру Config
// Заполняет структуру параметрами окружения
func InitConfig() (*Config, error) {
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("POSTGRES_USER", "admin")
	viper.SetDefault("POSTGRES_PASSWORD", "password")
	viper.SetDefault("POSTGRES_DB", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSL_MODE", "disable")

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config.DB); err != nil {
		return nil, err
	}

	return &config, nil
}
