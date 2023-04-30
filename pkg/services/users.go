package services

import (
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserById(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id uint) error
}

type GormUserService struct {
	DB *gorm.DB
}

func (s GormUserService) GetUserById(id uint) (*models.User, error) {
	return models.RetrieveUser(id, s.DB)
}

func (s GormUserService) CreateUser(user *models.User) error {
	return models.CreateUser(user, s.DB)
}

func (s GormUserService) DeleteUser(id uint) error {
	return models.DeleteUser(id, s.DB)
}
