package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Gender string

const (
	Female Gender = "Female"
	Male   Gender = "Male"
)

type User struct {
	gorm.Model
	Email      *string `gorm:"unique" validate:"required,email"`
	Password   string  `validate:"required"`
	FirstName  string  `validate:"required"`
	LastName   string  `validate:"required"`
	Patronymic string
	Gender     Gender    `validate:"oneof=Male Female"`
	Birthday   time.Time `validate:"required"`
}

// TableName gives table name of model
func (u User) TableName() string {
	return "users"
}

func CustomValidator() *validator.Validate {
	v := validator.New()
	return v
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	validate := CustomValidator()
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
