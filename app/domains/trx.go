package domains

import (
	"gorm.io/gorm"

	"finapp/models"
)

type TrxService interface {
	WithTrx(trxHandle *gorm.DB) TrxService
	Create(trxRequest *models.TrxRequest) error
}
