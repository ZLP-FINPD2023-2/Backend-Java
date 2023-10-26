package domains

import (
	"finapp/models"

	"gorm.io/gorm"
)

type UserService interface {
	WithTrx(trxHandle *gorm.DB) UserService
	Authorize(q *models.LoginRequest) (models.User, error)
	Register(q *models.RegisterRequest) error
}
