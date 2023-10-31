package domains

import "finapp/models"

type AuthService interface {
	Authorize(tokenString string) (bool, int, error)
	CreateToken(models.User) string
}
