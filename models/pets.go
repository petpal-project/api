package models

import (
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	UserID      uint         `json:"userId"`
	Name        string       `json:"name"`
	Breeds      []Breed      `gorm:"many2many:pet_breeds;"`
	SpeciesID   uint         `json:"speciesId"`
	Species     Species      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Age         uint         `json:"age"`
	Images      []Image      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Meals       []Meal       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Medications []Medication `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Events      []Event      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (pet Pet) GetUserID() uint { return pet.UserID }
func (pet Pet) GetID() uint     { return pet.ID }

func (pet *Pet) BeforeUpdate(DB *gorm.DB) error {
	return CheckOwnership[Pet]("pet", DB)
}

func (pet *Pet) BeforeDelete(DB *gorm.DB) error {
	return CheckOwnership[Pet]("pet", DB)
}

func CreatePet(pet *Pet, DB *gorm.DB) error {
	return DB.Create(&pet).Error
}

func RetrievePet(petId uint, userId uint, DB *gorm.DB) (*Pet, error) {
	var pet *Pet
	err := DB.First(&pet, "id = ? AND user_id = ?", petId, userId).Error
	return pet, err
}

func RetrievePets(userId uint, DB *gorm.DB) (*[]Pet, error) {
	var pets *[]Pet
	err := DB.Find(&pets, "user_id = ?", userId).Error
	return pets, err
}

func UpdatePet(userId uint, petId uint, pet *Pet, DB *gorm.DB) (*Pet, error) {
	err := DB.Set("user", userId).Set("pet", petId).Model(&pet).Where("id = ?", petId).Updates(&pet).Error
	return pet, err
}

func DeletePet(petId uint, userId uint, DB *gorm.DB) error {
	var pet *Pet
	return DB.Set("user", userId).Set("pet", petId).Delete(&pet, "id = ? AND user_id = ?", petId, userId).Error
}
