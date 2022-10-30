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

func CreatePet(pet *Pet, DB *gorm.DB) {
	DB.Create(&pet)
}

func RetrievePet(petId uint, userId uint, DB *gorm.DB) *Pet {
	var pet *Pet
	DB.First(&pet, "id = ? AND user_id = ?", petId, userId)
	return pet
}

func DeletePet(petId uint, userId uint, DB *gorm.DB) {
	var pet *Pet
	DB.Delete(&pet, "id = ? AND user_id = ?", petId, userId)
}
