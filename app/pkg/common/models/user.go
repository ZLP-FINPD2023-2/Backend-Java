package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"unique" json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Patronymic string `json:"patronymic"`
	Age        uint8  `json:"age"`
	Gender     bool   `json:"gender"`
}
