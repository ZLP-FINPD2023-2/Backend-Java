package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
	"finapp/repository"
)

type BudgetService struct {
	logger     lib.Logger
	repository repository.TrxRepository
}

func NewBudgetService(logger lib.Logger, repository repository.TrxRepository) domains.BudgetService {
	return BudgetService{
		logger:     logger,
		repository: repository,
	}
}

func (s BudgetService) WithTrx(trxHandle *gorm.DB) domains.BudgetService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s BudgetService) Get(c *gin.Context, userID uint) ([]models.Budget, error) {
	var budgets []models.Budget

	// Создание запроса
	query := s.repository.DB.Where("user_id = ?", userID)

	// Выполнение запроса
	err := query.Find(&budgets).Error

	return budgets, err
}

func (s BudgetService) Create(request *models.BudgetCreateRequest, userID uint) error {
	transaction := models.Budget{
		UserID: userID,
		Title:  request.Title,
	}

	return s.repository.Create(&transaction).Error
}
