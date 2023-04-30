package services

import (
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type SpeciesService interface {
	GetSpeciesById(id uint) (*models.Species, error)
}

type GormSpeciesService struct {
	DB *gorm.DB
}

func (s GormSpeciesService) GetSpeciesById(id uint) (*models.Species, error) {
	return models.RetrieveSpecies(id, s.DB)
}
