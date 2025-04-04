package services

import (
	"PrimeGin/internal/models"
	"PrimeGin/internal/utils"
	"errors"
)

type UserService struct{}

func (service *UserService) GetUsers() ([]models.User, error) {
	if utils.DB == nil { // 直接使用 utils.DB
		return nil, errors.New("database connection is not initialized")
	}

	var users []models.User
	if err := utils.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
