package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	UserID uint   `validate:"required"`
	Title  string `validate:"required" gorm:"unique"`
}

func (b Budget) TableName() string {
	return "budgets"
}

func (b Budget) customValidator() *validator.Validate {
	v := validator.New()
	return v
}

func (b *Budget) BeforeSave(db *gorm.DB) error {
	// Валидация
	validate := b.customValidator()
	if err := validate.Struct(b); err != nil {
		return err
	}

	return nil
}

type BudgetCreateRequest struct {
	Title string `json:"title"`
}

type BudgetGetResponse struct {
	Title string `json:"title"`
	ID    uint   `json:"id"`
}
