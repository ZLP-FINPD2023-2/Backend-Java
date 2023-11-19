package services

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
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

func (s TrxService) Get(c *gin.Context) ([]models.Trx, error) {
	var trxs []models.Trx

	// Нахождение параметров
	amountMinStr := c.Query("amount_min")
	amountMaxStr := c.Query("amount_max")
	dateFromStr := c.Query("date_from")
	dateToStr := c.Query("date_to")

	// Создание запроса
	query := s.repository.DB.Where(&models.Trx{})

	// Фильтрация запроса
	if amountMinStr != "" {
		amountMin, err := strconv.ParseFloat(amountMinStr, 64)
		if err != nil {
			return nil, err
		}
		query = query.Where("amount >= ?", amountMin)
	}

	if amountMaxStr != "" {
		amountMax, err := strconv.ParseFloat(amountMaxStr, 64)
		if err != nil {
			return nil, err
		}
		query = query.Where("amount <= ?", amountMax)
	}

	if dateFromStr != "" {
		dateFrom, err := time.Parse(models.DateFormat, dateFromStr)
		if err != nil {
			return nil, err
		}
		query = query.Where("date >= ?", dateFrom)
	}

	if dateToStr != "" {
		dateTo, err := time.Parse(models.DateFormat, dateToStr)
		if err != nil {
			return nil, err
		}
		query = query.Where("date <= ?", dateTo)
	}

	// Выполнение запроса
	err := query.Find(&trxs).Error

	return trxs, err
}
