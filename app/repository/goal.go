package repository

import (
	"gorm.io/gorm"

	"finapp/lib"
	"finapp/models"
)

type GoalRepository struct {
	logger   lib.Logger
	Database lib.Database
}

func NewGoalRepository(logger lib.Logger, db lib.Database) GoalRepository {
	return GoalRepository{
		logger:   logger,
		Database: db,
	}
}

func (r GoalRepository) WithTrx(trxHandle *gorm.DB) GoalRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r GoalRepository) Get(id, userID uint) (models.Goal, error) {
	var goal models.Goal
	err := r.Database.Where("user_id = ?", userID).Where("id = ?", id).First(&goal).Error
	return goal, err
}
