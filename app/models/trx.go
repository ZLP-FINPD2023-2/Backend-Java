package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TrxResponse struct {
	Name   string          `json:"name"`
	Date   time.Time       `json:"date"`
	Amount decimal.Decimal `json:"amount"`
}

type TrxRequest struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Amount string `json:"amount"`
}

type Trx struct {
	gorm.Model
	Name   string          `validate:"required"`
	Date   time.Time       `validate:"required,isNotFutureDate"`
	Amount decimal.Decimal `validate:"required"`
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

func customValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("isNotFutureDate", IsNotFutureDate)
	return v
}

func (trx *Trx) BeforeSave(db *gorm.DB) error {
	// Валидация
	validate := customValidator()
	if err := validate.Struct(trx); err != nil {
		return err
	}

	return nil
}
