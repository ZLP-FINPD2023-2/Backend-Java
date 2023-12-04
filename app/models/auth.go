package models

import "github.com/dgrijalva/jwt-go"

const DateFormat = "02-01-2006"

type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type RegisterRequest struct {
	Email      *string `json:"email"`
	Password   string  `json:"password"`
	FirstName  string  `json:"firstname"`
	LastName   string  `json:"lastname"`
	Patronymic string  `json:"patronymic,omitempty"`
	Gender     Gender  `json:"gender"`
	Birthday   string  `json:"birthday"`
}

type LoginRequest struct {
	Email    *string `json:"email"`
	Password string  `json:"password"`
}
