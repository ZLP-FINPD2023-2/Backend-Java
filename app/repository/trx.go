package repository

import (
	"gorm.io/gorm"

	"finapp/lib"
)

type TrxRepository struct {
	logger lib.Logger
	lib.Database
}

func NewTrxRepository(logger lib.Logger, db lib.Database) TrxRepository {
	return TrxRepository{
		logger:   logger,
		Database: db,
	}
}

func (r TrxRepository) WithTrx(trxHandle *gorm.DB) TrxRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
