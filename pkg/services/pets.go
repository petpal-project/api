package services

import (
	"pet-pal/api/pkg/datasources"
	"pet-pal/api/pkg/models"

	"gorm.io/gorm"
)

type PetService interface {
	GetPetById(id uint) (*models.Pet, error)
	GetPetsByUserId(id uint) (*[]models.Pet, error)
	CreatePet(pet *models.Pet) error
	UpdatePet(userId uint, pet models.Pet) (*models.Pet, error)
	DeletePet(petId uint, userId uint) error
}

type GormPetService struct {
	DB *gorm.DB
}

func (s GormPetService) GetPetById(id uint) (*models.Pet, error) {
	return datasources.RetrieveRecordById[models.Pet](id, s.DB)
}

func (s GormPetService) GetPetsByUserId(id uint) (*[]models.Pet, error) {
	return datasources.RetrieveMultipleRecords[models.Pet](id, s.DB)
}

func (s GormPetService) CreatePet(pet *models.Pet) error {
	return datasources.CreateRecord(pet, s.DB)
}

func (s GormPetService) UpdatePet(userId uint, pet models.Pet) (*models.Pet, error) {
	return datasources.UpdateRecord(userId, pet, s.DB)
}

func (s GormPetService) DeletePet(petId uint, userId uint) error {
	return datasources.DeleteRecord[models.Pet](petId, userId, s.DB)
}
