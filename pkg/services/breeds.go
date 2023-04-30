package services

import (
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type BreedService interface {
	GetBreedById(id uint) (*models.Breed, error)
	GetBreedsBySpeciesId(id uint) (*[]models.Breed, error)
}

type GormBreedService struct {
	DB *gorm.DB
}

func (s GormBreedService) GetBreedById(id uint) (*models.Breed, error) {
	return models.RetrieveBreed(id, s.DB)
}

func (s GormBreedService) GetBreedsBySpeciesId(id uint) (*[]models.Breed, error) {
	return models.RetrieveBreeds(id, s.DB)
}
