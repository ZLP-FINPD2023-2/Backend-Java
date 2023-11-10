package domains

import (
	"gorm.io/gorm"
)

type TrxService interface {
	WithTrx(trxHandle *gorm.DB) TrxService
}
