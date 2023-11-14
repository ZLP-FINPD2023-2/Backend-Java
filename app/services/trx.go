package services

import (
	"gorm.io/gorm"

	"finapp/domains"
	"finapp/lib"
	"finapp/repository"
)

type TrxService struct {
	logger     lib.Logger
	repository repository.TrxRepository
}

func NewTrxService(logger lib.Logger, repository repository.TrxRepository) domains.TrxService {
	return TrxService{
		logger:     logger,
		repository: repository,
	}
}
func (s TrxService) WithTrx(trxHandle *gorm.DB) domains.TrxService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}