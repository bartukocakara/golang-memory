package repository

import (
	"golang-memory/model"

	"github.com/jinzhu/gorm"
)

type ModelRepository interface {
	GetAll() ([]model.Model, error)
}

type modelRepository struct {
	connection *gorm.DB
}

func NewModelRepository() ModelRepository {
	return &modelRepository{
		connection: DB(),
	}
}

func (db *modelRepository) GetAll() (models []model.Model, err error) {
	return models, db.connection.Find(&models).Error
}