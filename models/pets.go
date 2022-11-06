package models

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	UserID       uint          `json:"userId"`
	Name         string        `json:"name"`
	Breeds       []Breed       `gorm:"many2many:pet_breeds;"`
	SpeciesID    uint          `json:"speciesId"`
	Species      Species       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Age          uint          `json:"age"`
	Images       []Image       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Meals        []Meal        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Medications  []Medication  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HealthEvents []HealthEvent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func CreatePet(pet *Pet, DB *gorm.DB) error {
	var err error = DB.Create(&pet).Error
	return err
}

func RetrievePet(petId uint, userId uint, DB *gorm.DB) (*Pet, error) {
	var pet *Pet
	var err error = DB.First(&pet, "id = ? AND user_id = ?", petId, userId).Error
	return pet, err
}

func RetrievePets(userId uint, DB *gorm.DB) (*[]Pet, error) {
	var pets *[]Pet
	var err error = DB.Find(&pets, "user_id = ?", userId).Error
	return pets, err
}

func UpdatePet(userId uint, petId uint, pet *Pet, DB *gorm.DB) error {
	var err error = DB.Model(&pet).Where("id = ? AND user_id = ?", petId, userId).Updates(&pet).Error
	return err
}

func DeletePet(petId uint, userId uint, DB *gorm.DB) error {
	var pet *Pet
	var err error = DB.Delete(&pet, "id = ? AND user_id = ?", petId, userId).Error
	return err
}
