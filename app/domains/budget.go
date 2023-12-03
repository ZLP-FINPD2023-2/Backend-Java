package domains

import (
	"gorm.io/gorm"

	"finapp/models"
)

type BudgetService interface {
	WithTrx(trxHandle *gorm.DB) BudgetService
	List(userID uint) ([]models.Budget, error)
	Create(request *models.BudgetCreateRequest, userID uint) error
}
