package services

import (
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type FoodService interface {
	GetFoodById(id uint) (*models.Food, error)
	GetFoodsBySpeciesId(speciesId uint) (*[]models.Food, error)
}

type GormFoodService struct {
	DB *gorm.DB
}

func (s GormFoodService) GetFoodById(id uint) (*models.Food, error) {
	return models.RetrieveFood(id, s.DB)
}

func (s GormFoodService) GetFoodsBySpeciesId(speciesId uint) (*[]models.Food, error) {
	return models.RetrieveFoods(speciesId, s.DB)
}
