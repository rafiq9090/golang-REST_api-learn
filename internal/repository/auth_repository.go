package repository

import (
	"task-api/internal/database"
	"task-api/internal/models"
)

type AuthRepo struct{}

var Auth = AuthRepo{}

func (AuthRepo) Register(user *models.User) error {
	return database.DB.Create(user).Error
}

func (AuthRepo) Login(email string, password string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (AuthRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
