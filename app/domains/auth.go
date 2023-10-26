package domains

import "finapp/models"

type AuthService interface {
	Authorize(tokenString string) (bool, error)
	CreateToken(models.User) string
	HashPassword(password string) ([]byte, error)
}
