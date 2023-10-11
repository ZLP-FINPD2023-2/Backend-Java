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
	Port      string `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	DB        DBConfig
}

var Cfg Config

// Заполняет структуру параметрами окружения
func InitConfig() error {
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("SECRET_KEY", "secret_key")
	viper.SetDefault("POSTGRES_USER", "admin")
	viper.SetDefault("POSTGRES_PASSWORD", "password")
	viper.SetDefault("POSTGRES_DB", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSL_MODE", "disable")

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Cfg.DB); err != nil {
		return err
	}

	return nil
}
