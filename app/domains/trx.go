package domains

import (
	"finapp/models"
	"gorm.io/gorm"
)

type TrxService interface {
	WithTrx(trxHandle *gorm.DB) TrxService
	Get(name string) (models.Trx, error)
}
