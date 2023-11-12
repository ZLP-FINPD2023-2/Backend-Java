package services

import (
	"time"

	"gorm.io/gorm"

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

func (s UserService) GetUserByEmail(email *string) (models.User, error) {
	var user models.User
	err := s.repository.Where("email = ?", email).First(&user).Error
	return user, err
}

func (s UserService) Get(id uint) (models.User, error) {
	var user models.User
	err := s.repository.First(&user, id).Error
	return user, err
}

// UpdateUser updates the user
func (s UserService) Update(user *models.User) error {
	return s.repository.Save(user).Error
}

// Delete deletes the user
func (s UserService) Delete(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}

func (s UserService) CreateTransaction(userID uint, amount float64, currency, reason string) error {
	return s.repository.CreateTransaction(s.repository.DB, userID, amount, currency, reason)
}
