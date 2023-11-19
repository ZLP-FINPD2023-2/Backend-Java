package domains

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finapp/models"
)

type TrxService interface {
	WithTrx(trxHandle *gorm.DB) TrxService
	Get(c *gin.Context) ([]models.Trx, error)
}
