package services

import (
	"gorm.io/gorm"
	"time"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
	"finapp/repository"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, repository repository.UserRepository) domains.UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) domains.UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// Authorize call to register the user
func (s UserService) Authorize(q *models.LoginRequest) (models.User, error) {
	var user models.User
	err := s.repository.Where("email = ?", q.Email).First(&user).Error
	return user, err
}

// Register call to register the user
func (s UserService) Register(q *models.RegisterRequest) error {
	birthday, err := time.Parse(models.DateFormat, q.Birthday)
	if err != nil {
		return err
	}
	user := models.User{
		Email:      q.Email,
		Password:   q.Password,
		FirstName:  q.FirstName,
		LastName:   q.LastName,
		Patronymic: q.Patronymic,
		Gender:     q.Gender,
		Birthday:   birthday,
	}

	return s.repository.Create(&user).Error
}

// UpdateUser updates the user
//func (s UserService) UpdateUser(user models.User) error {
//	return s.repository.Save(&user).Error
//}

// DeleteUser deletes the user
//func (s UserService) DeleteUser(id uint) error {
//	return s.repository.Delete(&models.User{}, id).Error
//}
