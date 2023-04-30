package services

import (
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type MedicineService interface {
	GetMedicineById(id uint) (*models.Medicine, error)
	GetMedicinesBySpeciesId(speciesId uint) (*[]models.Medicine, error)
}

type GormMedicineService struct {
	DB	*gorm.DB
}

func (s *GormMedicineService) GetMedicineById(id uint) (*models.Medicine, error) {
	return models.GetMedicine(id, s.DB)
}

func (s *GormMedicineService) GetMedicinesBySpeciesId(speciesId uint) (*[]models.Medicine, error) {
	return models.GetMedicines(speciesId, s.DB)
}