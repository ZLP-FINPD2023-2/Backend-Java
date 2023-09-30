package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/pkg/common/config"
	"app/pkg/common/models"
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

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Migrated database")

	return db, nil
}
