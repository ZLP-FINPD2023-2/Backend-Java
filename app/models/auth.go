package models

type RegisterRequest struct {
	Email      *string `json:"email"`
	Password   string  `json:"password"`
	FirstName  string  `json:"firstname"`
	LastName   string  `json:"lastname"`
	Patronymic string  `json:"patronymic,omitempty"`
}

type LoginRequest struct {
	Email    *string `json:"email"`
	Password string  `json:"password"`
}
