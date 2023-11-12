package repository

import (
	"finapp/models"
	"gorm.io/gorm"
	"time"

	"finapp/lib"
)

// UserRepository database structure
type UserRepository struct {
	logger lib.Logger
	lib.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository(logger lib.Logger, db lib.Database) UserRepository {
	return UserRepository{
		logger:   logger,
		Database: db,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r UserRepository) CreateTransaction(tx *gorm.DB, userID uint, amount float64, currency, reason string) error {
	transaction := models.Transaction{
		UserID:     userID,
		Amount:     amount,
		Currency:   currency,
		Reason:     reason,
		OccurredAt: time.Now(),
	}

	return tx.Create(&transaction).Error
}
