package services

import (
	"finapp/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"

	"finapp/domains"
	"finapp/lib"
	"finapp/repository"
)

type TrxService struct {
	logger     lib.Logger
	repository repository.TrxRepository
}

func NewTrxService(logger lib.Logger, repository repository.TrxRepository) domains.TrxService {
	return &TrxService{
		logger:     logger,
		repository: repository,
	}
}
func (s TrxService) WithTrx(trxHandle *gorm.DB) domains.TrxService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}
func (s TrxService) Create(trxRequest *models.TrxRequest) error {
	date, err := time.Parse(models.DateFormat, trxRequest.Date)
	if err != nil {
		return err
	}
	amount, err := decimal.NewFromString(trxRequest.Amount)
	if err != nil {
		return err
	}
	transaction := models.Trx{
		Name:   trxRequest.Name,
		Date:   date,
		Amount: amount,
	}
	return s.repository.Create(&transaction).Error
}
