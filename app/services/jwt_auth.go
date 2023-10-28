package services

import (
	"errors"

	"github.com/dgrijalva/jwt-go"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
)

// JWTAuthService service relating to authorization
type JWTAuthService struct {
	logger lib.Logger
	env    lib.Env
}

// NewJWTAuthService creates a new auth service
func NewJWTAuthService(logger lib.Logger, env lib.Env) domains.AuthService {
	return JWTAuthService{
		logger: logger,
		env:    env,
	}
}

// Authorize authorizes the generated token
func (s JWTAuthService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.env.SecretKey), nil
	})
	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("Token malformed")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("Token expired")
		}
	}
	return false, errors.New("Couldn't handle token")
}

// CreateToken creates jwt auth token
func (s JWTAuthService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(s.env.SecretKey))

	if err != nil {
		s.logger.Error("JWT validation failed: ", err)
	}

	return tokenString
}
