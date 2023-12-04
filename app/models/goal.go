package models

import (
	"finapp/lib/validators"

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

func (m *Goal) BeforeSave(db *gorm.DB) error {
	// Валидация
	validate := validators.CustomValidator()
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
