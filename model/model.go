package model

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Key     string  `json:"key"`
	Value    string `json:"value"`
}

func (Model) TableName() string {
	return "models"
}