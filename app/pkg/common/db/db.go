package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/pkg/common/config"
)

func Init(config config.DBConfig) (*gorm.DB, error) {
	user := config.User
	password := config.Password
	dbName := config.Name
	host := config.Host
	port := config.Port
	sslMode := config.SSLMode

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, dbName, host, port, sslMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database")
	return db, nil
}
