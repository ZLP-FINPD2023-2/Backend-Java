package models

import (
	"gorm.io/gorm"
	"time"
)

// Transaction представляет модель данных для транзакций
type Transaction struct {
	gorm.Model
	UserID     uint      `gorm:"column:user_id"`
	Amount     float64   `gorm:"column:amount"`
	Currency   string    `gorm:"column:currency"`
	Reason     string    `gorm:"column:reason"`
	OccurredAt time.Time `gorm:"column:occurred_at"`
}

type TransactionRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Reason   string  `json:"reason"`
}

// TableName указывает на имя таблицы в базе данных для модели Transaction
func (Transaction) TableName() string {
	return "transactions"
}
