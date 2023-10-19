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

type CorsConfig struct {
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
	AllowedMethods string `mapstructure:"ALLOWED_METHODS"`
	AllowedHeaders string `mapstructure:"ALLOWED_HEADERS"`
}

type SwagConfig struct {
	Host string `mapstructure:"SWAGGER_HOST"`
}

type Config struct {
	Port      string `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	DB        DBConfig
	Cors      CorsConfig
	Swagger   SwagConfig
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

	viper.SetDefault("ALLOWED_ORIGINS", "*")
	viper.SetDefault("ALLOWED_METHODS", "GET HEAD POST PUT DELETE OPTIONS PATCH")
	viper.SetDefault("ALLOWED_HEADERS", "Content-Type Authorization Accept Cache-Control Allow")

	viper.SetDefault("SWAGGER_HOST", "localhost:8080")

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Cfg.DB); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Cfg.Cors); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Cfg.Swagger); err != nil {
		return err
	}

	return nil
}
