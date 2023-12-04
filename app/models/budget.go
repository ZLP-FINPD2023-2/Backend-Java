package models

import (
	"gorm.io/gorm"

	"finapp/lib/validators"
)

type Budget struct {
	gorm.Model
	UserID uint   `validate:"required"`
	Title  string `validate:"required" gorm:"unique"`
	Goal   uint
}

func (b Budget) TableName() string {
	return "budgets"
}

func (b *Budget) BeforeSave(db *gorm.DB) error {
	// Валидация
	validate := validators.CustomValidator()
	if err := validate.Struct(b); err != nil {
		return err
	}

	return nil
}

type BudgetCreateRequest struct {
	Title string `json:"title"`
	Goal  uint   `json:"goal"`
}

type BudgetGetResponse struct {
	Title string `json:"title"`
	ID    uint   `json:"id"`
	Goal  uint   `json:"goal"`
}
