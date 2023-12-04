package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	UserID uint   `validate:"required"`
	Title  string `validate:"required" gorm:"unique"`
}

func (m Goal) TableName() string {
	return "goals"
}

func (m Goal) customValidator() *validator.Validate {
	v := validator.New()
	return v
}

func (m *Goal) BeforeSave(db *gorm.DB) error {
	// Валидация
	validate := m.customValidator()
	if err := validate.Struct(m); err != nil {
		return err
	}

	return nil
}

type GoalCreateRequest struct {
	Title string `json:"title"`
}

type GoalGetResponse struct {
	Title string `json:"title"`
	ID    uint   `json:"id"`
}
