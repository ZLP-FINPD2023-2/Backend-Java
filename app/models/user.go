package models

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      *string `gorm:"unique" validate:"required,email"`
	Password   string  `validate:"required,min=8"`
	FirstName  string  `validate:"required"`
	LastName   string  `validate:"required"`
	Patronymic string
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
	// Валидация
	validate := CustomValidator()
	if err := validate.Struct(user); err != nil {
		return err
	}

	// Хэширование пароля
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	return nil
}
