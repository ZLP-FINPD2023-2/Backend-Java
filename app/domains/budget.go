package domains

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finapp/models"
)

type BudgetService interface {
	WithTrx(trxHandle *gorm.DB) BudgetService
	Get(c *gin.Context, userID uint) ([]models.Budget, error)
	Create(request *models.BudgetCreateRequest, userID uint) error
}
