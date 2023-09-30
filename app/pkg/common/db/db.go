package db

import (
	"app/pkg/common/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/pkg/common/config"
)

func Init(config config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.User, config.Password, config.Name, config.Host, config.Port, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connect to sqlite...")
		dsn := "sqlite.db"
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	log.Println("Connected to database")
	db.AutoMigrate(&models.User{})
	return db, nil
}
