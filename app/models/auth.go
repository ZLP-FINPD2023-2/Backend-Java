package models

import "time"

type RegisterRequest struct {
	Email      *string   `json:"email"`
	Password   string    `json:"password"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Patronymic string    `json:"patronymic,omitempty"`
	Gender     Gender    `json:"gender"`
	BirthDate  time.Time `json:"birthDate"`
}

type LoginRequest struct {
	Email    *string `json:"email"`
	Password string  `json:"password"`
}
