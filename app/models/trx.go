package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TrxResponse struct {
	Title      string          `json:"title"`
	Date       time.Time       `json:"date"`
	Amount     decimal.Decimal `json:"amount"`
	BudgetFrom uint            `validate:"required"`
	BudgetTo   uint            `validate:"required"`
}

type TrxRequest struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	Amount     string `json:"amount"`
	BudgetFrom uint   `json:"from" validate:"required"`
	BudgetTo   uint   `json:"to" validate:"required"`
}

type Trx struct {
	gorm.Model
	UserID     uint            `validate:"required"`
	Title      string          `validate:"required"`
	Date       time.Time       `validate:"required,isNotFutureDate"`
	Amount     decimal.Decimal `validate:"required" sql:"type:decimal(20,2);"`
	BudgetFrom uint            `validate:"required"`
	BudgetTo   uint            `validate:"required"`
}

func (t Trx) TableName() string {
	return "transactions"
}

func IsNotFutureDate(fldLvl validator.FieldLevel) bool {
	dateToValidate, ok := fldLvl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	dateToValidate = dateToValidate.UTC()
	currentDate := time.Now().UTC()

	return dateToValidate.Before(currentDate) || dateToValidate.Equal(currentDate)
}

func (trx Trx) customValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("isNotFutureDate", IsNotFutureDate)
	return v
}

func (trx *Trx) BeforeSave(db *gorm.DB) error {
	// Валидация
	validate := trx.customValidator()
	if err := validate.Struct(trx); err != nil {
		return err
	}

	return nil
}
