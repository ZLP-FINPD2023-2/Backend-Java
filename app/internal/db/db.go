package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/internal/config"
	"app/internal/models"
)

var DB *gorm.DB

func Init() error {
	cfg := config.Cfg
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connect to sqlite...")
		dsn := "sqlite.db"
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
	}
	log.Println("Connected to database")

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Migrated database")

	return nil
}
