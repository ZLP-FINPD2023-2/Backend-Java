package repository

import (
	"gorm.io/gorm"

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
