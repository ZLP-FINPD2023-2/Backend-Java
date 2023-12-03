package services

import (
	"gorm.io/gorm"

	"finapp/domains"
	"finapp/lib"
	"finapp/models"
	"finapp/repository"
)

type GoalService struct {
	logger     lib.Logger
	repository repository.GoalRepository
}

func NewGoalService(
	logger lib.Logger,
	repository repository.GoalRepository,
) domains.GoalService {
	return GoalService{
		logger:     logger,
		repository: repository,
	}
}

func (s GoalService) WithTrx(trxHandle *gorm.DB) domains.GoalService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s GoalService) List(userID uint) ([]models.Goal, error) {
	var goals []models.Goal

	// Создание запроса
	query := s.repository.Database.Where("user_id = ?", userID)

	// Выполнение запроса
	err := query.Find(&goals).Error

	return goals, err
}

func (s GoalService) Create(request *models.GoalCreateRequest, userID uint) error {
	goal := models.Goal{
		UserID: userID,
		Title:  request.Title,
	}

	return s.repository.Database.Create(&goal).Error
}
