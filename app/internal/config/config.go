package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// App
	Host      string `mapstructure:"HOST"`
	Port      string `mapstructure:"PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	// DB
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBSSLMode  string `mapstructure:"DB_SSL_MODE"`
	// CORS
	CORSAllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
	CORSAllowedMethods string `mapstructure:"ALLOWED_METHODS"`
	CORSAllowedHeaders string `mapstructure:"ALLOWED_HEADERS"`
}

var Cfg Config

func InitConfig() error {
	// App
	viper.SetDefault("HOST", "localhost:8080")
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("SECRET_KEY", "secret_key")
	// DB
	viper.SetDefault("POSTGRES_USER", "admin")
	viper.SetDefault("POSTGRES_PASSWORD", "password")
	viper.SetDefault("POSTGRES_DB", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSL_MODE", "disable")
	// CORS
	viper.SetDefault("ALLOWED_ORIGINS", "*")
	viper.SetDefault("ALLOWED_METHODS", "GET HEAD POST PUT DELETE OPTIONS PATCH")
	viper.SetDefault("ALLOWED_HEADERS", "Content-Type Authorization Accept Cache-Control Allow")

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}
