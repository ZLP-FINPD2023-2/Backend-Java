package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// TODO: реализовать gender с помощью enum
type User struct {
	gorm.Model
	Email      string `json:"email" gorm:"unique" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	FirstName  string `json:"firstname" validate:"required"`
	LastName   string `json:"lastname" validate:"required"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        uint8  `json:"age" validate:"required,numeric,gte=0"`
	Gender     bool   `json:"gender" validate:"required"`
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
